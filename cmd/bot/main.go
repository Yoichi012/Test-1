package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/YourUsername/waifu-catcher/internal/bot"
	"github.com/YourUsername/waifu-catcher/internal/config"
	"github.com/YourUsername/waifu-catcher/internal/scheduler"
	"github.com/YourUsername/waifu-catcher/internal/storage"
)

func main() {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config load: %v", err)
	}

	// Connect to Mongo
	store, err := storage.NewMongoStore(cfg)
	if err != nil {
		log.Fatalf("mongo connect: %v", err)
	}
	defer func() {
		_ = store.Disconnect(context.Background())
	}()

	// Create Telegram bot
	tb, err := bot.NewBot(cfg, store)
	if err != nil {
		log.Fatalf("bot init: %v", err)
	}

	// Scheduler
	sched := scheduler.NewScheduler(cfg, store, tb)
	sched.Start()

	// Graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	log.Println("Bot running...")
	tb.Start(ctx)

	// Give scheduler time to stop
	ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	sched.Stop(ctx2)

	log.Println("Exiting.")
}
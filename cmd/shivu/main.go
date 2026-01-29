package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/yourusername/shivu-go/internal/bot"
	"github.com/yourusername/shivu-go/internal/config"
)

func main() {
	// Print banner
	printBanner()

	// Load configuration
	log.Println("Loading configuration...")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	log.Println("Configuration loaded successfully")

	// Initialize bot
	log.Println("Initializing bot...")
	botInstance, err := bot.NewBot(cfg.BotToken, cfg.MongoURL, cfg.DatabaseName)
	if err != nil {
		log.Fatalf("Failed to initialize bot: %v", err)
	}
	log.Printf("Bot initialized: @%s", cfg.BotUsername)

	// Setup graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start bot in a goroutine
	go func() {
		log.Println("Starting bot...")
		if err := botInstance.Start(); err != nil {
			log.Fatalf("Bot stopped with error: %v", err)
		}
	}()

	log.Println("Bot is running. Press CTRL+C to stop.")

	// Wait for shutdown signal
	<-quit
	log.Println("\nReceived shutdown signal...")

	// Graceful shutdown
	if err := botInstance.Stop(); err != nil {
		log.Printf("Error during shutdown: %v", err)
	}

	log.Println("Shutdown complete. Goodbye!")
}

func printBanner() {
	banner := `
╔═══════════════════════════════════════════╗
║                                           ║
║          SHIVU BOT - GO VERSION          ║
║                                           ║
║     Character Collection Telegram Bot     ║
║                                           ║
╚═══════════════════════════════════════════╝
	`
	log.Println(banner)
}

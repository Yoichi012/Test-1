package scheduler

import (
	"context"
	"log"
	"time"

	"github.com/robfig/cron/v3"

	"github.com/YourUsername/waifu-catcher/internal/config"
	"github.com/YourUsername/waifu-catcher/internal/storage"
	"github.com/YourUsername/waifu-catcher/internal/bot"
)

type Scheduler struct {
	cron  *cron.Cron
	cfg   *config.Config
	store *storage.MongoStore
	tb    *bot.Bot
}

func NewScheduler(cfg *config.Config, store *storage.MongoStore, tb *bot.Bot) *Scheduler {
	return &Scheduler{
		cron:  cron.New(cron.WithChain()),
		cfg:   cfg,
		store: store,
		tb:    tb,
	}
}

func (s *Scheduler) Start() {
	// Run every SpawnInterval minutes
	spec := "@every " + time.Duration(s.cfg.SpawnIntervalMins*int(time.Minute)).String()
	_, err := s.cron.AddFunc(spec, func() {
		// TODO: implement spawn: select group(s) and post a character
		log.Println("Scheduler tick: spawn job TODO")
	})
	if err != nil {
		log.Printf("cron add error: %v", err)
	}
	s.cron.Start()
}

func (s *Scheduler) Stop(ctx context.Context) {
	s.cron.Stop()
}
package bot

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/YourUsername/waifu-catcher/internal/config"
	"github.com/YourUsername/waifu-catcher/internal/storage"
	"github.com/YourUsername/waifu-catcher/internal/handlers"
)

type Bot struct {
	api    *tgbotapi.BotAPI
	cfg    *config.Config
	store  *storage.MongoStore
	ucfg   tgbotapi.UpdateConfig
}

func NewBot(cfg *config.Config, store *storage.MongoStore) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		return nil, err
	}
	api.Debug = false
	b := &Bot{
		api:   api,
		cfg:   cfg,
		store: store,
		ucfg:  tgbotapi.NewUpdate(0),
	}
	// register handlers with dependency injection
	handlers.Register(store, api, cfg)
	// Print bot username
	botUser, _ := api.GetMe()
	log.Printf("Authorized on account %s", botUser.UserName)
	return b, nil
}

func (b *Bot) Start(ctx context.Context) {
	u := b.ucfg
	u.Timeout = 60
	updates := b.api.GetUpdatesChan(u)

	for {
		select {
		case <-ctx.Done():
			log.Println("Bot received shutdown")
			return
		case upd := <-updates:
			handlers.HandleUpdate(b.api, &upd)
		}
	}
}
package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/YourUsername/waifu-catcher/internal/config"
	"github.com/YourUsername/waifu-catcher/internal/storage"
)

var store *storage.MongoStore
var api *tgbotapi.BotAPI
var cfg *config.Config

// Register stores DI references and sets up any caches
func Register(s *storage.MongoStore, a *tgbotapi.BotAPI, c *config.Config) {
	store = s
	api = a
	cfg = c
	log.Println("Handlers registered")
}
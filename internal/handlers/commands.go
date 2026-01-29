package handlers

import (
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleUpdate routes updates to command handlers
func HandleUpdate(apiInst *tgbotapi.BotAPI, update *tgbotapi.Update) {
	// use global api ref
	if update.Message == nil {
		// TODO: implement callback_query handling if needed
		return
	}

	if update.Message.IsCommand() {
		cmd := strings.ToLower(update.Message.Command())
		switch cmd {
		case "start":
			handleStart(apiInst, update)
		case "ping":
			handlePing(apiInst, update)
		case "upload":
			handleUploadCommand(apiInst, update)
		case "guess":
			handleGuessCommand(apiInst, update)
		case "changetime":
			handleChangeTime(apiInst, update)
		// add other command routes
		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Unknown command")
			apiInst.Send(msg)
		}
	}
}

func handleStart(apiInst *tgbotapi.BotAPI, update *tgbotapi.Update) {
	text := "Welcome to Waifu Catcher Bot (Go port)!\nCommands: /upload /guess /ping"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	apiInst.Send(msg)
}

func handlePing(apiInst *tgbotapi.BotAPI, update *tgbotapi.Update) {
	start := time.Now()
	m := tgbotapi.NewMessage(update.Message.Chat.ID, "Pong!")
	sent, _ := apiInst.Send(m)
	latency := time.Since(start).Milliseconds()
	edit := tgbotapi.NewEditMessageText(update.Message.Chat.ID, sent.MessageID, "Pong! "+time.Duration(latency).String())
	apiInst.Send(edit)
}
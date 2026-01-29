package handlers

import (
	"context"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson"
)

// /changetime <minutes> - owner only global update
func handleChangeTime(apiInst *tgbotapi.BotAPI, update *tgbotapi.Update) {
	if update.Message.From == nil {
		return
	}
	if int64(update.Message.From.ID) != cfg.OwnerID {
		apiInst.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Only bot owner can use this"))
		return
	}
	parts := update.Message.CommandArguments()
	mins, err := strconv.Atoi(parts)
	if err != nil {
		apiInst.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Usage: /changetime <minutes>"))
		return
	}
	// update all user_totals collection
	col := store.DB.Collection("user_totals")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = col.UpdateMany(ctx, bson.M{}, bson.M{"$set": bson.M{"message_frequency": mins}})
	if err != nil {
		apiInst.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Failed to update frequency"))
		return
	}
	apiInst.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Global frequency updated ✅"))
}
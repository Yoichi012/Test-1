package handlers

import (
	"context"
	"log"
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson"
)

// handleGuessCommand: simplified guessing flow — picks a random character and posts image URL
func handleGuessCommand(apiInst *tgbotapi.BotAPI, update *tgbotapi.Update) {
	col := store.DB.Collection("characters")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := col.Find(ctx, bson.M{})
	if err != nil {
		apiInst.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "No characters available"))
		return
	}
	defer cur.Close(ctx)
	var chars []map[string]interface{}
	for cur.Next(ctx) {
		var m map[string]interface{}
		if err := cur.Decode(&m); err == nil {
			chars = append(chars, m)
		}
	}
	if len(chars) == 0 {
		apiInst.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "No characters found"))
		return
	}
	rand.Seed(time.Now().UnixNano())
	ch := chars[rand.Intn(len(chars))]
	imageURL, _ := ch["image_url"].(string)
	text := "Guess the character from image:\n" + imageURL + "\nUse /guess <name> to answer"
	apiInst.Send(tgbotapi.NewMessage(update.Message.Chat.ID, text))
	log.Printf("spawned guess in chat %d", update.Message.Chat.ID)
}
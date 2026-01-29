package handlers

import (
	"context"
	"errors"
	"log"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/YourUsername/waifu-catcher/internal/models"
)

// handleUploadCommand parses /upload args and stores character
// Usage: /upload img_url character-name anime-name rarity
func handleUploadCommand(apiInst *tgbotapi.BotAPI, update *tgbotapi.Update) {
	args := update.Message.CommandArguments()
	parts := strings.Fields(args)
	if len(parts) < 4 {
		apiInst.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Usage: /upload img_url character-name anime-name rarity"))
		return
	}
	img := parts[0]
	name := parts[1]
	anime := parts[2]
	rarity, err := strconv.Atoi(parts[3])
	if err != nil {
		apiInst.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Rarity must be a number"))
		return
	}

	// Create Character
	ch := models.Character{
		ImageURL:     img,
		Name:         name,
		Anime:        anime,
		Rarity:       rarity,
		UploaderID:   int64(update.Message.From.ID),
		UploaderName: update.Message.From.UserName,
		CreatedAt:    time.Now(),
	}

	// save to Mongo
	col := store.DB.Collection("characters")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objID := primitive.NewObjectID()
	ch.ID = objID.Hex()
	_, err = col.InsertOne(ctx, ch)
	if err != nil {
		log.Printf("insert error: %v", err)
		apiInst.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Failed to save character"))
		return
	}

	apiInst.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Character uploaded ✅"))
}

// Note: more sophisticated upload (media download, catbox upload) can be implemented in utils/catbox.go
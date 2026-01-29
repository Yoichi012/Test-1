package handlers

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson"
)

// change user's balance by amount
func changeBalance(userID int64, amount int) error {
	col := store.DB.Collection("user_balance")
	ctx := context.Background()
	_, err := col.UpdateOne(ctx, bson.M{"user_id": userID}, bson.M{"$inc": bson.M{"balance": amount}}, optionsUpsert())
	return err
}

func optionsUpsert() interface{} {
	// small helper to avoid importing options in many files
	return bson.M{"upsert": true}
}

func handleBalanceCmd(apiInst *tgbotapi.BotAPI, update *tgbotapi.Update) {
	col := store.DB.Collection("user_balance")
	ctx := context.Background()
	var res map[string]interface{}
	err := col.FindOne(ctx, bson.M{"user_id": update.Message.From.ID}).Decode(&res)
	if err != nil {
		apiInst.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "0 coins"))
		return
	}
	bal := res["balance"]
	apiInst.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Balance: "+fmt.Sprintf("%v", bal)))
}
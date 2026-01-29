package models

import "time"

type Character struct {
	ID            string    `bson:"_id,omitempty" json:"id,omitempty"`
	ImageURL      string    `bson:"image_url" json:"image_url"`
	Name          string    `bson:"name" json:"name"`
	Anime         string    `bson:"anime" json:"anime"`
	Rarity        int       `bson:"rarity" json:"rarity"`
	UploaderID    int64     `bson:"uploader_id" json:"uploader_id"`
	UploaderName  string    `bson:"uploader_name" json:"uploader_name"`
	TelegramFileID string   `bson:"telegram_file_id,omitempty" json:"telegram_file_id,omitempty"`
	CreatedAt     time.Time `bson:"created_at" json:"created_at"`
}

type User struct {
	ID        int64     `bson:"_id" json:"_id"`
	Username  string    `bson:"username,omitempty" json:"username,omitempty"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

type CollectionItem struct {
	UserID      int64     `bson:"user_id" json:"user_id"`
	CharacterID string    `bson:"character_id" json:"character_id"`
	AddedAt     time.Time `bson:"added_at" json:"added_at"`
}

type UserTotals struct {
	ChatID          string `bson:"chat_id"`
	MessageFrequency int   `bson:"message_frequency"`
}
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AnimeCharacter represents a character in the database
type AnimeCharacter struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string            `bson:"name" json:"name"`
	Anime       string            `bson:"anime" json:"anime"`
	Rarity      string            `bson:"rarity" json:"rarity"` // Common, Rare, Epic, Legendary, etc.
	ImageURL    string            `bson:"image_url" json:"image_url"`
	Description string            `bson:"description,omitempty" json:"description,omitempty"`
	CreatedAt   time.Time         `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time         `bson:"updated_at" json:"updated_at"`
}

// UserTotal represents user's total statistics
type UserTotal struct {
	UserID         int64     `bson:"user_id" json:"user_id"`
	Username       string    `bson:"username" json:"username"`
	FirstName      string    `bson:"first_name" json:"first_name"`
	TotalCharacters int      `bson:"total_characters" json:"total_characters"`
	LastUpdated    time.Time `bson:"last_updated" json:"last_updated"`
}

// UserCharacter represents a character owned by a user
type UserCharacter struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      int64             `bson:"user_id" json:"user_id"`
	CharacterID primitive.ObjectID `bson:"character_id" json:"character_id"`
	Character   *AnimeCharacter    `bson:"character,omitempty" json:"character,omitempty"` // Populated field
	Count       int               `bson:"count" json:"count"` // How many of this character user has
	FirstCaught time.Time         `bson:"first_caught" json:"first_caught"`
	LastCaught  time.Time         `bson:"last_caught" json:"last_caught"`
}

// GroupUserTotal represents user statistics in a specific group
type GroupUserTotal struct {
	UserID         int64     `bson:"user_id" json:"user_id"`
	GroupID        int64     `bson:"group_id" json:"group_id"`
	Username       string    `bson:"username" json:"username"`
	FirstName      string    `bson:"first_name" json:"first_name"`
	TotalCharacters int      `bson:"total_characters" json:"total_characters"`
	LastUpdated    time.Time `bson:"last_updated" json:"last_updated"`
}

// TopGlobalGroup represents top groups globally
type TopGlobalGroup struct {
	GroupID     int64     `bson:"group_id" json:"group_id"`
	GroupName   string    `bson:"group_name" json:"group_name"`
	TotalUsers  int       `bson:"total_users" json:"total_users"`
	TotalCatches int      `bson:"total_catches" json:"total_catches"`
	LastUpdated time.Time `bson:"last_updated" json:"last_updated"`
}

// PMUser represents users who have interacted with bot in PM
type PMUser struct {
	UserID      int64     `bson:"user_id" json:"user_id"`
	Username    string    `bson:"username" json:"username"`
	FirstName   string    `bson:"first_name" json:"first_name"`
	LastName    string    `bson:"last_name,omitempty" json:"last_name,omitempty"`
	FirstSeen   time.Time `bson:"first_seen" json:"first_seen"`
	LastSeen    time.Time `bson:"last_seen" json:"last_seen"`
	IsBlocked   bool      `bson:"is_blocked" json:"is_blocked"`
}

// UserBalance represents user's balance/currency
type UserBalance struct {
	UserID    int64     `bson:"user_id" json:"user_id"`
	Balance   int64     `bson:"balance" json:"balance"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

// TradeRequest represents a trade between two users
type TradeRequest struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FromUserID      int64             `bson:"from_user_id" json:"from_user_id"`
	ToUserID        int64             `bson:"to_user_id" json:"to_user_id"`
	OfferedCharID   primitive.ObjectID `bson:"offered_char_id" json:"offered_char_id"`
	RequestedCharID primitive.ObjectID `bson:"requested_char_id" json:"requested_char_id"`
	Status          string            `bson:"status" json:"status"` // pending, accepted, rejected, cancelled
	CreatedAt       time.Time         `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time         `bson:"updated_at" json:"updated_at"`
}

// LeaderboardEntry represents a leaderboard entry
type LeaderboardEntry struct {
	Rank       int    `json:"rank"`
	UserID     int64  `json:"user_id"`
	Username   string `json:"username"`
	FirstName  string `json:"first_name"`
	Count      int    `json:"count"`
}

package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB wraps the MongoDB collections for easier access
type DB struct {
	Collections *Collections
	ctx         context.Context
}

// Collections holds all MongoDB collections
type Collections struct {
	AnimeCharacters      *mongo.Collection
	UserTotals          *mongo.Collection
	UserCollection      *mongo.Collection
	GroupUserTotals     *mongo.Collection
	TopGlobalGroups     *mongo.Collection
	PMUsers             *mongo.Collection
	UserBalance         *mongo.Collection
}

// NewDB creates a new database wrapper
func NewDB(database *mongo.Database) *DB {
	return &DB{
		Collections: &Collections{
			AnimeCharacters:  database.Collection("anime_characters_lol"),
			UserTotals:      database.Collection("user_totals_lmaoooo"),
			UserCollection:  database.Collection("user_collection_lmaoooo"),
			GroupUserTotals: database.Collection("group_user_totalsssssss"),
			TopGlobalGroups: database.Collection("top_global_groups"),
			PMUsers:         database.Collection("total_pm_users"),
			UserBalance:     database.Collection("user_balance"),
		},
		ctx: context.Background(),
	}
}

// ChangeBalance updates user's balance by a specified amount
// Amount can be positive (add) or negative (subtract)
// Creates the user document if it doesn't exist
func (db *DB) ChangeBalance(userID int64, amount int64) error {
	ctx, cancel := context.WithTimeout(db.ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"user_id": userID}
	update := bson.M{
		"$inc": bson.M{"balance": amount},
		"$setOnInsert": bson.M{
			"user_id":    userID,
			"created_at": time.Now(),
		},
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}
	opts := options.Update().SetUpsert(true)

	_, err := db.Collections.UserBalance.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return fmt.Errorf("failed to change balance: %w", err)
	}

	return nil
}

// GetBalance retrieves user's current balance
func (db *DB) GetBalance(userID int64) (int64, error) {
	ctx, cancel := context.WithTimeout(db.ctx, 5*time.Second)
	defer cancel()

	var result struct {
		Balance int64 `bson:"balance"`
	}

	filter := bson.M{"user_id": userID}
	err := db.Collections.UserBalance.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, nil
		}
		return 0, fmt.Errorf("failed to get balance: %w", err)
	}

	return result.Balance, nil
}

// SetBalance sets user's balance to a specific amount
func (db *DB) SetBalance(userID int64, amount int64) error {
	ctx, cancel := context.WithTimeout(db.ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"user_id": userID}
	update := bson.M{
		"$set": bson.M{
			"balance":    amount,
			"updated_at": time.Now(),
		},
		"$setOnInsert": bson.M{
			"user_id":    userID,
			"created_at": time.Now(),
		},
	}
	opts := options.Update().SetUpsert(true)

	_, err := db.Collections.UserBalance.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return fmt.Errorf("failed to set balance: %w", err)
	}

	return nil
}

// AddPMUser adds or updates a PM user
func (db *DB) AddPMUser(userID int64, username, firstName, lastName string) error {
	ctx, cancel := context.WithTimeout(db.ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"user_id": userID}
	update := bson.M{
		"$set": bson.M{
			"username":   username,
			"first_name": firstName,
			"last_name":  lastName,
			"last_seen":  time.Now(),
		},
		"$setOnInsert": bson.M{
			"user_id":    userID,
			"first_seen": time.Now(),
			"is_blocked": false,
		},
	}
	opts := options.Update().SetUpsert(true)

	_, err := db.Collections.PMUsers.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return fmt.Errorf("failed to add PM user: %w", err)
	}

	return nil
}

// UpdateUserTotal updates user's total character count
func (db *DB) UpdateUserTotal(userID int64, username, firstName string, count int) error {
	ctx, cancel := context.WithTimeout(db.ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"user_id": userID}
	update := bson.M{
		"$set": bson.M{
			"username":          username,
			"first_name":        firstName,
			"total_characters":  count,
			"last_updated":      time.Now(),
		},
	}
	opts := options.Update().SetUpsert(true)

	_, err := db.Collections.UserTotals.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return fmt.Errorf("failed to update user total: %w", err)
	}

	return nil
}

// IncrementUserTotal increments user's character count by specified amount
func (db *DB) IncrementUserTotal(userID int64, username, firstName string, increment int) error {
	ctx, cancel := context.WithTimeout(db.ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"user_id": userID}
	update := bson.M{
		"$inc": bson.M{"total_characters": increment},
		"$set": bson.M{
			"username":     username,
			"first_name":   firstName,
			"last_updated": time.Now(),
		},
	}
	opts := options.Update().SetUpsert(true)

	_, err := db.Collections.UserTotals.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return fmt.Errorf("failed to increment user total: %w", err)
	}

	return nil
}

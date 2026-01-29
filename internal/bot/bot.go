package bot

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Bot represents the main bot instance
type Bot struct {
	API         *tgbotapi.BotAPI
	DB          *mongo.Database
	Collections *Collections
	Context     context.Context
}

// Collections holds all MongoDB collections
type Collections struct {
	AnimeCharacters      *mongo.Collection
	UserTotals          *mongo.Collection
	UserCollection      *mongo.Collection
	GroupUserTotals     *mongo.Collection
	TopGlobalGroups     *mongo.Collection
	PMUsers             *mongo.Collection
	UserBalance         *mongo.Collection // New: User balance collection
}

// Logger configuration
var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	DebugLogger *log.Logger
)

// InitLogger initializes the logging system
func InitLogger() error {
	// Open log file
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	// Create loggers
	InfoLogger = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	DebugLogger = log.New(logFile, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Also log to stdout
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	return nil
}

// NewBot creates and initializes a new bot instance
func NewBot(token, mongoURL, dbName string) (*Bot, error) {
	// Initialize logger
	if err := InitLogger(); err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	InfoLogger.Println("Initializing bot...")

	// Create Telegram Bot API instance
	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot API: %w", err)
	}

	InfoLogger.Printf("Authorized on account: @%s", botAPI.Self.UserName)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURL)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Ping the database to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	InfoLogger.Println("Successfully connected to MongoDB")

	// Get database
	db := client.Database(dbName)

	// Initialize collections
	collections := &Collections{
		AnimeCharacters:  db.Collection("anime_characters_lol"),
		UserTotals:      db.Collection("user_totals_lmaoooo"),
		UserCollection:  db.Collection("user_collection_lmaoooo"),
		GroupUserTotals: db.Collection("group_user_totalsssssss"),
		TopGlobalGroups: db.Collection("top_global_groups"),
		PMUsers:         db.Collection("total_pm_users"),
		UserBalance:     db.Collection("user_balance"), // New collection
	}

	// Create bot instance
	bot := &Bot{
		API:         botAPI,
		DB:          db,
		Collections: collections,
		Context:     context.Background(),
	}

	InfoLogger.Println("Bot initialized successfully")

	return bot, nil
}

// Start begins polling for updates
func (b *Bot) Start() error {
	InfoLogger.Println("Starting bot polling...")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.API.GetUpdatesChan(u)

	for update := range updates {
		// Process update in a goroutine for concurrent handling
		go b.HandleUpdate(update)
	}

	return nil
}

// HandleUpdate processes incoming updates
func (b *Bot) HandleUpdate(update tgbotapi.Update) {
	// This will be implemented with handlers in next step
	if update.Message != nil {
		InfoLogger.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	}
}

// Stop gracefully shuts down the bot
func (b *Bot) Stop() error {
	InfoLogger.Println("Stopping bot...")

	// Disconnect from MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := b.DB.Client().Disconnect(ctx); err != nil {
		return fmt.Errorf("failed to disconnect from MongoDB: %w", err)
	}

	b.API.StopReceivingUpdates()

	InfoLogger.Println("Bot stopped successfully")
	return nil
}

// SendMessage sends a message to a chat
func (b *Bot) SendMessage(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := b.API.Send(msg)
	return err
}

// SendPhoto sends a photo to a chat
func (b *Bot) SendPhoto(chatID int64, photoURL string, caption string) error {
	msg := tgbotapi.NewPhoto(chatID, tgbotapi.FileURL(photoURL))
	msg.Caption = caption
	_, err := b.API.Send(msg)
	return err
}

// SendVideo sends a video to a chat
func (b *Bot) SendVideo(chatID int64, videoURL string, caption string) error {
	msg := tgbotapi.NewVideo(chatID, tgbotapi.FileURL(videoURL))
	msg.Caption = caption
	_, err := b.API.Send(msg)
	return err
}

// ReplyToMessage replies to a specific message
func (b *Bot) ReplyToMessage(chatID int64, replyToMessageID int, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyToMessageID = replyToMessageID
	_, err := b.API.Send(msg)
	return err
}

// ChangeBalance updates user's balance by a specified amount
// Amount can be positive (add) or negative (subtract)
func (b *Bot) ChangeBalance(userID int64, amount int64) error {
	ctx, cancel := context.WithTimeout(b.Context, 5*time.Second)
	defer cancel()

	// Use upsert to create document if it doesn't exist
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

	_, err := b.Collections.UserBalance.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		ErrorLogger.Printf("Failed to change balance for user %d: %v", userID, err)
		return fmt.Errorf("failed to update balance: %w", err)
	}

	InfoLogger.Printf("Balance changed for user %d by %d", userID, amount)
	return nil
}

// GetBalance retrieves user's current balance
func (b *Bot) GetBalance(userID int64) (int64, error) {
	ctx, cancel := context.WithTimeout(b.Context, 5*time.Second)
	defer cancel()

	var result struct {
		Balance int64 `bson:"balance"`
	}

	filter := bson.M{"user_id": userID}
	err := b.Collections.UserBalance.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// User doesn't exist, return 0 balance
			return 0, nil
		}
		return 0, fmt.Errorf("failed to get balance: %w", err)
	}

	return result.Balance, nil
}

// CreateBackgroundTask creates a background task (goroutine)
func (b *Bot) CreateBackgroundTask(task func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				ErrorLogger.Printf("Background task panic recovered: %v", r)
			}
		}()
		task()
	}()
}

// GetUserProfilePhotos gets user profile photos
func (b *Bot) GetUserProfilePhotos(userID int64) (tgbotapi.UserProfilePhotos, error) {
	photos := tgbotapi.NewUserProfilePhotos(userID)
	return b.API.GetUserProfilePhotos(photos)
}

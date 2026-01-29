package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	// Telegram Bot Configuration
	BotToken    string
	BotUsername string
	OwnerID     int64
	SudoUsers   []int64

	// Telegram API Configuration (for Pyrogram equivalent)
	ApiID   int
	ApiHash string

	// Group/Channel IDs
	GroupID       int64
	CharaChannelID int64
	SupportChat   string
	UpdateChat    string

	// MongoDB Configuration
	MongoURL     string
	DatabaseName string

	// Photo URLs
	PhotoURLs []string

	// Logging
	LogLevel string
}

var AppConfig *Config

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if it exists (for local development)
	_ = godotenv.Load()

	cfg := &Config{
		BotToken:       getEnv("BOT_TOKEN", "6707490163:AAHZzqjm3rbEZsObRiNaT7DMtw_i5WPo_0o"),
		BotUsername:    getEnv("BOT_USERNAME", "Collect_Em_AllBot"),
		MongoURL:       getEnv("MONGO_URI", "mongodb+srv://HaremDBBot:ThisIsPasswordForHaremDB@haremdb.swzjngj.mongodb.net/?retryWrites=true&w=majority"),
		DatabaseName:   getEnv("DATABASE_NAME", "Character_catcher"),
		LogLevel:       getEnv("LOG_LEVEL", "info"),
		SupportChat:    getEnv("SUPPORT_CHAT", "Collect_em_support"),
		UpdateChat:     getEnv("UPDATE_CHAT", "Collect_em_support"),
	}

	// Parse Owner ID
	ownerIDStr := getEnv("OWNER_ID", "6765826972")
	ownerID, err := strconv.ParseInt(ownerIDStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid OWNER_ID: %w", err)
	}
	cfg.OwnerID = ownerID

	// Parse Sudo Users
	sudoUsersStr := getEnv("SUDO_USERS", "6845325416,6765826972")
	sudoUsers, err := parseIntSlice(sudoUsersStr)
	if err != nil {
		return nil, fmt.Errorf("invalid SUDO_USERS: %w", err)
	}
	cfg.SudoUsers = sudoUsers

	// Parse Group ID
	groupIDStr := getEnv("GROUP_ID", "-1002133191051")
	groupID, err := strconv.ParseInt(groupIDStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid GROUP_ID: %w", err)
	}
	cfg.GroupID = groupID

	// Parse Chara Channel ID
	charaChannelIDStr := getEnv("CHARA_CHANNEL_ID", "-1002133191051")
	charaChannelID, err := strconv.ParseInt(charaChannelIDStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid CHARA_CHANNEL_ID: %w", err)
	}
	cfg.CharaChannelID = charaChannelID

	// Parse API ID
	apiIDStr := getEnv("API_ID", "26626068")
	apiID, err := strconv.Atoi(apiIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid API_ID: %w", err)
	}
	cfg.ApiID = apiID

	// API Hash
	cfg.ApiHash = getEnv("API_HASH", "bf423698bcbe33cfd58b11c78c42caa2")

	// Parse Photo URLs
	photoURLsStr := getEnv("PHOTO_URLS", "https://telegra.ph/file/b925c3985f0f325e62e17.jpg,https://telegra.ph/file/4211fb191383d895dab9d.jpg")
	cfg.PhotoURLs = strings.Split(photoURLsStr, ",")

	// Validate required fields
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	AppConfig = cfg
	return cfg, nil
}

// Validate checks if all required configuration is present
func (c *Config) Validate() error {
	if c.BotToken == "" {
		return fmt.Errorf("BOT_TOKEN is required")
	}
	if c.MongoURL == "" {
		return fmt.Errorf("MONGO_URI is required")
	}
	if c.OwnerID == 0 {
		return fmt.Errorf("OWNER_ID is required")
	}
	return nil
}

// IsSudoUser checks if a user ID is in the sudo users list
func (c *Config) IsSudoUser(userID int64) bool {
	for _, id := range c.SudoUsers {
		if id == userID {
			return true
		}
	}
	return userID == c.OwnerID
}

// GetRandomPhotoURL returns a random photo URL from the list
func (c *Config) GetRandomPhotoURL() string {
	if len(c.PhotoURLs) == 0 {
		return ""
	}
	// For now, return first one. You can add random logic later
	return c.PhotoURLs[0]
}

// Helper function to get environment variable with default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// Helper function to parse comma-separated integers
func parseIntSlice(s string) ([]int64, error) {
	if s == "" {
		return []int64{}, nil
	}

	parts := strings.Split(s, ",")
	result := make([]int64, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		num, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}

	return result, nil
}

// MustLoadConfig loads config or panics
func MustLoadConfig() *Config {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	return cfg
}

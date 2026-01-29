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
	// Logging
	Logger bool

	// Bot Credentials
	Token       string
	BotUsername string

	// Telegram API Credentials
	ApiID   int
	ApiHash string

	// Owner and Sudo Users
	OwnerID   int64
	SudoUsers []int64

	// Group and Channel IDs
	GroupID       int64
	CharaChannelID int64

	// Database
	MongoURL     string
	DatabaseName string

	// Media
	VideoURLs []string

	// Community Links
	SupportChat string
	UpdateChat  string
}

var AppConfig *Config

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if it exists (for local development)
	_ = godotenv.Load()

	cfg := &Config{
		Logger:       true,
		Token:        getEnv("BOT_TOKEN", "8551975632:AAH1vrphQvEf_O5w9IfecwUmHJ_QwQbgBwM"),
		BotUsername:  getEnv("BOT_USERNAME", "Senpai_Waifu_Grabbing_Bot"),
		MongoURL:     getEnv("MONGO_URL", "mongodb+srv://ravi:ravi12345@cluster0.hndinhj.mongodb.net/?retryWrites=true&w=majority"),
		DatabaseName: getEnv("DATABASE_NAME", "Character_catcher"),
		SupportChat:  getEnv("SUPPORT_CHAT", "THE_DRAGON_SUPPORT"),
		UpdateChat:   getEnv("UPDATE_CHAT", "PICK_X_UPDATE"),
	}

	// Parse API ID
	apiIDStr := getEnv("API_ID", "35660683")
	apiID, err := strconv.Atoi(apiIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid API_ID: %w", err)
	}
	cfg.ApiID = apiID

	// API Hash
	cfg.ApiHash = getEnv("API_HASH", "7afb42cd73fb5f3501062ffa6a1f87f7")

	// Parse Owner ID
	ownerIDStr := getEnv("OWNER_ID", "7818323042")
	ownerID, err := strconv.ParseInt(ownerIDStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid OWNER_ID: %w", err)
	}
	cfg.OwnerID = ownerID

	// Parse Sudo Users
	sudoUsersStr := getEnv("SUDO_USERS", "7818323042,8453236527")
	sudoUsers, err := parseIntSlice(sudoUsersStr)
	if err != nil {
		return nil, fmt.Errorf("invalid SUDO_USERS: %w", err)
	}
	cfg.SudoUsers = sudoUsers

	// Add OWNER_ID to SUDO_USERS if not already present
	if !contains(cfg.SudoUsers, cfg.OwnerID) {
		cfg.SudoUsers = append(cfg.SudoUsers, cfg.OwnerID)
	}

	// Parse Group ID
	groupIDStr := getEnv("GROUP_ID", "-1003129952280")
	groupID, err := strconv.ParseInt(groupIDStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid GROUP_ID: %w", err)
	}
	cfg.GroupID = groupID

	// Parse Chara Channel ID
	charaChannelIDStr := getEnv("CHARA_CHANNEL_ID", "-1003150808065")
	charaChannelID, err := strconv.ParseInt(charaChannelIDStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid CHARA_CHANNEL_ID: %w", err)
	}
	cfg.CharaChannelID = charaChannelID

	// Parse Video URLs
	videoURLsStr := getEnv("VIDEO_URL", "https://files.catbox.moe/iqeaeb.mp4,https://files.catbox.moe/fp7m2d.mp4,https://files.catbox.moe/cv8r9i.mp4,https://files.catbox.moe/kz2usa.mp4,https://files.catbox.moe/u3gfz5.mp4,https://files.catbox.moe/4w63xt.mp4,https://files.catbox.moe/3mv64w.mp4,https://files.catbox.moe/n2m9av.mp4,https://files.catbox.moe/lrjr1o.mp4,https://files.catbox.moe/xdmuzm.mp4,https://files.catbox.moe/lqsdnr.mp4,https://files.catbox.moe/3mv64w.mp4")
	cfg.VideoURLs = parseStringSlice(videoURLsStr)

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	AppConfig = cfg
	return cfg, nil
}

// Validate checks if all required configuration is present
func (c *Config) Validate() error {
	errors := []string{}

	if c.Token == "" {
		errors = append(errors, "BOT_TOKEN is required")
	}

	if c.ApiID == 0 {
		errors = append(errors, "API_ID is required")
	}

	if c.ApiHash == "" {
		errors = append(errors, "API_HASH is required")
	}

	if c.OwnerID == 0 {
		errors = append(errors, "OWNER_ID is required")
	}

	if c.MongoURL == "" {
		errors = append(errors, "MONGO_URL is required")
	}

	if c.GroupID == 0 {
		errors = append(errors, "GROUP_ID is required")
	}

	if c.CharaChannelID == 0 {
		errors = append(errors, "CHARA_CHANNEL_ID is required")
	}

	if len(errors) > 0 {
		fmt.Println("❌ Configuration Error(s):")
		for _, err := range errors {
			fmt.Printf("   - %s\n", err)
		}
		fmt.Println("\n💡 Please set the required environment variables and try again.")
		return fmt.Errorf("configuration validation failed")
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

// GetRandomVideoURL returns a random video URL from the list
func (c *Config) GetRandomVideoURL() string {
	if len(c.VideoURLs) == 0 {
		return ""
	}
	// You can add random logic later
	return c.VideoURLs[0]
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

// Helper function to parse comma-separated strings
func parseStringSlice(s string) []string {
	if s == "" {
		return []string{}
	}

	parts := strings.Split(s, ",")
	result := make([]string, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			result = append(result, part)
		}
	}

	return result
}

// Helper function to check if slice contains value
func contains(slice []int64, value int64) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// MustLoadConfig loads config or panics
func MustLoadConfig() *Config {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	return cfg
}

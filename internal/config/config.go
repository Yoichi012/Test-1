package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken     string
	MongoURI          string
	MongoDB           string
	SpawnIntervalMins int
	OwnerID           int64
	SudoUsers         []int64
	GroupID           int64
	CharaChannelID    int64
}

// Load reads env variables (uses .env automatically if present)
func Load() (*Config, error) {
	_ = godotenv.Load()

	c := &Config{
		TelegramToken: os.Getenv("TELEGRAM_TOKEN"),
		MongoURI:      os.Getenv("MONGO_URI"),
		MongoDB:       os.Getenv("MONGO_DB"),
	}

	if c.MongoDB == "" {
		c.MongoDB = "waifu_db"
	}

	minStr := os.Getenv("SPAWN_INTERVAL_MINUTES")
	if minStr == "" {
		c.SpawnIntervalMins = 60
	} else {
		if v, err := strconv.Atoi(minStr); err == nil {
			c.SpawnIntervalMins = v
		} else {
			c.SpawnIntervalMins = 60
		}
	}

	if owner := os.Getenv("OWNER_ID"); owner != "" {
		if v, err := strconv.ParseInt(owner, 10, 64); err == nil {
			c.OwnerID = v
		}
	}

	if s := os.Getenv("SUDO_USERS"); s != "" {
		parts := strings.Split(s, ",")
		for _, p := range parts {
			if v, err := strconv.ParseInt(strings.TrimSpace(p), 10, 64); err == nil {
				c.SudoUsers = append(c.SudoUsers, v)
			}
		}
	}

	if g := os.Getenv("GROUP_ID"); g != "" {
		if v, err := strconv.ParseInt(g, 10, 64); err == nil {
			c.GroupID = v
		}
	}
	if ch := os.Getenv("CHARA_CHANNEL_ID"); ch != "" {
		if v, err := strconv.ParseInt(ch, 10, 64); err == nil {
			c.CharaChannelID = v
		}
	}

	return c, nil
}
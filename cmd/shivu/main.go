package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/yourusername/shivu-go/internal/bot"
	"github.com/yourusername/shivu-go/internal/config"
)

func main() {
	// Print banner
	printBanner()

	// Load configuration
	log.Println("📋 Loading configuration...")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("❌ Failed to load configuration: %v", err)
	}
	log.Println("✅ Configuration loaded successfully")

	// Display configuration summary
	log.Printf("🤖 Bot: @%s", cfg.BotUsername)
	log.Printf("👤 Owner ID: %d", cfg.OwnerID)
	log.Printf("👥 Sudo Users: %v", cfg.SudoUsers)
	log.Printf("💾 Database: %s", cfg.DatabaseName)

	// Initialize bot
	log.Println("🚀 Initializing bot...")
	botInstance, err := bot.NewBot(cfg.Token, cfg.MongoURL, cfg.DatabaseName)
	if err != nil {
		log.Fatalf("❌ Failed to initialize bot: %v", err)
	}
	log.Printf("✅ Bot initialized: @%s", cfg.BotUsername)

	// Setup graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start bot in a goroutine
	go func() {
		log.Println("▶️  Starting bot...")
		if err := botInstance.Start(); err != nil {
			log.Fatalf("❌ Bot stopped with error: %v", err)
		}
	}()

	log.Println("✨ Bot is running. Press CTRL+C to stop.")
	log.Printf("💬 Support: @%s", cfg.SupportChat)
	log.Printf("📢 Updates: @%s", cfg.UpdateChat)

	// Wait for shutdown signal
	<-quit
	log.Println("\n⏸️  Received shutdown signal...")

	// Graceful shutdown
	if err := botInstance.Stop(); err != nil {
		log.Printf("⚠️  Error during shutdown: %v", err)
	}

	log.Println("👋 Shutdown complete. Goodbye!")
}

func printBanner() {
	banner := `
╔═══════════════════════════════════════════════════════════╗
║                                                           ║
║           🌸 SHIVU BOT - GO VERSION 🌸                   ║
║                                                           ║
║        Character Collection Telegram Bot                 ║
║                                                           ║
║        Powered by Go | MongoDB | Telegram API            ║
║                                                           ║
╚═══════════════════════════════════════════════════════════╝
	`
	log.Println(banner)
}

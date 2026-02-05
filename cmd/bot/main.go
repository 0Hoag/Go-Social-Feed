package main

import (
	"log"
	"os"

	"github.com/hoag/go-social-feed/internal/adapters/etherscan"
	"github.com/hoag/go-social-feed/internal/adapters/gemini"
	"github.com/hoag/go-social-feed/internal/adapters/telegram"
	"github.com/hoag/go-social-feed/internal/core/scanner"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Load Config
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}

	ethKey := os.Getenv("ETHERSCAN_API_KEY")
	bscKey := os.Getenv("BSCSCAN_API_KEY")
	baseKey := os.Getenv("BASESCAN_API_KEY")
	arbKey := os.Getenv("ARBISCAN_API_KEY")
	polyKey := os.Getenv("POLYGONSCAN_API_KEY")
	botToken := os.Getenv("SCANNER_BOT_TOKEN")
	geminiKey := os.Getenv("GEMINI_API_KEY")

	if botToken == "" {
		log.Fatal("❌ Error: SCANNER_BOT_TOKEN is required in .env")
	}

	log.Println("🚀 Starting ChainGuardian Scanner Bot...")

	// 2. Init Dependencies
	apiKeys := map[string]string{
		etherscan.NetworkETH:      ethKey,
		etherscan.NetworkBSC:      bscKey,
		etherscan.NetworkBase:     baseKey,
		etherscan.NetworkArbitrum: arbKey,
		etherscan.NetworkPolygon:  polyKey,
	}
	ethClient := etherscan.NewClient(apiKeys)

	var geminiClient *gemini.Client
	if geminiKey != "" {
		geminiClient = gemini.NewClient(geminiKey)
		log.Println("✅ Gemini AI Integration: ENABLED")
	} else {
		log.Println("⚠️ Gemini AI Integration: DISABLED (Using Regex Only)")
	}

	scanEngine := scanner.NewEngine(geminiClient)

	// 3. Init Bot
	bot, err := telegram.NewScannerBot(botToken, ethClient, scanEngine)
	if err != nil {
		log.Fatalf("❌ Failed to init bot: %v", err)
	}

	// 4. Start Listener
	log.Println("✅ Bot is online and listening for addresses...")
	bot.Start()
}

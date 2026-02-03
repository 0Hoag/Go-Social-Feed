package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hoag/go-social-feed/internal/adapters/etherscan"
	"github.com/hoag/go-social-feed/internal/core/scanner"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found (ETHERSCAN_API_KEY might be missing)")
	}

	apiKey := os.Getenv("ETHERSCAN_API_KEY")
	if apiKey == "" {
		log.Fatal("Error: ETHERSCAN_API_KEY is not set")
	}

	// Parse flags
	contractAddr := flag.String("addr", "", "Contract Address to scan (Ethereum Mainnet)")
	flag.Parse()

	if *contractAddr == "" {
		fmt.Println("Usage: go run cmd/scanner/main.go --addr <0x...>")
		os.Exit(1)
	}

	fmt.Printf("🔍 Scanning Contract: %s\n", *contractAddr)

	// 1. Fetch Source Code
	client := etherscan.NewClient(apiKey)
	source, err := client.GetContractSource(*contractAddr)
	if err != nil {
		log.Fatalf("❌ Failed to fetch source code: %v", err)
	}

	if source == "" {
		log.Fatalf("❌ No source code found (Contract might not be verified)")
	}

	fmt.Printf("✅ Source code fetched (%d bytes)\n", len(source))

	// 2. Scan Logic
	engine := scanner.NewEngine()
	result := engine.Scan(source)

	// 3. Print Results
	fmt.Println("\n📊 === SCAN RESULT ===")
	fmt.Printf("Trust Score: %d/100\n", result.TrustScore)

	if len(result.Issues) == 0 {
		fmt.Println("✅ No critical issues detected (Basic Scan)")
	} else {
		fmt.Println("\n⚠️  Risk Factors detected:")
		for _, issue := range result.Issues {
			fmt.Printf("- [%s] %s: %s (-%d)\n", issue.Type, issue.Name, issue.Description, issue.Impact)
		}
	}
}

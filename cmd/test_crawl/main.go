package main

import (
	"context"
	"fmt"

	"github.com/hoag/go-social-feed/internal/crawler"
	"github.com/hoag/go-social-feed/internal/crawler/sites"
	"github.com/hoag/go-social-feed/internal/processor"
	"github.com/hoag/go-social-feed/pkg/log"
)

func main() {
	l := log.InitializeTestZapLogger()

	// 1. Crawl
	manager := crawler.NewManager(l)
	manager.Register(sites.NewCoindeskCrawler())

	fmt.Println("Starting crawl...")
	articles, err := manager.Run(context.Background())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// 2. Process
	proc := processor.NewSimpleProcessor(l)

	fmt.Printf("Crawled %d articles. Showing top 5 with translation:\n", len(articles))
	for i, a := range articles {
		if i >= 5 {
			break
		}

		fmt.Printf("--------------------------------------------------\n")
		fmt.Printf("[ORIGINAL] %s\n", a.Title)

		res, err := proc.Process(context.Background(), a)
		if err != nil {
			fmt.Printf("[ERROR] %v\n", err)
		} else {
			fmt.Printf("[TRANSLATED] %s\n", res.TranslatedTitle)
		}
		fmt.Printf("[LINK] %s\n", a.SourceURL)
	}
}

package crawler

import (
	"context"
	"time"
)

// Article represents a raw crawled article
type Article struct {
	Title       string
	Summary     string
	Content     string
	SourceURL   string
	ImageURL    string
	Source      string // e.g., "coindesk", "cointelegraph"
	PublishedAt time.Time
	CrawledAt   time.Time
}

// SiteCrawler defines the interface for a specific website crawler
type SiteCrawler interface {
	// Name returns the unique name of the crawler (e.g., "coindesk")
	Name() string
	// Crawl fetches the latest articles from the site
	Crawl(ctx context.Context) ([]Article, error)
}

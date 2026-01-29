package sites

import (
	"context"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/hoag/go-social-feed/internal/crawler"
)

type CoindeskCrawler struct {
}

func NewCoindeskCrawler() *CoindeskCrawler {
	return &CoindeskCrawler{}
}

func (c *CoindeskCrawler) Name() string {
	return "coindesk"
}

func (c *CoindeskCrawler) Crawl(ctx context.Context) ([]crawler.Article, error) {
	var articles []crawler.Article

	// Create collector by libary Colly
	collector := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36"),
	)

	// Coindesk structure often changes. We target common article wrappers.
	// This selector targets the main news texts.
	// Note: We might need to refine this based on the actual HTML structure.
	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		title := strings.TrimSpace(e.Text)

		// Filter relevant links
		if title == "" || len(strings.Split(title, " ")) < 3 { // Skip short link texts
			return
		}

		// Ensure full URL
		if !strings.HasPrefix(link, "http") {
			link = "https://www.coindesk.com" + link
		}

		// Basic filter to ensure it looks like an article
		if strings.Contains(link, "/business/") ||
			strings.Contains(link, "/markets/") ||
			strings.Contains(link, "/policy/") ||
			strings.Contains(link, "/tech/") ||
			strings.Contains(link, "/opinion/") {

			// Deduplicate if already added (simple check)
			for _, a := range articles {
				if a.SourceURL == link {
					return
				}
			}

			// Create detail collector to fetch image and full content
			detailCollector := collector.Clone()
			var imageURL string
			var fullContent strings.Builder

			detailCollector.OnHTML("meta[property='og:image']", func(e *colly.HTMLElement) {
				imageURL = e.Attr("content")
			})

			// Extract article body content
			// CoinDesk uses various selectors, try multiple
			detailCollector.OnHTML("article p, .article-body p, [data-module-name='article-body'] p", func(e *colly.HTMLElement) {
				text := strings.TrimSpace(e.Text)
				if text != "" && len(text) > 20 { // Filter out short/empty paragraphs
					fullContent.WriteString(text)
					fullContent.WriteString("\n\n")
				}
			})

			detailCollector.Visit(link)

			articles = append(articles, crawler.Article{
				Title:       title,
				SourceURL:   link,
				ImageURL:    imageURL,
				Content:     strings.TrimSpace(fullContent.String()), // Full article text
				Source:      "coindesk",
				CrawledAt:   time.Now(),
				PublishedAt: time.Now(),
			})
		}
	})

	err := collector.Visit("https://www.coindesk.com/")
	if err != nil {
		return nil, crawler.ErrCrawlFailed
	}

	return articles, nil
}

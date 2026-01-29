package sites

import (
	"context"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/hoag/go-social-feed/internal/crawler"
)

type coiTelegraphCrawler struct {
	c *colly.Collector
}

func NewCoinTelegraphCrawler() crawler.SiteCrawler {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36"),
	)
	return &coiTelegraphCrawler{
		c: c,
	}
}

func (c *coiTelegraphCrawler) Name() string {
	return "cointelegraph"
}

func (c *coiTelegraphCrawler) Crawl(ctx context.Context) ([]crawler.Article, error) {
	var articles []crawler.Article

	// Use RSS Feed
	c.c.OnXML("//item", func(e *colly.XMLElement) {
		title := e.ChildText("title")
		link := e.ChildText("link")
		summary := e.ChildText("description") // Use description as summary

		// Image might be in media:content or embedded in description
		// We use local-name() to be namespace agnostic or just simple filtering
		imageURL := e.ChildAttr("*[name()='media:content']", "url")
		if imageURL == "" {
			// Try extracting from description if it contains HTML image
			// But for now let's rely on media:content which is standard in CT RSS
		}

		// Published Date
		pubDate := e.ChildText("pubDate")
		// Parse pubDate if needed, but for now we use time.Now() approximation or try parsing
		// Mon, 02 Jan 2006 15:04:05 MST
		publishedAt, err := time.Parse(time.RFC1123, pubDate)
		if err != nil {
			publishedAt = time.Now()
		}

		if title != "" && link != "" {
			articles = append(articles, crawler.Article{
				Title:       strings.TrimSpace(title),
				Summary:     strings.TrimSpace(summary),
				SourceURL:   strings.TrimSpace(link),
				ImageURL:    imageURL,
				Source:      "cointelegraph",
				CrawledAt:   time.Now(),
				PublishedAt: publishedAt,
			})
		}
	})

	err := c.c.Visit("https://cointelegraph.com/rss")
	if err != nil {
		return nil, crawler.ErrCrawlFailed
	}

	c.c.Wait()
	return articles, nil
}

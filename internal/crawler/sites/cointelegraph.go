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
	// Add Headers to look like a real browser
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.9")
		r.Headers.Set("Referer", "https://google.com")
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
	})
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

		// Visit article page to extract full content
		detailCollector := c.c.Clone()
		var fullContent strings.Builder

		detailCollector.OnHTML(".post-content p, article p, .article__content p", func(e *colly.HTMLElement) {
			text := strings.TrimSpace(e.Text)
			if text != "" && len(text) > 20 {
				fullContent.WriteString(text)
				fullContent.WriteString("\n\n")
			}
		})

		// Visit the article page
		if link != "" {
			detailCollector.Visit(link)
		}

		if title != "" && link != "" {
			articles = append(articles, crawler.Article{
				Title:       strings.TrimSpace(title),
				Summary:     strings.TrimSpace(summary),
				SourceURL:   strings.TrimSpace(link),
				ImageURL:    imageURL,
				Content:     strings.TrimSpace(fullContent.String()), // Full article text
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

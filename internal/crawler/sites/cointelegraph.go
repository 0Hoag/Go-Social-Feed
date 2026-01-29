package sites

import (
	"context"

	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/hoag/go-social-feed/internal/crawler"
	"github.com/hoag/go-social-feed/pkg/log"
)

type coiTelegraphCrawler struct {
	c *colly.Collector
	l log.Logger
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

	// CoinTelegraph RSS Feed
	// We use RSS because it's cleaner for CoinTelegraph
	// RSS URL: https://cointelegraph.com/rss

	// However, RSS usually doesn't have the full image or full content.
	// But let's start with RSS for list and then maybe visit detail if needed.
	// Actually, for CoinTelegraph, parsing the main page or category pages might be better for images.
	// Let's try parsing the main tags page or latest news.

	// Better approach for consistency: Scrape the "Latest News" section on their site.
	// URL: https://cointelegraph.com/tags/bitcoin (or just potentially the rss)

	// Let's use the RSS feed for simplicity and reliability of finding new links,
	// then we can optionally visit the link to get og:image if RSS doesn't have it.

	// Wait, colly xml handling is different. Let's stick to HTML scraping of the "Latest News" list if possible,
	// or just use RSS with standard XML parsing.

	// Let's try standard HTML scraping of the main page or a category page to be consistent with Coindesk.
	// https://cointelegraph.com/

	c.c.OnHTML("li.posts-listing__item", func(e *colly.HTMLElement) {
		// This selector might need adjustment based on actual CT layout.
		// CT layout changes often.

		// Let's rely on a more stable RSS approach for CoinTelegraph if we want 100% success,
		// but since we are using Colly, let's try to find article blocks.

		// Current CT layout (approx):
		// <li class="posts-listing__item">
		//    <article class="post-card-inline">
		//       <a href="...">Title</a>
		//       ...
		//    </article>
		// </li>

		title := e.ChildText("span.post-card-inline__title")
		link := e.ChildAttr("a.post-card-inline__figure-link", "href")
		summary := e.ChildText("p.post-card-inline__text")

		// Image is often lazy loaded or in a srcset.
		// Try finding standard img tag.
		imageURL := e.ChildAttr("img.post-card-inline__cover-img", "src")

		// Normalization
		if !strings.HasPrefix(link, "http") {
			link = "https://cointelegraph.com" + link
		}

		if title != "" && link != "" {
			articles = append(articles, crawler.Article{
				Title:       strings.TrimSpace(title),
				Summary:     strings.TrimSpace(summary),
				SourceURL:   link,
				ImageURL:    imageURL, // Might need detail fetch
				Source:      "cointelegraph",
				CrawledAt:   time.Now(),
				PublishedAt: time.Now(), // Approx
			})
		}
	})

	// RSS Approach with Colly is trickier as it expects HTML.
	// Let's try scraping the RSS XML as text and parse it? No, that's overcomplicating.
	// Let's stick to HTML.

	// Alternative: CoinTelegraph "Latest" page
	// https://cointelegraph.com/category/latest-news
	// Selectors: ul.posts-listing > li.posts-listing__item

	err := c.c.Visit("https://cointelegraph.com/category/latest-news")
	if err != nil {
		c.l.Errorf(ctx, "CoinTelegraph: Crawl failed: %v", err)
		return nil, crawler.ErrCrawlFailed
	}

	c.c.Wait()

	// If image is missing, we might want to visit the detail page, similar to Coindesk
	for i := range articles {
		if articles[i].ImageURL == "" || strings.Contains(articles[i].ImageURL, "base64") {
			detailCollector := c.c.Clone()
			detailCollector.OnHTML("meta[property='og:image']", func(e *colly.HTMLElement) {
				articles[i].ImageURL = e.Attr("content")
			})
			_ = detailCollector.Visit(articles[i].SourceURL)
		}
	}

	return articles, nil
}

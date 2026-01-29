package crawler

import (
	"context"
	"sync"
	"time"

	"github.com/hoag/go-social-feed/pkg/log"
)

// Manager handles multiple crawlers
type Manager struct {
	crawlers []SiteCrawler
	l        log.Logger
}

func NewManager(l log.Logger) *Manager {
	return &Manager{
		crawlers: make([]SiteCrawler, 0),
		l:        l,
	}
}

func (m *Manager) Register(c SiteCrawler) {
	m.crawlers = append(m.crawlers, c)
}

func (m *Manager) Run(ctx context.Context) ([]Article, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var allArticles []Article

	for _, c := range m.crawlers {
		wg.Add(1)
		go func(crawler SiteCrawler) {
			defer wg.Done()

			// Set timeout for each crawler
			ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
			defer cancel()

			articles, err := crawler.Crawl(ctx)
			if err != nil {
				m.l.Errorf(ctx, "Crawler %s failed: %v", crawler.Name(), err)
				return
			}
			m.l.Infof(ctx, "Crawler %s fetched %d articles", crawler.Name(), len(articles))

			mu.Lock()
			allArticles = append(allArticles, articles...)
			mu.Unlock()
		}(c)
	}

	wg.Wait()
	return allArticles, nil
}

package processor

import (
	"context"

	"github.com/hoag/go-social-feed/internal/crawler"
)

type ProcessedContent struct {
	OriginalTitle     string
	TranslatedTitle   string
	OriginalSummary   string
	TranslatedSummary string
	Content           string
	SourceURL         string
	ImageURL          string
}

type ContentProcessor interface {
	Process(ctx context.Context, article crawler.Article) (ProcessedContent, error)
}

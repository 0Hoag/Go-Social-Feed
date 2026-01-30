package processor

import (
	"context"

	"github.com/hoag/go-social-feed/internal/crawler"
)

type ProcessedContent struct {
	OriginalTitle         string
	TranslatedTitle       string // Short, concise title (main idea)
	OriginalSummary       string
	TranslatedSummary     string // Brief summary for feed (2-3 sentences)
	TranslatedFullContent string // Full article translated to Vietnamese
	Content               string
	SourceURL             string
	ImageURL              string
}

type ContentProcessor interface {
	Process(ctx context.Context, article crawler.Article) (ProcessedContent, error)
}

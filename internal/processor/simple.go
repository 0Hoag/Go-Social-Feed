package processor

import (
	"context"

	"github.com/bregydoc/gtranslate"
	"github.com/hoag/go-social-feed/internal/crawler"
	"github.com/hoag/go-social-feed/pkg/log"
)

type SimpleProcessor struct {
	l log.Logger
}

func NewSimpleProcessor(l log.Logger) *SimpleProcessor {
	return &SimpleProcessor{l: l}
}

func (p *SimpleProcessor) Process(ctx context.Context, article crawler.Article) (ProcessedContent, error) {
	// 1. Translate Title
	titleVi, err := gtranslate.TranslateWithParams(
		article.Title,
		gtranslate.TranslationParams{
			From: "en",
			To:   "vi",
		},
	)
	if err != nil {
		p.l.Errorf(ctx, "Failed to translate title: %v", err)
		titleVi = article.Title // Fallback
	}

	// 2. "Summarize" (Truncate)
	// Since we don't have a real LLM for free summarization, we just take the first 300 chars
	// In a real crawl, 'article.Content' might be long HTML.
	// Here we assume article.Title or article.Summary is passed.

	rawSummary := article.Title // Often crawler only gets title + link
	if len(article.Summary) > 0 {
		rawSummary = article.Summary
	}

	summaryVi, err := gtranslate.TranslateWithParams(
		rawSummary,
		gtranslate.TranslationParams{
			From: "en",
			To:   "vi",
		},
	)
	if err != nil {
		p.l.Errorf(ctx, "Failed to translate summary: %v", err)
		summaryVi = rawSummary
	}

	return ProcessedContent{
		OriginalTitle:     article.Title,
		TranslatedTitle:   titleVi,
		OriginalSummary:   rawSummary,
		TranslatedSummary: summaryVi,
		SourceURL:         article.SourceURL,
		ImageURL:          article.ImageURL,
	}, nil
}

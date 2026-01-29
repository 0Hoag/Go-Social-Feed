package processor

import (
	"context"
	"strings"

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
		p.l.Errorf(ctx, "processor.simple.Process.TranslateTitle: %v", err)
		return ProcessedContent{}, err
	}

	// Extract first 2-3 sentences from content for summary
	summary := article.Summary
	if summary == "" && article.Content != "" {
		// Take first 300 characters or first 2 sentences
		content := article.Content
		if len(content) > 300 {
			content = content[:300]
		}
		// Find last period to avoid cutting mid-sentence
		lastPeriod := strings.LastIndex(content, ".")
		if lastPeriod > 100 { // Ensure we don't cut too short if no period is found early
			content = content[:lastPeriod+1]
		}
		summary = content
	}

	// If still no summary, use title
	if summary == "" {
		summary = article.Title
	}

	// Translate summary
	summaryVi, err := gtranslate.TranslateWithParams(
		summary,
		gtranslate.TranslationParams{
			From: "en",
			To:   "vi",
		},
	)
	if err != nil {
		p.l.Errorf(ctx, "processor.simple.Process.TranslateSummary: %v", err)
		return ProcessedContent{}, err
	}

	return ProcessedContent{
		OriginalTitle:     article.Title,
		TranslatedTitle:   titleVi,
		OriginalSummary:   summary,
		TranslatedSummary: summaryVi,
		SourceURL:         article.SourceURL,
		ImageURL:          article.ImageURL,
	}, nil
}

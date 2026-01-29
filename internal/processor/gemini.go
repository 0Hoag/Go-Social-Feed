package processor

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/hoag/go-social-feed/internal/crawler"
	"github.com/hoag/go-social-feed/pkg/log"
	"google.golang.org/api/option"
)

type GeminiProcessor struct {
	l      log.Logger
	client *genai.Client
	model  *genai.GenerativeModel
}

func NewGeminiProcessor(ctx context.Context, l log.Logger, apiKey string) (*GeminiProcessor, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	model := client.GenerativeModel("gemini-1.5-flash")
	model.SetTemperature(0.7) // Creative but focused

	return &GeminiProcessor{
		l:      l,
		client: client,
		model:  model,
	}, nil
}

func (p *GeminiProcessor) Process(ctx context.Context, article crawler.Article) (ProcessedContent, error) {
	// Prompt Engineering: merging translation and summarization
	prompt := fmt.Sprintf(`
	You are a professional crypto news editor for a Vietnamese audience.
	Read the following article title and summary.
	Task:
	1. Translate the Title to Vietnamese (catchy, standard journalism style).
	2. Summarize the main points into Vietnamese (2-3 sentences, clear and concise).

	Input:
	Title: %s
	Summary: %s

	Output Format (strict):
	Title: [Vietnamese Title]
	Summary: [Vietnamese Summary]
	`, article.Title, article.Summary)

	resp, err := p.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		p.l.Errorf(ctx, "Gemini GenerateContent failed: %v", err)
		return ProcessedContent{}, err
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return ProcessedContent{}, fmt.Errorf("empty response from Gemini")
	}

	rawText := fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0])

	lines := strings.Split(rawText, "\n")
	var viTitle, viSummary string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Title:") {
			viTitle = strings.TrimSpace(strings.TrimPrefix(line, "Title:"))
		} else if strings.HasPrefix(line, "Summary:") {
			viSummary = strings.TrimSpace(strings.TrimPrefix(line, "Summary:"))
		} else if viSummary != "" && line != "" {
			// multi-line summary
			viSummary += " " + line
		}
	}

	// Fallback if parsing fails slightly (AI sometimes adds bolding etc)
	if viTitle == "" {
		viTitle = article.Title // fallback to original
	}
	if viSummary == "" {
		viSummary = "..." // better than nothing
	}

	return ProcessedContent{
		OriginalTitle:     article.Title,
		TranslatedTitle:   strings.ReplaceAll(viTitle, "**", ""), // Remove markdown bold if any
		OriginalSummary:   article.Summary,
		TranslatedSummary: strings.ReplaceAll(viSummary, "**", ""),
		SourceURL:         article.SourceURL,
		ImageURL:          article.ImageURL,
	}, nil
}

func (p *GeminiProcessor) Close() {
	p.client.Close()
}

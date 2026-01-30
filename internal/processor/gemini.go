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

	model := client.GenerativeModel("gemini-flash-latest")
	model.SetTemperature(0.7) // Creative but focused

	return &GeminiProcessor{
		l:      l,
		client: client,
		model:  model,
	}, nil
}

func (p *GeminiProcessor) Process(ctx context.Context, article crawler.Article) (ProcessedContent, error) {
	// Use full content if available, otherwise fall back to summary
	contentToAnalyze := article.Content
	if contentToAnalyze == "" {
		contentToAnalyze = article.Summary
	}

	// Limit content length to avoid token limits (approximately 10,000 chars = ~2,500 tokens)
	// Gemini Flash has 1M token context, but we'll be conservative
	maxContentLength := 10000
	if len(contentToAnalyze) > maxContentLength {
		contentToAnalyze = contentToAnalyze[:maxContentLength] + "..."
	}

	// Enhanced prompt: Generate 3 distinct outputs
	prompt := fmt.Sprintf(`You are a Vietnamese crypto news editor. Read this English article and create Vietnamese content.

	Article Title: %s
	Full Content:
	%s

	Tasks:
	1. Translate the Title to Vietnamese. IMPORTANT: Keep it SHORT and CONCISE, similar length to original English title. Do NOT summarize the article in the title.
	2. Write a BRIEF Vietnamese summary for news feed (2-3 sentences maximum).
	3. Translate the FULL article content to Vietnamese (natural journalism style).

	Output Format (strict):
	Title: [Short Vietnamese Title]
	Summary: [Brief 2-3 sentence summary]
	FullContent: [Complete Vietnamese translation of the article]
	`, article.Title, contentToAnalyze)

	resp, err := p.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		if !strings.Contains(err.Error(), "429") {
			p.l.Errorf(ctx, "Gemini GenerateContent failed: %v", err)
		}
		return ProcessedContent{}, err
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return ProcessedContent{}, fmt.Errorf("empty response from Gemini")
	}

	rawText := fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0])

	// Parse the response
	lines := strings.Split(rawText, "\n")
	var viTitle, viSummary, viFullContent string
	currentSection := ""

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "Title:") {
			viTitle = strings.TrimSpace(strings.TrimPrefix(line, "Title:"))
			currentSection = "title"
		} else if strings.HasPrefix(line, "Summary:") {
			viSummary = strings.TrimSpace(strings.TrimPrefix(line, "Summary:"))
			currentSection = "summary"
		} else if strings.HasPrefix(line, "FullContent:") {
			viFullContent = strings.TrimSpace(strings.TrimPrefix(line, "FullContent:"))
			currentSection = "fullcontent"
		} else if line != "" {
			// Append to current section
			switch currentSection {
			case "title":
				viTitle += " " + line
			case "summary":
				viSummary += " " + line
			case "fullcontent":
				viFullContent += "\n" + line
			}
		}
	}

	// Fallback if parsing fails
	if viTitle == "" {
		p.l.Warnf(ctx, "Gemini failed to generate title, using original")
		viTitle = article.Title
	}
	if viSummary == "" {
		p.l.Warnf(ctx, "Gemini failed to generate summary, using fallback")
		viSummary = "Nội dung đang được cập nhật..."
	}
	if viFullContent == "" {
		p.l.Warnf(ctx, "Gemini failed to generate full content, using summary")
		viFullContent = viSummary
	}

	return ProcessedContent{
		OriginalTitle:         article.Title,
		TranslatedTitle:       strings.ReplaceAll(viTitle, "**", ""),
		OriginalSummary:       article.Summary,
		TranslatedSummary:     strings.ReplaceAll(viSummary, "**", ""),
		TranslatedFullContent: strings.TrimSpace(strings.ReplaceAll(viFullContent, "**", "")),
		SourceURL:             article.SourceURL,
		ImageURL:              article.ImageURL,
	}, nil
}

func (p *GeminiProcessor) Close() {
	p.client.Close()
}

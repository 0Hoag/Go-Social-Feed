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

	model := client.GenerativeModel("gemini-1.5-flash") // Use specific version for stability
	model.SetTemperature(0.7)
	model.SetMaxOutputTokens(8192) // Ensure enough tokens for full translation

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
	maxContentLength := 15000 // Increased limit
	if len(contentToAnalyze) > maxContentLength {
		contentToAnalyze = contentToAnalyze[:maxContentLength] + "..."
	}

	// Enhanced prompt: Generate 3 distinct outputs
	prompt := fmt.Sprintf(`You are a professional crypto news editor for the Vietnamese market. 
	Your task is to translate the following article into Vietnamese, ensuring NO information is lost.
	
	Article Title: %s
	Full Content:
	%s

	Tasks:
	1. Translate the Title: Keep it concise but accurate.
	2. Write a Summary: 2-3 sentences max.
	3. Translate the Full Content: 
	   - Translate paragraph by paragraph.
	   - Do NOT summarize. Translate the full meaning of every sentence.
	   - Use professional, natural Vietnamese terminology for crypto/finance.
	   - Maintain the original structure and length.
       - Use markdown formatting (bold, italic) where appropriate.

	Output Format (strict):
	Title: [Vietnamese Title]
	Summary: [Vietnamese Summary]
	FullContent:
    [Full Vietnamese Translation - Preserve Paragraphs]
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
		trimmedLine := strings.TrimSpace(line)

		if strings.HasPrefix(trimmedLine, "Title:") {
			viTitle = strings.TrimSpace(strings.TrimPrefix(trimmedLine, "Title:"))
			currentSection = "title"
		} else if strings.HasPrefix(trimmedLine, "Summary:") {
			viSummary = strings.TrimSpace(strings.TrimPrefix(trimmedLine, "Summary:"))
			currentSection = "summary"
		} else if strings.HasPrefix(trimmedLine, "FullContent:") {
			// Check if there is content on the same line
			content := strings.TrimPrefix(trimmedLine, "FullContent:")
			if strings.TrimSpace(content) != "" {
				viFullContent = strings.TrimSpace(content)
			}
			currentSection = "fullcontent"
		} else {
			// Append to current section
			switch currentSection {
			case "title":
				if trimmedLine != "" {
					viTitle += " " + trimmedLine
				}
			case "summary":
				if trimmedLine != "" {
					viSummary += " " + trimmedLine
				}
			case "fullcontent":
				// Preserve empty lines for paragraph separation
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

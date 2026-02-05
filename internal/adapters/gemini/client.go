package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Client struct {
	model *genai.GenerativeModel
}

type AIAnalysisResult struct {
	TrustScore int         `json:"trust_score"`
	Issues     []IssueData `json:"issues"`
}

type IssueData struct {
	Type        string `json:"type"` // "CRITICAL", "WARNING", "INFO"
	Name        string `json:"name"`
	Description string `json:"description"`
	Impact      int    `json:"impact"`
}

func NewClient(apiKey string) *Client {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating Gemini client: %v", err)
	}

	// Use gemini-1.5-flash for speed and efficiency
	model := client.GenerativeModel("gemini-flash-latest")
	model.SetTemperature(0.1) // Low temperature for consistent analysis

	return &Client{model: model}
}

func (c *Client) AnalyzeContract(sourceCode string) (*AIAnalysisResult, error) {
	ctx := context.Background()

	// Truncate if too long (Gemini Flash has ~1M token context, but let's be safe/efficient)
	if len(sourceCode) > 100000 {
		sourceCode = sourceCode[:100000] + "\n... (truncated)"
	}

	prompt := fmt.Sprintf(`
	You are a Senior Smart Contract Auditor. Analyze the following Solidity code for security vulnerabilities.
	Focus on logical flaws, backdoors, hidden mint functions, honeypot mechanisms, and high taxes.

	Source Code:
	%s

	Output STRICTLY JSON in the following format (no markdown code blocks):
	{
	"trust_score": <0-100 integer, 100 is safe, 0 is scam>,
	"issues": [
		{
		"type": "CRITICAL" | "WARNING" | "INFO",
		"name": "<Short Name>",
		"description": "<Concise explanation>",
		"impact": <negative integer, e.g. 40>
		}
	]
	}

	If no issues found, return empty issues array and high trust score.
`, sourceCode)

	resp, err := c.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, fmt.Errorf("gemini generate error: %w", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("empty response from Gemini")
	}

	part := resp.Candidates[0].Content.Parts[0]
	text, ok := part.(genai.Text)
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	// Clean Markdown code blocks if present (Gemini sometimes adds them despite instructions)
	cleanText := strings.TrimSpace(string(text))
	cleanText = strings.TrimPrefix(cleanText, "```json")
	cleanText = strings.TrimPrefix(cleanText, "```")
	cleanText = strings.TrimSuffix(cleanText, "```")

	var result AIAnalysisResult
	if err := json.Unmarshal([]byte(cleanText), &result); err != nil {
		return nil, fmt.Errorf("json parse error: %v | raw: %s", err, cleanText)
	}

	return &result, nil
}

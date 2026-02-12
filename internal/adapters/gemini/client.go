package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Client struct {
	model *genai.GenerativeModel
}

type AIAnalysisResult struct {
	TrustScore   int         `json:"trust_score"`
	Issues       []IssueData `json:"issues"`
	SafeFeatures []string    `json:"safe_features"`
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

	// Use gemini-1.5-flash for stability
	model := client.GenerativeModel("gemini-1.5-flash")
	model.SetTemperature(0.1)

	return &Client{model: model}
}

func (c *Client) AnalyzeContract(sourceCode string) (*AIAnalysisResult, error) {
	// Set a reasonable timeout for AI analysis (e.g. 2 minutes)
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	// Truncate if too long
	if len(sourceCode) > 100000 {
		sourceCode = sourceCode[:100000] + "\n... (truncated)"
	}

	prompt := fmt.Sprintf(`
	You are an Elite Smart Contract Security Auditor. Your job is to analyze Solidity code with extreme depth and nuance.
	
	### OBJECTIVE:
	Determine if the contract is a **Legitimate Project** (DeFi, Stablecoin, Top Token) or a **Scam/Rug Pull**.
	
	### OUTPUT LANGUAGE:
	**CRITICAL: ALL OUTPUT MUST BE IN VIETNAMESE (TIẾNG VIỆT).**
	**CRITICAL: DO NOT USE ENGLISH FOR DESCRIPTIONS OR EXPLANATIONS.**
	Use professional security terminology (e.g., "Cơ chế quản trị", "Hàm Mint", "Lỗ hổng bảo mật", "Thư viện chuẩn").

	### ANALYSIS RULES:
	1. **Context Matters**: 
	   - If code allows **Minting**: Is it restricted to a specific role (e.g. 'minter'), capped, or time-locked? If yes, this is Standard DeFi (Low Risk). If 'onlyOwner' can mint unlimited without checks -> Critical Risk.
	   - If code has **Blacklist**: Is it a standard implementation (like USDC/USDT) for compliance? If yes -> Warning (Centralization). If it blocks allows transfer only to owner -> Critical (Honeypot).
	   - If code is **Proxy/Upgradable**: This is standard for professional projects (e.g. ENA, USDC). Flag as "Upgradable Contract" (Info/Warning), not Critical.

	2. **Code Quality Heuristics**:
	   - Professional formatting, Natspec comments, modular imports (OpenZeppelin) => High Trust Score (+20 pts).
	   - Obfuscated code, lack of comments, single monolithic file => Scam Indicator.

	3. **Scoring Logic**:
	   - **Legit DeFi/Top Coins** (ENA, UNI, USDT): Score should be **65-95** (depending on centralization).
	   - **Scams/Honeypots**: Score should be **0-30**.
	   - **Average Meme Coins**: Score **30-60** (High risk but honest code).

	### OUTPUT FORMAT (JSON ONLY):
	{
		"trust_score": <0-100 integer>,
		"issues": [
			{
				"type": "CRITICAL" | "WARNING" | "INFO",
				"name": "<Tên lỗi (Bằng TIẾNG VIỆT)>",
				"description": "<Mô tả chi tiết bằng TIẾNG VIỆT. Tuyệt đối KHÔNG dùng Tiếng Anh. Giải thích tại sao nguy hiểm.>",
				"impact": <positive integer deduction>
			}
		],
		"safe_features": [
			"<Tính năng an toàn (Bằng TIẾNG VIỆT), vd: 'Mint giới hạn', 'Chuẩn OpenZeppelin', 'Proxy nâng cấp'>"
		]
	}

	Source Code:
	%s
	
	Output STRICTLY JSON. Checks must be rigorous but fair. Response MUST be in Vietnamese.
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

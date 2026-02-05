package scanner

import (
	"log"
	"regexp"

	"github.com/hoag/go-social-feed/internal/adapters/gemini"
)

type IssueType string

const (
	IssueCritical IssueType = "CRITICAL"
	IssueWarning  IssueType = "WARNING"
	IssueInfo     IssueType = "INFO"
)

type Issue struct {
	Type        IssueType
	Name        string
	Description string
	Impact      int // Negative score impact
}

type ScanResult struct {
	TrustScore int
	Issues     []Issue
}

type Engine struct {
	regexRules   map[string]*regexp.Regexp
	geminiClient *gemini.Client
}

func NewEngine(aiClient *gemini.Client) *Engine {
	return &Engine{
		geminiClient: aiClient,
		regexRules: map[string]*regexp.Regexp{
			// Critical: Honeypot indicators
			"blacklist_function":   regexp.MustCompile(`(?i)(function\s+blacklist|mapping\s*\(address\s*=>\s*bool\)\s*.*blacklist)`),
			"transfer_restriction": regexp.MustCompile(`(?i)(require\s*\(.*!isBlacklisted)`),

			// Critical: Rugpull indicators
			"mint_function":      regexp.MustCompile(`(?i)(function\s+mint.*public|function\s+mint.*external)`),
			"infinite_allowance": regexp.MustCompile(`(?i)(allowance\s*=\s*type\(uint256\)\.max)`),

			// Warning: High Tax / Fees
			"high_tax": regexp.MustCompile(`(?i)(fee\s*=\s*[1-9][0-9])`), // Primitive check for fee >= 10

			// Warning: Hidden ownership
			"renounce_ownership": regexp.MustCompile(`(?i)(function\s+renounceOwnership)`),
		},
	}
}

func (e *Engine) Scan(sourceCode string) ScanResult {
	// 1. Try AI Scan First (Deep Analysis)
	if e.geminiClient != nil {
		log.Println("🧠 Running Gemini AI Analysis...")
		aiResult, err := e.geminiClient.AnalyzeContract(sourceCode)
		if err == nil {
			// Map AI result to internal format
			issues := []Issue{}
			for _, i := range aiResult.Issues {
				issues = append(issues, Issue{
					Type:        IssueType(i.Type),
					Name:        i.Name,
					Description: i.Description,
					Impact:      i.Impact,
				})
			}
			return ScanResult{
				TrustScore: aiResult.TrustScore,
				Issues:     issues,
			}
		} else {
			log.Printf("⚠️ Gemini Scan Failed: %v. Falling back to Regex.", err)
		}
	}

	// 2. Fallback: Basic Regex Scan
	log.Println("⚡ Running Basic Regex Scan...")
	issues := []Issue{}
	score := 100

	// 1. Basic Regex Checks
	if e.regexRules["blacklist_function"].MatchString(sourceCode) {
		issues = append(issues, Issue{
			Type:        IssueWarning,
			Name:        "Centralization Risk: Blacklist",
			Description: "Admin role can restrict token transfers (Common in Stablecoins/CEX tokens).",
			Impact:      25,
		})
		score -= 25
	}

	if e.regexRules["mint_function"].MatchString(sourceCode) {
		issues = append(issues, Issue{
			Type:        IssueCritical,
			Name:        "Critical: Uncapped Supply",
			Description: "Owner can mint infinite tokens (High Rug Pull Risk if not Time-locked).",
			Impact:      40,
		})
		score -= 40
	}

	if e.regexRules["high_tax"].MatchString(sourceCode) {
		issues = append(issues, Issue{
			Type:        IssueWarning,
			Name:        "High Tax Variables",
			Description: "Detected potential high fee settings (>10%).",
			Impact:      15,
		})
		score -= 15
	}

	// Cap score at 0
	if score < 0 {
		score = 0
	}

	return ScanResult{
		TrustScore: score,
		Issues:     issues,
	}
}

package scanner

import (
	"regexp"
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
	regexRules map[string]*regexp.Regexp
}

func NewEngine() *Engine {
	return &Engine{
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
	issues := []Issue{}
	score := 100

	// 1. Basic Regex Checks
	if e.regexRules["blacklist_function"].MatchString(sourceCode) {
		issues = append(issues, Issue{
			Type:        IssueWarning,
			Name:        "Blacklist Function Detected",
			Description: "Admin might be able to block your wallet.",
			Impact:      25,
		})
		score -= 25
	}

	if e.regexRules["mint_function"].MatchString(sourceCode) {
		issues = append(issues, Issue{
			Type:        IssueCritical,
			Name:        "Uncapped Mint Function",
			Description: "Owner might be able to mint infinite tokens (Rug Pull risk).",
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

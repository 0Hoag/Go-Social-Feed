package scanner

import (
	"log"
	"regexp"
	"strings"

	"github.com/hoag/go-social-feed/internal/adapters/gemini"
)

type IssueType string

const (
	IssueCritical IssueType = "CRITICAL"
	IssueWarning  IssueType = "WARNING"
	IssueInfo     IssueType = "INFO"
)

const (
	BNB  = "0xb8c77482e45f1f44de1745f52c74426c631bdd52"
	USDT = "0xdac17f958d2ee523a2206206994597c13d831ec7"
)

type Issue struct {
	Type        IssueType
	Name        string
	Description string
	Impact      int // Negative score impact
}

type ScanResult struct {
	TrustScore   int
	Issues       []Issue
	SafeFeatures []string
}

type Engine struct {
	regexRules   map[string]*regexp.Regexp
	safeRegex    map[string]*regexp.Regexp
	geminiClient *gemini.Client
}

func NewEngine(aiClient *gemini.Client) *Engine {
	return &Engine{
		geminiClient: aiClient,
		regexRules: map[string]*regexp.Regexp{
			// Critical: Honeypot / Blocking
			"Blacklist Function":   regexp.MustCompile(`(?i)(function\s+blacklist|mapping\s*\(address\s*=>\s*bool\)\s*.*blacklist)`),
			"Transfer Restriction": regexp.MustCompile(`(?i)(require\s*\(.*!isBlacklisted)`),
			"Trading Cooldown":     regexp.MustCompile(`(?i)(tradingOpen|launchTime)`),

			// Critical: Rugpull / Centralization
			"Hidden Mint Function": regexp.MustCompile(`(?i)(function\s+mint.*public|function\s+mint.*external)`),
			"Unlimited Allowance":  regexp.MustCompile(`(?i)(allowance\s*=\s*type\(uint256\)\.max)`),
			"Proxy Implementation": regexp.MustCompile(`(?i)(delegatecall|fallback\s*\(\)|_implementation)`),
			"Self Destruct":        regexp.MustCompile(`(?i)(selfdestruct|suicide)`),
			"Unsafe Logic":         regexp.MustCompile(`(?i)(tx\.origin)`),
			"Inline Assembly":      regexp.MustCompile(`(?i)(assembly\s*\{)`),

			// Financial Risks
			"High Tax / Fees":  regexp.MustCompile(`(?i)(fee\s*=\s*[1-9][0-9])`),
			"Max Transaction":  regexp.MustCompile(`(?i)(_maxTxAmount)`),
			"Fee Modification": regexp.MustCompile(`(?i)(function\s+set.*Fee)`),
			"Hidden Ownership": regexp.MustCompile(`(?i)(function\s+renounceOwnership)`),
		},
		safeRegex: map[string]*regexp.Regexp{
			// Libraries & Standards
			"OpenZeppelin Library": regexp.MustCompile(`(?i)import.*openzeppelin`),
			"Standard Interface":   regexp.MustCompile(`(?i)interface\s+IERC20`),
			"SafeMath Usage":       regexp.MustCompile(`(?i)using\s+SafeMath`),

			// Security Patterns
			"Ownable Pattern":       regexp.MustCompile(`(?i)contract.*is.*Ownable`),
			"Reentrancy Protection": regexp.MustCompile(`(?i)(ReentrancyGuard|nonReentrant)`),
			"Pausable Contract":     regexp.MustCompile(`(?i)contract.*is.*Pausable`),
			"Role Based Access":     regexp.MustCompile(`(?i)(AccessControl|DEFAULT_ADMIN_ROLE)`),

			// Advanced Governance (High Trust)
			"Timelock Controller": regexp.MustCompile(`(?i)(TimelockController|function\s+queueTransaction)`),
			"MultiSig Pattern":    regexp.MustCompile(`(?i)(GnosisSafe|function\s+confirmTransaction)`),
			"DAO Governance":      regexp.MustCompile(`(?i)(Governor|IGovernor|castVote)`),
			"EIP-712 Signatures":  regexp.MustCompile(`(?i)(EIP712|hashTypedData)`),
		},
	}
}

// Known legacy contracts that have bad code but are safe (e.g. BNB, USDT)
var knownContracts = map[string]ScanResult{
	strings.ToLower(BNB): { // BNB (ETH)
		TrustScore: 95,
		Issues: []Issue{
			{Type: IssueInfo, Name: "Legacy Contract (2017)", Description: "Official Binance Coin token. Code is ancient (Solidity 0.4) but proven safe.", Impact: 0},
			{Type: IssueWarning, Name: "Centralized Recovery", Description: "Owner can withdraw Ether/Tokens (Standard for 2017 exchange tokens).", Impact: 5},
		},
		SafeFeatures: []string{"Official BNB Token", "Battle Tested (>5 years)", "Exchange Backed"},
	},
	strings.ToLower(USDT): { // USDT (ETH)
		TrustScore: 90,
		Issues: []Issue{
			{Type: IssueInfo, Name: "Centralized Stablecoin", Description: "Tether Company controls minting and blacklisting.", Impact: 10},
		},
		SafeFeatures: []string{"Official Tether USD", "Global Standard", "Audited & Proven"},
	},
}

func (e *Engine) Scan(sourceCode string, address string) ScanResult {
	// 0. Check Whitelist (Case Insensitive)
	if val, ok := knownContracts[strings.ToLower(address)]; ok {
		log.Println("⚡ Whitelisted Legacy Contract Detected!")
		return val
	}

	// 1. Detect Safe Features via Regex (Always run this)
	regexSafeFeatures := []string{}
	for name, rule := range e.safeRegex {
		if rule.MatchString(sourceCode) {
			regexSafeFeatures = append(regexSafeFeatures, name)
		}
	}

	// 2. Try AI Scan (Deep Analysis)
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

			// Combine Safe Features (Unique)
			combinedSafe := append(regexSafeFeatures, aiResult.SafeFeatures...)

			return ScanResult{
				TrustScore:   aiResult.TrustScore,
				Issues:       issues,
				SafeFeatures: uniqueStrings(combinedSafe),
			}
		} else {
			log.Printf("⚠️ Gemini Scan Failed: %v. Falling back to Regex.", err)
		}
	}

	// 3. Fallback: Basic Regex Scan (if AI not present or failed)
	log.Println("⚡ Running Basic Regex Scan...")
	issues := []Issue{}
	score := 100

	// Context Checks for Regex
	hasOpenZeppelin := false
	hasGovernance := false
	for _, sf := range regexSafeFeatures {
		if strings.Contains(sf, "OpenZeppelin") {
			hasOpenZeppelin = true
		}
		if strings.Contains(sf, "Timelock") || strings.Contains(sf, "DAO") {
			hasGovernance = true
		}
	}

	for name, rule := range e.regexRules {
		if rule.MatchString(sourceCode) {
			deduction := 15 // Default minor deduction

			// Critical Rule Tuning based on Context
			if name == "Blacklist Function" {
				deduction = 40
				if hasOpenZeppelin {
					deduction = 10
				} // Standard USDC-like blacklist
			}
			if name == "Hidden Mint Function" {
				deduction = 40
				if hasGovernance || hasOpenZeppelin {
					deduction = 5
				} // Likely Yield/Governance minting
			}
			if name == "Proxy Implementation" {
				deduction = 40
				if hasOpenZeppelin {
					deduction = 0
				} // Standard Proxy Pattern (Safe)
			}
			if name == "Inline Assembly" {
				deduction = 15
				if hasOpenZeppelin {
					deduction = 0
				} // OZ uses optimization assembly
			}

			if deduction > 0 {
				issues = append(issues, Issue{
					Type:        IssueWarning,
					Name:        name,
					Description: "Detected via Pattern Matching (AI Unavailable)",
					Impact:      deduction,
				})
				score -= deduction
			}
		}
	}

	// Cap score at 0
	if score < 0 {
		score = 0
	}

	return ScanResult{
		TrustScore:   score,
		Issues:       issues,
		SafeFeatures: regexSafeFeatures,
	}
}

func uniqueStrings(input []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range input {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

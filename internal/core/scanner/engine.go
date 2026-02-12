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
			"Hàm Blacklist (Cấm ví)":  regexp.MustCompile(`(?i)(function\s+blacklist|mapping\s*\(address\s*=>\s*bool\)\s*.*blacklist)`),
			"Hạn chế Giao dịch":       regexp.MustCompile(`(?i)(require\s*\(.*!isBlacklisted)`),
			"Thời gian chờ giao dịch": regexp.MustCompile(`(?i)(tradingOpen|launchTime)`),

			// Critical: Rugpull / Centralization
			"Hàm Mint Ẩn (In tiền)":           regexp.MustCompile(`(?i)(function\s+mint.*public|function\s+mint.*external)`),
			"Cấp quyền Vô hạn":                regexp.MustCompile(`(?i)(allowance\s*=\s*type\(uint256\)\.max)`),
			"Proxy (Có thể nâng cấp)":         regexp.MustCompile(`(?i)(delegatecall|fallback\s*\(\)|_implementation)`),
			"Hàm Tự hủy (Self Destruct)":      regexp.MustCompile(`(?i)(selfdestruct|suicide)`),
			"Logic Không an toàn (tx.origin)": regexp.MustCompile(`(?i)(tx\.origin)`),
			"Assembly Nội bộ (Khó kiểm tra)":  regexp.MustCompile(`(?i)(assembly\s*\{)`),

			// Financial Risks
			"Thuế / Phí Cao":             regexp.MustCompile(`(?i)(fee\s*=\s*[1-9][0-9])`),
			"Giới hạn Giao dịch (MaxTx)": regexp.MustCompile(`(?i)(_maxTxAmount)`),
			"Hàm Chỉnh sửa Phí":          regexp.MustCompile(`(?i)(function\s+set.*Fee)`),
			"Hàm Từ bỏ Quyền sở hữu":     regexp.MustCompile(`(?i)(function\s+renounceOwnership)`),
		},
		safeRegex: map[string]*regexp.Regexp{
			// Libraries & Standards
			"Thư viện OpenZeppelin":    regexp.MustCompile(`(?i)import.*openzeppelin`),
			"Tuân thủ chuẩn Interface": regexp.MustCompile(`(?i)interface\s+IERC20`),
			"Sử dụng SafeMath":         regexp.MustCompile(`(?i)using\s+SafeMath`),

			// Security Patterns
			"Mô hình Ownable":            regexp.MustCompile(`(?i)contract.*is.*Ownable`),
			"Bảo vệ Reentrancy":          regexp.MustCompile(`(?i)(ReentrancyGuard|nonReentrant)`),
			"Có thể Tạm dừng (Pausable)": regexp.MustCompile(`(?i)contract.*is.*Pausable`),
			"Phân quyền (Role Based)":    regexp.MustCompile(`(?i)(AccessControl|DEFAULT_ADMIN_ROLE)`),

			// Advanced Governance (High Trust)
			"Timelock (Khóa thời gian)": regexp.MustCompile(`(?i)(TimelockController|function\s+queueTransaction)`),
			"Ví Đa chữ ký (MultiSig)":   regexp.MustCompile(`(?i)(GnosisSafe|function\s+confirmTransaction)`),
			"Quản trị DAO":              regexp.MustCompile(`(?i)(Governor|IGovernor|castVote)`),
			"Chữ ký EIP-712":            regexp.MustCompile(`(?i)(EIP712|hashTypedData)`),
		},
	}
}

// Known legacy contracts that have bad code but are safe (e.g. BNB, USDT)
var knownContracts = map[string]ScanResult{
	strings.ToLower(BNB): { // BNB (ETH)
		TrustScore: 95,
		Issues: []Issue{
			{Type: IssueInfo, Name: "Contract Cũ (2017)", Description: "Token BNB gốc. Code cũ (Solidity 0.4) nhưng đã được kiểm chứng an toàn.", Impact: 0},
			{Type: IssueWarning, Name: "Phục hồi Tập trung", Description: "Chủ sở hữu có thể rút Ether/Token (Tiêu chuẩn của token sàn 2017).", Impact: 5},
		},
		SafeFeatures: []string{"Token BNB Chính chủ", "Đã kiểm chứng (>5 năm)", "Sàn Binance bảo chứng"},
	},
	strings.ToLower(USDT): { // USDT (ETH)
		TrustScore: 90,
		Issues: []Issue{
			{Type: IssueInfo, Name: "Stablecoin Tập trung", Description: "Công ty Tether kiểm soát việc in tiền và khóa ví.", Impact: 10},
		},
		SafeFeatures: []string{"Tether USD Chính chủ", "Tiêu chuẩn Toàn cầu", "Đã được kiểm toán & chứng minh"},
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
			if name == "Hàm Blacklist (Cấm ví)" {
				deduction = 40
				if hasOpenZeppelin {
					deduction = 10
				} // Standard USDC-like blacklist
			}
			if name == "Hàm Mint Ẩn (In tiền)" {
				deduction = 40
				if hasGovernance || hasOpenZeppelin {
					deduction = 5
				} // Likely Yield/Governance minting
			}
			if name == "Proxy (Có thể nâng cấp)" {
				deduction = 40
				if hasOpenZeppelin {
					deduction = 0
				} // Standard Proxy Pattern (Safe)
			}
			if name == "Assembly Nội bộ (Khó kiểm tra)" {
				deduction = 15
				if hasOpenZeppelin {
					deduction = 0
				} // OZ uses optimization assembly
			}

			if deduction > 0 {
				issues = append(issues, Issue{
					Type:        IssueWarning,
					Name:        name,
					Description: "Phát hiện qua so khớp mẫu (AI không khả dụng)",
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

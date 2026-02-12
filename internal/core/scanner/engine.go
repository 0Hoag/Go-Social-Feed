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

func (e *Engine) Scan(sourceCode string, address string, language string) ScanResult {
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
		aiResult, err := e.geminiClient.AnalyzeContract(sourceCode, language)
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
			var description string

			// Build detailed description based on rule name and language
			isVi := (language == "vi")

			// Critical Rule Tuning based on Context
			if name == "Hàm Blacklist (Cấm ví)" || name == "Blacklist Function" {
				deduction = 40
				if hasOpenZeppelin {
					deduction = 10
				} // Standard USDC-like blacklist
				if isVi {
					description = "Hợp đồng có khả năng chặn địa chỉ ví cụ thể khỏi giao dịch. Nếu đây là stablecoin (USDT/USDC) thì đây là tính năng tuân thủ pháp luật bình thường. Nhưng nếu là token meme/DeFi, chủ sở hữu có thể lạm dụng để chặn người dùng bán token (Honeypot)."
				} else {
					description = "Contract can block specific wallet addresses from trading. If this is a stablecoin (USDT/USDC), this is a normal compliance feature. But if it's a meme/DeFi token, the owner could abuse this to prevent users from selling (Honeypot)."
				}
			} else if name == "Hàm Mint Ẩn (In tiền)" || name == "Hidden Mint Function" {
				deduction = 40
				if hasGovernance || hasOpenZeppelin {
					deduction = 5
				} // Likely Yield/Governance minting
				if isVi {
					description = "Hợp đồng có hàm tạo thêm token (mint). Nếu có cơ chế quản trị (DAO/Timelock) hoặc giới hạn rõ ràng thì an toàn. Nhưng nếu chủ sở hữu có thể mint vô hạn mà không kiểm soát, họ có thể pha loãng giá trị token của bạn bất cứ lúc nào (Rug Pull)."
				} else {
					description = "Contract has a function to create new tokens (mint). If there's governance (DAO/Timelock) or clear limits, it's safe. But if the owner can mint unlimited tokens without control, they can dilute your token value anytime (Rug Pull)."
				}
			} else if name == "Proxy (Có thể nâng cấp)" || name == "Proxy Implementation" {
				deduction = 40
				if hasOpenZeppelin {
					deduction = 0
				} // Standard Proxy Pattern (Safe)
				if isVi {
					description = "Hợp đồng có thể được nâng cấp (Proxy Pattern). Đây là chuẩn cho các dự án DeFi chuyên nghiệp (ENA, USDC). Tuy nhiên, nếu không có Timelock hoặc MultiSig, chủ sở hữu có thể thay đổi logic hợp đồng bất cứ lúc nào mà không cần thông báo."
				} else {
					description = "Contract is upgradeable (Proxy Pattern). This is standard for professional DeFi projects (ENA, USDC). However, without Timelock or MultiSig, the owner can change contract logic anytime without notice."
				}
			} else if name == "Assembly Nội bộ (Khó kiểm tra)" || name == "Inline Assembly" {
				deduction = 15
				if hasOpenZeppelin {
					deduction = 0
				} // OZ uses optimization assembly
				if isVi {
					description = "Hợp đồng sử dụng Assembly (mã máy cấp thấp). Điều này có thể là tối ưu hóa hợp pháp (OpenZeppelin dùng), nhưng cũng có thể che giấu logic độc hại khó phát hiện. Cần kiểm toán kỹ lưỡng."
				} else {
					description = "Contract uses Assembly (low-level machine code). This could be legitimate optimization (OpenZeppelin uses it), but it can also hide malicious logic that's hard to detect. Requires thorough audit."
				}
			} else if name == "Thuế / Phí Cao" || name == "High Tax" {
				deduction = 20
				if isVi {
					description = "Hợp đồng có phí giao dịch cao (>=10%). Điều này có nghĩa là mỗi lần bạn mua/bán, một phần lớn giá trị sẽ bị trừ đi. Phí quá cao có thể là dấu hiệu của dự án lừa đảo hoặc thiết kế kém."
				} else {
					description = "Contract has high transaction fees (>=10%). This means every time you buy/sell, a large portion of value is deducted. Excessively high fees may indicate a scam project or poor design."
				}
			} else if name == "Hàm Chỉnh sửa Phí" || name == "Fee Modification" {
				deduction = 15
				if isVi {
					description = "Chủ sở hữu có thể thay đổi phí giao dịch bất cứ lúc nào. Họ có thể tăng phí lên 99% sau khi bạn mua, khiến bạn không thể bán được (Honeypot). Rủi ro cao nếu không có giới hạn rõ ràng."
				} else {
					description = "Owner can change transaction fees anytime. They could raise fees to 99% after you buy, making it impossible to sell (Honeypot). High risk if there are no clear limits."
				}
			} else if name == "Ẩn chủ hữu (Hidden Ownership)" || name == "Hidden Ownership" {
				deduction = 25
				if isVi {
					description = "Không tìm thấy hàm từ bỏ quyền sở hữu (renounceOwnership). Điều này có nghĩa là chủ sở hữu vẫn giữ toàn quyền kiểm soát hợp đồng và có thể thay đổi bất cứ điều gì. Rủi ro tập trung hóa cao."
				} else {
					description = "No renounceOwnership function found. This means the owner retains full control over the contract and can change anything. High centralization risk."
				}
			} else {
				// Generic fallback for other rules
				if isVi {
					description = "Phát hiện qua so khớp mẫu. Cần kiểm tra thủ công để xác định mức độ rủi ro chính xác."
				} else {
					description = "Detected via pattern matching. Manual review needed to determine exact risk level."
				}
			}

			if deduction > 0 {
				issues = append(issues, Issue{
					Type:        IssueWarning,
					Name:        name,
					Description: description,
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

package telegram

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hoag/go-social-feed/internal/adapters/dexscreener"
	"github.com/hoag/go-social-feed/internal/adapters/etherscan"
	"github.com/hoag/go-social-feed/internal/core/scanner"
)

type ScannerBot struct {
	bot        *tgbotapi.BotAPI
	ethClient  *etherscan.Client
	dexClient  *dexscreener.Client
	scanEngine *scanner.Engine

	// Rate limiting
	lastScan map[int64]time.Time
	mu       sync.Mutex
}

func NewScannerBot(token string, ethClient *etherscan.Client, engine *scanner.Engine) (*ScannerBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot: %w", err)
	}

	bot.Debug = false // Disable debug valid to reduce noise
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &ScannerBot{
		bot:        bot,
		ethClient:  ethClient,
		dexClient:  dexscreener.NewClient(),
		scanEngine: engine,
		lastScan:   make(map[int64]time.Time),
	}, nil
}

func (s *ScannerBot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := s.bot.GetUpdatesChan(u)

	log.Println("✅ Bot is online and listening for addresses...")

	for update := range updates {
		if update.Message == nil {
			continue
		}

		go s.handleMessage(update.Message)
	}
}

func (s *ScannerBot) handleMessage(msg *tgbotapi.Message) {
	chatID := msg.Chat.ID
	userID := msg.From.ID
	text := strings.TrimSpace(msg.Text)

	// Rate Limiting Check
	s.mu.Lock()
	lastTime, exists := s.lastScan[userID]
	if exists && time.Since(lastTime) < 10*time.Second {
		remaining := 10*time.Second - time.Since(lastTime)
		s.mu.Unlock()
		s.sendReply(chatID, fmt.Sprintf("⏳ Please wait %d seconds before scanning again.", int(remaining.Seconds())))
		return
	}
	s.lastScan[userID] = time.Now()
	s.mu.Unlock()

	// 1. Send "Typing..." action
	action := tgbotapi.NewChatAction(chatID, tgbotapi.ChatTyping)
	s.bot.Send(action)

	// 2. Initial Reply
	statusMsg, _ := s.bot.Send(tgbotapi.NewMessage(chatID, "🔍 **ChainGuardian AI** is analyzing..."))

	// 3. Determine if Input is Address or Symbol
	addrRegex := regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`)
	isAddress := addrRegex.MatchString(text) || text == "0xMOCK"

	var contractAddress string

	if isAddress {
		contractAddress = text
	} else {
		// Assume Symbol Search
		if len(text) > 10 || strings.Contains(text, " ") {
			s.editMessage(chatID, statusMsg.MessageID, "❌ Invalid input. Please send a Contract Address (0x...) or a Token Symbol (e.g. PEPE).")
			return
		}

		s.editMessage(chatID, statusMsg.MessageID, fmt.Sprintf("🔎 Searching for token **%s**...", strings.ToUpper(text)))

		foundAddr, foundNet, foundName, err := s.dexClient.SearchTopToken(text)
		if err != nil {
			s.editMessage(chatID, statusMsg.MessageID, fmt.Sprintf("❌ Not found: %s\n(%v)", text, err))
			return
		}

		s.editMessage(chatID, statusMsg.MessageID, fmt.Sprintf("✅ Found **%s** on **%s**\n📍 Address: `%s`\n🚀 Starting scan...", foundName, strings.ToUpper(foundNet), foundAddr))
		contractAddress = foundAddr

		// Small delay to let user read
		time.Sleep(1 * time.Second)
	}

	// 4. Start Scan Process
	s.processScan(chatID, statusMsg.MessageID, contractAddress)
}

func (s *ScannerBot) processScan(chatID int64, messageID int, address string) {
	// Fetch Source Code (Multi-chain Auto-Discovery)
	var source string
	var name string // Etherscan name (might be slightly different from DexScreener but authoritative for code)
	var err error
	var networkFound string

	// Try all networks
	networks := []string{
		etherscan.NetworkETH,
		etherscan.NetworkBSC,
		etherscan.NetworkBase,
		etherscan.NetworkArbitrum,
		etherscan.NetworkPolygon,
	}

	for _, net := range networks {
		// Update status
		s.editMessage(chatID, messageID, fmt.Sprintf("🔍 Scanning on **%s** network...", strings.ToUpper(net)))

		source, name, err = s.ethClient.GetContractSource(net, address)
		if err == nil && source != "" {
			networkFound = net
			break
		}
	}

	if networkFound == "" {
		s.editMessage(chatID, messageID, fmt.Sprintf("❌ **Scan Failed**\nContract not found on ETH, BSC, or BASE.\n(Note: Verify your API keys and contract address)"))
		return
	}

	// Run Scan Engine
	// Update status before heavy AI task
	s.editMessage(chatID, messageID, fmt.Sprintf("🧠 **Elite Auditor AI** is analyzing logic on **%s**...\n(This might take up to 60s for complex contracts)", strings.ToUpper(networkFound)))

	result := s.scanEngine.Scan(source, address)

	// Format Output
	report := s.formatReport(address, networkFound, name, result)

	// Send Final Report
	s.editMessage(chatID, messageID, report)
}

func (s *ScannerBot) formatReport(address, network, name string, result scanner.ScanResult) string {
	// Create "0xHOANG..." short address
	shortAddr := address
	if len(address) > 10 {
		shortAddr = address[:6] + "..." + address[len(address)-4:]
	}

	var issuesList string
	for _, issue := range result.Issues {
		icon := "⚠️"
		if issue.Type == scanner.IssueCritical {
			icon = "❌" // Critical
		} else if issue.Type == scanner.IssueWarning {
			icon = "🔸" // Warning
		}

		// New detailed format
		issuesList += fmt.Sprintf("• %s **%s** (-%d pts)\n   _Violates:_ %s\n", icon, issue.Name, issue.Impact, issue.Description)
	}

	if len(result.Issues) == 0 {
		issuesList = "• ✅ No critical vulnerabilities found (+0 pts)\n"
	}

	// Determine Safety Advice
	advice := "✅ Contract looks legitimate. Safe to trade."
	scoreColor := "✅"
	riskLevel := "SAFE"

	if result.TrustScore < 50 {
		advice = "⚠️ **Advice:** Exercise extreme caution! High potential for rug pull or scam. Do not invest unless you understand the risks."
		scoreColor = "🛑"
		riskLevel = "HIGH RISK"
	} else if result.TrustScore < 80 {
		advice = "⚠️ **Advice:** Suspicious centralized features detected. Audit heavily before investing."
		scoreColor = "⚠️"
		riskLevel = "MEDIUM RISK"
	}

	safeList := "• Source code verified on Explorer\n"
	for _, feature := range result.SafeFeatures {
		safeList += fmt.Sprintf("• ✅ %s\n", feature)
	}

	return fmt.Sprintf(`🛡 *ChainGuardian AI Report*
	---------------------------
	📍 *Network:* %s
	📍 *Name:* %s
	📍 *Contract:* `+"`%s`"+`
	📊 *Trust Score:* %s **%d/100** (%s)

	❌ *RISK ANALYSIS:*
	%s
	✅ *SAFETY FEATURES:*
	%s

	%s`, strings.ToUpper(network), name, shortAddr, scoreColor, result.TrustScore, riskLevel, issuesList, safeList, advice)
}

func (s *ScannerBot) sendReply(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "Markdown"
	s.bot.Send(msg)
}

func (s *ScannerBot) editMessage(chatID int64, msgID int, text string) {
	edit := tgbotapi.NewEditMessageText(chatID, msgID, text)
	edit.ParseMode = "Markdown"
	s.bot.Send(edit)
}

package telegram

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hoag/go-social-feed/internal/adapters/etherscan"
	"github.com/hoag/go-social-feed/internal/core/scanner"
)

type ScannerBot struct {
	bot        *tgbotapi.BotAPI
	ethClient  *etherscan.Client
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

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &ScannerBot{
		bot:        bot,
		ethClient:  ethClient,
		scanEngine: engine,
		lastScan:   make(map[int64]time.Time),
	}, nil
}

func (s *ScannerBot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := s.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		go s.handleMessage(update.Message)
	}
}

func (s *ScannerBot) handleMessage(msg *tgbotapi.Message) {
	// 1. Regular text processing
	text := strings.TrimSpace(msg.Text)

	// Regex for Ethereum Address
	ethRegex := regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`)

	if text == "0xMOCK" || ethRegex.MatchString(text) {
		s.processScan(msg, text)
	} else if text == "/start" || text == "/help" {
		s.sendReply(msg.Chat.ID, "👋 Welcome directly to **ChainGuardian AI**!\n\nSend me a Smart Contract address (starting with `0x...`) to scan it for vulnerabilities.\n\n⚠️ Rate Limit: 1 scan/minute.")
	} else {
		s.sendReply(msg.Chat.ID, "I only understand Ethereum addresses (0x...) or /start.")
	}
}

func (s *ScannerBot) processScan(msg *tgbotapi.Message, address string) {
	chatID := msg.Chat.ID
	userID := msg.From.ID

	// 2. Rate Limiting Check
	s.mu.Lock()
	lastTime, exists := s.lastScan[userID]
	if exists && time.Since(lastTime) < time.Minute {
		remaining := time.Minute - time.Since(lastTime)
		s.mu.Unlock()
		s.sendReply(chatID, fmt.Sprintf("⏳ **Rate Limit Exceeded**\nPlease wait %d seconds before scanning again.", int(remaining.Seconds())))
		return
	}
	s.lastScan[userID] = time.Now()
	s.mu.Unlock()

	// 3. User Feedback: "Typing..." action
	action := tgbotapi.NewChatAction(chatID, tgbotapi.ChatTyping)
	s.bot.Send(action)

	// Send initial status
	statusMsg, _ := s.bot.Send(tgbotapi.NewMessage(chatID, "🔍 **ChainGuardian AI** is analyzing..."))

	// 4. Fetch Source Code (Multi-chain Auto-Discovery)
	var source string
	var err error
	var networkFound string

	// Try all networks
	networks := []string{etherscan.NetworkETH, etherscan.NetworkBSC, etherscan.NetworkBase}

	for _, net := range networks {
		// Update status
		s.editMessage(chatID, statusMsg.MessageID, fmt.Sprintf("🔍 Scanning on **%s** network...", strings.ToUpper(net)))

		source, err = s.ethClient.GetContractSource(net, address)
		if err == nil && source != "" {
			networkFound = net
			break
		}
	}

	if networkFound == "" {
		s.editMessage(chatID, statusMsg.MessageID, fmt.Sprintf("❌ **Scan Failed**\nContract not found on ETH, BSC, or BASE.\n(Note: Verify your API keys and contract address)"))
		return
	}

	// 5. Run Scan Engine
	result := s.scanEngine.Scan(source)

	// 6. Format Output (User's Template)
	report := s.formatReport(address, networkFound, result)

	// 7. Send Final Report
	s.editMessage(chatID, statusMsg.MessageID, report)
}

func (s *ScannerBot) formatReport(address, network string, result scanner.ScanResult) string {
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

		issuesList += fmt.Sprintf("• %s %s (-%d pts)\n", icon, issue.Name, issue.Impact)
	}

	if len(result.Issues) == 0 {
		issuesList = "• No critical vulnerabilities found (+0 pts)\n"
	}

	// Determine Safety Advice
	advice := "✅ Contract looks safe, but always DYOR."
	scoreColor := "✅"
	riskLevel := "AN TOÀN"

	if result.TrustScore < 50 {
		advice = "⚠️ **Advice:** Hãy cực kỳ cẩn trọng trước khi nạp tiền!"
		scoreColor = "🛑"
		riskLevel = "RỦI RO CAO"
	} else if result.TrustScore < 80 {
		advice = "⚠️ **Advice:** Có một số dấu hiệu đáng ngờ. Hãy check kỹ."
		scoreColor = "⚠️"
		riskLevel = "TRUNG BÌNH"
	}

	safeList := ""
	if result.TrustScore > 0 {
		safeList = "• Source code verified (+10 pts)\n"
		// Real engine assumes verified code if we got this far
	}

	return fmt.Sprintf(`🛡 *ChainGuardian AI Report*
	---------------------------
	📍 *Network:* %s
	📍 *Contract:* `+"`%s`"+`
	📊 *Trust Score:* %s **%d/100** (%s)

	❌ *Issues Found:*
	%s
	✅ *Safe:*
	%s

	%s`, strings.ToUpper(network), shortAddr, scoreColor, result.TrustScore, riskLevel, issuesList, safeList, advice)
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

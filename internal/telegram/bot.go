package telegram

import (
	"context"
	"fmt"
	"html"
	"regexp"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hoag/go-social-feed/pkg/log"
)

type TelegramClient struct {
	bot    *tgbotapi.BotAPI
	chatID int64
	l      log.Logger
}

func NewTelegramClient(token string, chatID int64, l log.Logger) (*TelegramClient, error) {
	if token == "" || chatID == 0 {
		return nil, nil // Return nil if not configured, safe to ignore
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("failed to init telegram bot: %w", err)
	}

	return &TelegramClient{
		bot:    bot,
		chatID: chatID,
		l:      l,
	}, nil
}

func (t *TelegramClient) SendPost(ctx context.Context, title, summary, imageURL, sourceURL string) error {
	if t == nil || t.bot == nil {
		return nil // No-op if disabled
	}

	// Clean up summary (remove HTML tags)
	cleanSummary := stripHTMLTags(summary)

	// Format processing using HTML ParseMode to handle special chars safer
	// Escape all dynamic content to avoid breaking HTML structure
	caption := fmt.Sprintf(""+
		"🚀 <b>NEW ARTICLE FOUND</b>\n"+
		"📝 <b>Title</b>: %s\n"+
		"🔗 <b>Source</b>: %s\n"+
		"📜 <b>Content</b>: %s\n",
		html.EscapeString(title),
		html.EscapeString(sourceURL),
		html.EscapeString(cleanSummary))

	if len(caption) > 1024 {
		caption = caption[:1021] + "..."
	}

	var msg tgbotapi.Chattable

	if imageURL != "" {
		// Verify image URL access first to avoid Telegram API error
		// Or just try. If failed, fallback to text.
		photo := tgbotapi.NewPhoto(t.chatID, tgbotapi.FileURL(imageURL))
		photo.Caption = caption
		photo.ParseMode = "HTML"
		msg = photo
	} else {
		txt := tgbotapi.NewMessage(t.chatID, caption)
		txt.ParseMode = "HTML"
		// Disable web page preview if you want, but likely we want it if no image
		msg = txt
	}

	_, err := t.bot.Send(msg)
	if err != nil {
		// If photo failed (e.g. invalid URL or bad format), try sending just text
		if imageURL != "" {
			t.l.Warnf(ctx, "Telegram: Failed to send photo, falling back to text: %v", err)
			txt := tgbotapi.NewMessage(t.chatID, caption)
			txt.ParseMode = "HTML"
			_, err = t.bot.Send(txt)
			return err
		}
		return err
	}

	return nil
}

// stripHTMLTags removes HTML tags from a string using regex
func stripHTMLTags(s string) string {
	re := regexp.MustCompile(`<[^>]*>`)
	return strings.TrimSpace(re.ReplaceAllString(s, " "))
}

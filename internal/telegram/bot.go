package telegram

import (
	"fmt"

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

func (t *TelegramClient) SendPost(title, summary, imageURL, sourceURL string) error {
	if t == nil || t.bot == nil {
		return nil // No-op if disabled
	}

	// Format processing
	// Telegram Caption Limit is 1024 chars.
	caption := fmt.Sprintf(""+
		"🚀 *NEW ARTICLE FOUND*\n"+
		"📝 *Title*: %s\n"+
		"🔗 *Source*: %s\n"+
		"📜 *Content*: %s\n",
		title, sourceURL, summary)

	if len(caption) > 1024 {
		caption = caption[:1021] + "..."
	}

	var msg tgbotapi.Chattable

	if imageURL != "" {
		// Verify image URL access first to avoid Telegram API error
		// Or just try. If failed, fallback to text.
		photo := tgbotapi.NewPhoto(t.chatID, tgbotapi.FileURL(imageURL))
		photo.Caption = caption
		photo.ParseMode = "Markdown"
		msg = photo
	} else {
		txt := tgbotapi.NewMessage(t.chatID, caption)
		txt.ParseMode = "Markdown"
		// Disable web page preview if you want, but likely we want it if no image
		msg = txt
	}

	_, err := t.bot.Send(msg)
	if err != nil {
		// If photo failed (e.g. invalid URL or bad format), try sending just text
		if imageURL != "" {
			t.l.Warnf(nil, "Telegram: Failed to send photo, falling back to text: %v", err)
			txt := tgbotapi.NewMessage(t.chatID, caption)
			txt.ParseMode = "Markdown"
			_, err = t.bot.Send(txt)
			return err
		}
		return err
	}

	return nil
}

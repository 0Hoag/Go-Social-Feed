package locale

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
)

type Locale struct{}

const (
	ViLanguage = "vi" // Vietnamese (Việt Nam)
	EnLanguage = "en" // English (Tiếng Anh))
	JaLanguage = "ja" // Japanese (Tiếng Nhật)
)

var SupportedLanguages = map[string]bool{
	ViLanguage: true,
	EnLanguage: true,
	JaLanguage: true,
}

var ErrLocaleNotFound = errors.New("locale not found")

var DefaultLanguage = ViLanguage

func GetLanguage(c *gin.Context) string {
	lang := c.GetHeader("Lang")
	if _, exists := SupportedLanguages[lang]; exists {
		return lang
	}
	return DefaultLanguage
}

func SetLocaleToContext(ctx context.Context, lang string) context.Context {
	return context.WithValue(ctx, Locale{}, lang)
}

func GetLocaleFromContext(ctx context.Context) (string, bool) {
	l, ok := ctx.Value(Locale{}).(string)
	return l, ok
}

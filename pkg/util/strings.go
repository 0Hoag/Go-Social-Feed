package util

import "github.com/microcosm-cc/bluemonday"

func TruncateHTMLString(html string, length int) string {
	p := bluemonday.StrictPolicy()
	sanitized := p.Sanitize(html)

	runes := []rune(sanitized)

	if len(runes) > length {
		return string(runes[:length]) + "..."
	}

	return string(runes)
}

package scanner

import (
	"github.com/hoag/go-social-feed/internal/core/scanner"
)

type ScanTokenInput struct {
	Token string `json:"token"`
}

type ScanTokenOutput struct {
	Network      string          `json:"network"`
	Name         string          `json:"name"`
	Address      string          `json:"address"`
	TrustScore   int             `json:"trust_score"`
	Issues       []scanner.Issue `json:"issues"`
	SafeFeatures []string        `json:"safe_features"`
}

package http

import (
	scan "github.com/hoag/go-social-feed/internal/core/scanner"
	"github.com/hoag/go-social-feed/internal/scanner"
)

type scannerTokenInput struct {
	Token    string `form:"token"`
	Language string `form:"lang"` // From query param or header
}

func (r scannerTokenInput) ToScanTokenInput() scanner.ScanTokenInput {
	lang := r.Language
	if lang == "" {
		lang = "en" // Default to English
	}
	return scanner.ScanTokenInput{
		Token:    r.Token,
		Language: lang,
	}
}

func (r scannerTokenInput) validate() error {
	if len(r.Token) == 0 {
		return errWrongBody
	}

	return nil
}

type issue struct {
	Type        scan.IssueType `json:"type"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Impact      int            `json:"impact"`
}

func toIssues(issues []scan.Issue) []issue {
	result := []issue{}
	for _, i := range issues {
		result = append(result, issue{
			Type:        i.Type,
			Name:        i.Name,
			Description: i.Description,
			Impact:      i.Impact,
		})
	}
	return result
}

type scannerTokenOutput struct {
	Network      string   `json:"network"`
	Name         string   `json:"name"`
	Address      string   `json:"address"`
	TrustScore   int      `json:"trust_score"`
	Issues       []issue  `json:"issues"`
	SafeFeatures []string `json:"safe_features"`
}

func (h handler) ToScanTokenOutput(token scanner.ScanTokenOutput) scannerTokenOutput {
	return scannerTokenOutput{
		Network:      token.Network,
		Name:         token.Name,
		Address:      token.Address,
		TrustScore:   token.TrustScore,
		Issues:       toIssues(token.Issues),
		SafeFeatures: token.SafeFeatures,
	}
}

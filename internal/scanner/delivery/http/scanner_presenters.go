package http

import (
	scan "github.com/hoag/go-social-feed/internal/core/scanner"
	"github.com/hoag/go-social-feed/internal/scanner"
)

type scannerTokenInput struct {
	Token string `json:"token"`
}

func (r scannerTokenInput) ToScanTokenInput() scanner.ScanTokenInput {
	return scanner.ScanTokenInput{
		Token: r.Token,
	}
}

func (r scannerTokenInput) validate() error {
	if len(r.Token) == 0 {
		return errWrongBody
	}

	return nil
}

type issue struct {
	Type        scan.IssueType
	Name        string
	Description string
	Impact      int
}

func toIssues(issues []scan.Issue) []issue {
	var result []issue
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

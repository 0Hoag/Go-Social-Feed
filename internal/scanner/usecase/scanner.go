package usecase

import (
	"context"
	"strings"

	"github.com/hoag/go-social-feed/internal/adapters/etherscan"
	scanDomain "github.com/hoag/go-social-feed/internal/scanner"
)

func (uc ScannerUC) ScanToken(ctx context.Context, input scanDomain.ScanTokenInput) (scanDomain.ScanTokenOutput, error) {
	query := input.Token
	address := query
	network := "eth"
	name := query

	// 1. Resolve Symbol if needed
	isAddress := (strings.HasPrefix(query, "0x") && len(query) == 42) || query == "0xMOCK"
	if !isAddress {
		foundAddr, _, foundName, err := uc.dexClient.SearchTopToken(query)
		if err != nil {
			uc.l.Errorf(ctx, "Token not found on DexScreener: %v", err)
			return scanDomain.ScanTokenOutput{}, scanDomain.ErrTokenNotFound
		}
		address = foundAddr
		name = foundName
	}

	// 2. Fetch Source Code (Try all networks like Telegram Bot)
	var sourceCode string
	var err error
	var networkFound string

	networks := []string{
		etherscan.NetworkETH,
		etherscan.NetworkBSC,
		etherscan.NetworkBase,
		etherscan.NetworkArbitrum,
		etherscan.NetworkPolygon,
	}

	for _, net := range networks {
		sourceCode, name, err = uc.ethClient.GetContractSource(net, address)
		if err == nil && sourceCode != "" {
			networkFound = net
			break
		}
	}

	if networkFound == "" {
		uc.l.Errorf(ctx, "scanner.usecase.scanner.ScanToken: source code not found on any network")
		return scanDomain.ScanTokenOutput{}, scanDomain.ErrSourceCodeNotFound
	}

	network = networkFound

	// 3. Analyze (pass language preference)
	result := uc.engine.Scan(sourceCode, address, input.Language)

	return scanDomain.ScanTokenOutput{
		Network:      network,
		Name:         name,
		Address:      address,
		TrustScore:   result.TrustScore,
		Issues:       result.Issues,
		SafeFeatures: result.SafeFeatures,
	}, nil
}

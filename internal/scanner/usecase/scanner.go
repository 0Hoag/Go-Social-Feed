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
	isAddress := strings.HasPrefix(query, "0x") && len(query) == 42
	if !isAddress {
		foundAddr, foundNet, foundName, err := uc.dexClient.SearchTopToken(query)
		if err != nil {
			uc.l.Errorf(ctx, "Token not found: %v", err)
			return scanDomain.ScanTokenOutput{}, scanDomain.ErrScanToken
		}
		address = foundAddr
		network = foundNet
		name = foundName
	}

	// 2. Fetch Source Code
	var sourceCode string
	var err error

	if isAddress {
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
				network = net
				break
			}
		}
	} else {
		sourceCode, name, err = uc.ethClient.GetContractSource(network, address)
	}

	if err != nil || sourceCode == "" {
		uc.l.Errorf(ctx, "scanner.usecase.scanner.ScanToken: %v", err)
		return scanDomain.ScanTokenOutput{}, scanDomain.ErrScanToken
	}

	// 3. Analyze
	result := uc.engine.Scan(sourceCode, address)

	return scanDomain.ScanTokenOutput{
		Network:      network,
		Name:         name,
		Address:      address,
		TrustScore:   result.TrustScore,
		Issues:       result.Issues,
		SafeFeatures: result.SafeFeatures,
	}, nil
}

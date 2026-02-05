package usecase

import (
	"github.com/hoag/go-social-feed/internal/adapters/dexscreener"
	"github.com/hoag/go-social-feed/internal/adapters/etherscan"
	"github.com/hoag/go-social-feed/internal/core/scanner"
	pkgLog "github.com/hoag/go-social-feed/pkg/log"
)

type ScannerUC struct {
	l         pkgLog.Logger
	engine    *scanner.Engine
	dexClient *dexscreener.Client
	ethClient *etherscan.Client
}

func New(l pkgLog.Logger, engine *scanner.Engine, dex *dexscreener.Client, eth *etherscan.Client) *ScannerUC {
	return &ScannerUC{
		l:         l,
		engine:    engine,
		dexClient: dex,
		ethClient: eth,
	}
}

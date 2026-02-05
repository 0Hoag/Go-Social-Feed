package dexscreener

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/hoag/go-social-feed/internal/adapters/etherscan"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

// Response structures matching DexScreener API
type SearchResponse struct {
	Pairs []Pair `json:"pairs"`
}

type Pair struct {
	ChainId   string    `json:"chainId"`
	BaseToken Token     `json:"baseToken"`
	Liquidity Liquidity `json:"liquidity"`
	Volume    Volume    `json:"volume"`
}

type Token struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
}

type Liquidity struct {
	Usd float64 `json:"usd"`
}

type Volume struct {
	H24 float64 `json:"h24"`
}

// SearchTopToken finds the best matching token for a query (Symbol or Name)
// Returns: address, network, name, error
func (c *Client) SearchTopToken(query string) (string, string, string, error) {
	url := fmt.Sprintf("https://api.dexscreener.com/latest/dex/search?q=%s", query)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return "", "", "", fmt.Errorf("dexscreener request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", "", fmt.Errorf("dexscreener status: %d", resp.StatusCode)
	}

	var result SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", "", "", fmt.Errorf("failed to decode json: %w", err)
	}

	if len(result.Pairs) == 0 {
		return "", "", "", fmt.Errorf("no token found for query: %s", query)
	}

	// Filter and Sort: prioritize high liquidity and specific chains
	// Supported chains in our bot: ethereum, bsc, base, arbitrum, polygon
	supportedChains := map[string]string{
		"ethereum": etherscan.NetworkETH,
		"bsc":      etherscan.NetworkBSC,
		"base":     etherscan.NetworkBase,
		"arbitrum": etherscan.NetworkArbitrum,
		"polygon":  etherscan.NetworkPolygon,
	}

	// 1. Filter only supported chains
	var candidates []Pair
	for _, p := range result.Pairs {
		if _, ok := supportedChains[p.ChainId]; ok {
			candidates = append(candidates, p)
		}
	}

	if len(candidates) == 0 {
		return "", "", "", fmt.Errorf("found tokens but not on supported chains (ETH, BSC, BASE)")
	}

	// 2. Sort by Liquidity USD descending (finding the 'real' token)
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].Liquidity.Usd > candidates[j].Liquidity.Usd
	})

	topMatch := candidates[0]

	// Map DexScreener chainID to our internal network ID
	network := supportedChains[topMatch.ChainId]

	return topMatch.BaseToken.Address, network, topMatch.BaseToken.Name, nil
}

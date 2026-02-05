package etherscan

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	NetworkETH      = "eth"
	NetworkBSC      = "bsc"
	NetworkBase     = "base"
	NetworkArbitrum = "arbitrum"
	NetworkPolygon  = "polygon"
)

type Client struct {
	apiKeys    map[string]string
	baseURLs   map[string]string
	httpClient *http.Client
}

type SourceCodeResponse struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Result  json.RawMessage `json:"result"`
}

type ContractSource struct {
	SourceCode   string `json:"SourceCode"`
	ContractName string `json:"ContractName"`
	ABI          string `json:"ABI"`
}

func NewClient(apiKeys map[string]string) *Client {
	return &Client{
		apiKeys: apiKeys,
		baseURLs: map[string]string{
			NetworkETH:      "https://api.etherscan.io/v2/api",
			NetworkBSC:      "https://api.bscscan.com/api",
			NetworkBase:     "https://api.basescan.org/api",
			NetworkArbitrum: "https://api.arbiscan.io/api",
			NetworkPolygon:  "https://api.polygonscan.com/api",
		},
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetContractSource fetches the Solidity source code and name for a given contract address on a specific network
func (c *Client) GetContractSource(network, address string) (string, string, error) {
	apiKey, ok := c.apiKeys[network]
	if !ok || apiKey == "" {
		return "", "", fmt.Errorf("no api key for network: %s", network)
	}

	baseURL, ok := c.baseURLs[network]
	if !ok {
		return "", "", fmt.Errorf("unsupported network: %s", network)
	}

	// MOCK MODE
	if apiKey == "MOCK" {
		fmt.Println("⚠️  Running in MOCK MODE (Simulated Response)")
		return `// SPDX-License-Identifier: MIT
		pragma solidity ^0.8.0;
		contract MockRiskyToken {
			function mint(address to, uint256 amount) public {}
		}`, "MockRiskyToken", nil
	}

	// Different chains might have slightly different v2/v1 query params, but standard getsourcecode is usually consistant.
	// Etherscan V2 uses chainid param. BscScan/BaseScan might still use V1 or ignore it.
	// For safety, we try the standard V1-style appended with chainid if it's ETH v2, or just standard for others.

	var url string
	if network == NetworkETH {
		url = fmt.Sprintf("%s?chainid=1&module=contract&action=getsourcecode&address=%s&apikey=%s", baseURL, address, apiKey)
	} else {
		// BSC and Base standard endpoints
		url = fmt.Sprintf("%s?module=contract&action=getsourcecode&address=%s&apikey=%s", baseURL, address, apiKey)
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return "", "", fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("failed to read response body: %w", err)
	}

	var parsedResp SourceCodeResponse
	if err := json.Unmarshal(body, &parsedResp); err != nil {
		return "", "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	if parsedResp.Status == "0" {
		var errorMsg string
		_ = json.Unmarshal(parsedResp.Result, &errorMsg)
		return "", "", fmt.Errorf("etherscan API error (%s): %s - %s", network, parsedResp.Message, errorMsg)
	}

	var results []ContractSource
	if err := json.Unmarshal(parsedResp.Result, &results); err != nil {
		return "", "", fmt.Errorf("api error (structure mismatch?): %w - raw: %s", err, string(parsedResp.Result))
	}

	if len(results) == 0 {
		return "", "", fmt.Errorf("no source code found")
	}

	return results[0].SourceCode, results[0].ContractName, nil
}

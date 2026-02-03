package etherscan

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	apiKey     string
	httpClient *http.Client
	baseURL    string
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

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:  apiKey,
		baseURL: "https://api.etherscan.io/v2/api",
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetContractSource fetches the Solidity source code for a given contract address
func (c *Client) GetContractSource(address string) (string, error) {
	// MOCK MODE for testing without API Key
	if c.apiKey == "MOCK" {
		fmt.Println("⚠️  Running in MOCK MODE (Simulated Etherscan Response)")
		return `
		// SPDX-License-Identifier: MIT
		pragma solidity ^0.8.0;

		contract MockRiskyToken {
			mapping(address => uint256) public balances;
			mapping(address => bool) public blacklist;
			address public owner;

			constructor() {
				owner = msg.sender;
				balances[owner] = 1000000;
			}

			// Critical: Rug Pull Risk (Uncapped Mint)
			function mint(address to, uint256 amount) public {
				balances[to] += amount;
			}

			// Critical: Honeypot Risk (Blacklist)
			function setBlacklist(address user, bool value) public {
				require(msg.sender == owner);
				blacklist[user] = value;
			}

			function transfer(address to, uint256 amount) public {
				require(!blacklist[msg.sender], "You are blacklisted!");
				balances[msg.sender] -= amount;
				balances[to] += amount;
			}
		}
		`, nil
	}

	url := fmt.Sprintf("%s?chainid=1&module=contract&action=getsourcecode&address=%s&apikey=%s", c.baseURL, address, c.apiKey)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var parsedResp SourceCodeResponse
	if err := json.Unmarshal(body, &parsedResp); err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	if parsedResp.Status == "0" {
		var errorMsg string
		_ = json.Unmarshal(parsedResp.Result, &errorMsg)
		return "", fmt.Errorf("etherscan API error: %s - %s", parsedResp.Message, errorMsg)
	}

	var results []ContractSource
	if err := json.Unmarshal(parsedResp.Result, &results); err != nil {
		return "", fmt.Errorf("etherscan API error (v2 structure mismatch?): %w - raw: %s", err, string(parsedResp.Result))
	}

	if len(results) == 0 {
		return "", fmt.Errorf("no source code found")
	}

	return results[0].SourceCode, nil
}

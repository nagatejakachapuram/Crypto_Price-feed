package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type CoinGeckoResponse struct {
	Usd float64 `json:"usd"`
}

func fetchCurrentPrice(crypto, apiKey string) (float64, error) {
	// Construct the API URL
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd", crypto)

	// Make the HTTP request
	response, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("failed to make request: %w", err)
	}
	defer response.Body.Close()

	// Check for HTTP errors
	if response.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("non-200 response status: %s", response.Status)
	}

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response body: %w", err)
	}

	// Log the response body for debugging
	fmt.Println("API Response:", string(body))

	// Unmarshal the JSON response
	var result map[string]CoinGeckoResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// Extract the price
	coinResponse, ok := result[crypto]
	if !ok {
		return 0, fmt.Errorf("cryptocurrency not found in response")
	}

	return coinResponse.Usd, nil
}

func main() {
	crypto := "bitcoin" // Example cryptocurrency
	apiKey := "c9d78726-9358-4724-bc8c-b8072f94d4f9"
	price, err := fetchCurrentPrice(crypto, apiKey)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("The current price of %s is $%.2f\n", crypto, price)
}

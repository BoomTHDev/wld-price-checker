package main

import (
	"encoding/json"
	"net/http"
)

type CoinData struct {
	Data struct {
		WLD struct {
			Quote struct {
				THB struct {
					Price float64 `json:"price"`
				} `json:"THB"`
			} `json:"quote"`
		} `json:"WLD"`
	} `json:"data"`
}

func getWorldcoinPrice(apiKey string) (float64, error) {
	url := "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest?symbol=WLD&convert=THB"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("X-CMC_PRO_API_KEY", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	coinPrice := CoinData{}
	if err := json.NewDecoder(resp.Body).Decode(&coinPrice); err != nil {
		return 0, err
	}

	return coinPrice.Data.WLD.Quote.THB.Price, nil
}

package repository

import (
	"encoding/json"
	"net/http"

	"github.com/boomthdev/wld-price-cheker/entities"
)

type coinRepositoryImpl struct {
	apiKey string
}

func NewCoinRepositoryImpl(apiKey string) CoinRepository {
	return &coinRepositoryImpl{apiKey: apiKey}
}

func (r *coinRepositoryImpl) GetWorldcoinPrice() (float64, error) {
	url := "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest?symbol=WLD&convert=THB"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("X-CMC_PRO_API_KEY", r.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	coinPrice := entities.WorldCoin{}
	if err := json.NewDecoder(resp.Body).Decode(&coinPrice); err != nil {
		return 0, err
	}

	return coinPrice.Data.WLD.Quote.THB.Price, nil
}

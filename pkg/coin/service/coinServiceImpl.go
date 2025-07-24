package service

import (
	_coinRepository "github.com/boomthdev/wld-price-cheker/pkg/coin/repository"
	"github.com/boomthdev/wld-price-cheker/pkg/custom"
)

type coinServiceImpl struct {
	coinRepository _coinRepository.CoinRepository
}

func NewCoinServiceImpl(coinRepository _coinRepository.CoinRepository) CoinService {
	return &coinServiceImpl{coinRepository: coinRepository}
}

func (s *coinServiceImpl) GetWorldcoinPrice() (float64, *custom.AppError) {
	price, err := s.coinRepository.GetWorldcoinPrice()
	if err != nil {
		return 0, custom.ErrIntervalServer("Failed to get Worldcoin price", err)
	}

	return price, nil
}

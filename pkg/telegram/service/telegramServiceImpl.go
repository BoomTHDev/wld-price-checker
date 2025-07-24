package service

import (
	"fmt"

	_coinRepository "github.com/boomthdev/wld-price-cheker/pkg/coin/repository"
	"github.com/boomthdev/wld-price-cheker/pkg/custom"
	_telegramRepository "github.com/boomthdev/wld-price-cheker/pkg/telegram/repository"
)

type telegramServiceImpl struct {
	telegramRepo   _telegramRepository.TelegramRepository
	coinRepository _coinRepository.CoinRepository
}

func NewTelegramService(telegramRepo _telegramRepository.TelegramRepository, coinRepository _coinRepository.CoinRepository) TelegramService {
	return &telegramServiceImpl{telegramRepo: telegramRepo, coinRepository: coinRepository}
}

func (s *telegramServiceImpl) SendPriceUpdate() *custom.AppError {
	price, err := s.coinRepository.GetWorldcoinPrice()
	if err != nil {
		return custom.ErrIntervalServer("Failed to get Worldcoin price", err)
	}
	message := fmt.Sprintf("🚀 Worldcoin ราคาแตะ %.2f บาทแล้ว", price)
	if err := s.telegramRepo.SendTelegramNotification(message); err != nil {
		return custom.ErrIntervalServer("Failed to send Telegram notification", err)
	}
	return nil
}

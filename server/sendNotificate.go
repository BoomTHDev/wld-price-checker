package server

import (
	"log"
	"time"

	_coinRepository "github.com/boomthdev/wld-price-cheker/pkg/coin/repository"
	_coinService "github.com/boomthdev/wld-price-cheker/pkg/coin/service"
	_telegramRepository "github.com/boomthdev/wld-price-cheker/pkg/telegram/repository"
	_telegramService "github.com/boomthdev/wld-price-cheker/pkg/telegram/service"
)

const targetPrice = 36.0

var lastNotified time.Time

func (s *fiberServer) sendNotificate() {
	for {
		coinRepository := _coinRepository.NewCoinRepositoryImpl(s.conf.CoinEnv.APIKey)
		coinService := _coinService.NewCoinServiceImpl(coinRepository)

		price, err := coinService.GetWorldcoinPrice()
		if err != nil {
			log.Printf("failed to get price: %v\n", err)
			return
		}

		if price >= targetPrice && time.Since(lastNotified) >= 60*time.Second {
			telegramRepository := _telegramRepository.NewTelegramRepository(s.conf.TelegramEnv.BotToken, s.conf.TelegramEnv.ChatID)
			telegramService := _telegramService.NewTelegramService(telegramRepository, coinRepository)
			if err := telegramService.SendPriceUpdate(); err != nil {
				log.Printf("failed to send telegram: %v\n", err)
			} else {
				log.Println("Notification sent!")
				lastNotified = time.Now()
			}
		}

		time.Sleep(60 * time.Second)
	}
}

package server

import (
	_coinRepository "github.com/boomthdev/wld-price-cheker/pkg/coin/repository"
	_telegramController "github.com/boomthdev/wld-price-cheker/pkg/telegram/controller"
	_telegramRepository "github.com/boomthdev/wld-price-cheker/pkg/telegram/repository"
	_telegramService "github.com/boomthdev/wld-price-cheker/pkg/telegram/service"
)

func (s *fiberServer) initTelegramRouter() {
	telegramRepository := _telegramRepository.NewTelegramRepository(s.conf.TelegramEnv.BotToken, s.conf.TelegramEnv.ChatID)
	coinRepository := _coinRepository.NewCoinRepositoryImpl(s.conf.CoinEnv.APIKey)
	telegramService := _telegramService.NewTelegramService(telegramRepository, coinRepository)
	_ = _telegramController.NewTelegramController(telegramService)
}

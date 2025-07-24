package server

import (
	_coinController "github.com/boomthdev/wld-price-cheker/pkg/coin/controller"
	_coinRepository "github.com/boomthdev/wld-price-cheker/pkg/coin/repository"
	_coinService "github.com/boomthdev/wld-price-cheker/pkg/coin/service"
)

func (s *fiberServer) initCoinRouter() {
	coinRepository := _coinRepository.NewCoinRepositoryImpl(s.conf.CoinEnv.APIKey)
	coinService := _coinService.NewCoinServiceImpl(coinRepository)
	coinController := _coinController.NewCoinController(coinService)

	coinRouter := s.app.Group("/coin")
	coinRouter.Get("/price", coinController.GetWorldcoinPrice)
}

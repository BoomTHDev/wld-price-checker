package service

import "github.com/boomthdev/wld-price-cheker/pkg/custom"

type CoinService interface {
	GetWorldcoinPrice() (float64, *custom.AppError)
}

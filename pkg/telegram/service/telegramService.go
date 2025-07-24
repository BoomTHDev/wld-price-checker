package service

import (
	"github.com/boomthdev/wld-price-cheker/pkg/custom"
)

type TelegramService interface {
	SendPriceUpdate() *custom.AppError
}

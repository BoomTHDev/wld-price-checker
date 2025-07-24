package controller

import (
	"github.com/boomthdev/wld-price-cheker/pkg/telegram/service"
)

type TelegramController struct {
	telegramService service.TelegramService
}

func NewTelegramController(telegramService service.TelegramService) *TelegramController {
	return &TelegramController{telegramService: telegramService}
}

func (c *TelegramController) SendPriceUpdate() error {
	return c.telegramService.SendPriceUpdate()
}

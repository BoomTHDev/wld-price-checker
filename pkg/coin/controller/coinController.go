package controller

import (
	_coinService "github.com/boomthdev/wld-price-cheker/pkg/coin/service"
	"github.com/gofiber/fiber/v2"
)

type CoinController struct {
	CoinService _coinService.CoinService
}

func NewCoinController(coinService _coinService.CoinService) CoinController {
	return CoinController{CoinService: coinService}
}

func (c *CoinController) GetWorldcoinPrice(ctx *fiber.Ctx) error {
	price, appErr := c.CoinService.GetWorldcoinPrice()
	if appErr != nil {
		return appErr
	}
	return ctx.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"WLD": fiber.Map{
				"price": price,
			},
		},
	})
}

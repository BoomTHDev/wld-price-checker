package server

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/boomthdev/wld-price-cheker/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type fiberServer struct {
	app  *fiber.App
	conf *config.Config
}

var (
	once           sync.Once
	serverInstance *fiberServer
)

func NewFiberServer(conf *config.Config) *fiberServer {
	fiberApp := fiber.New(fiber.Config{
		IdleTimeout: time.Second * time.Duration(conf.Server.Timeout),
	})

	once.Do(func() {
		serverInstance = &fiberServer{
			app:  fiberApp,
			conf: conf,
		}
	})

	return serverInstance
}

func (s *fiberServer) Start() {
	s.app.Use(logger.New())
	s.app.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(s.conf.Server.AllowOrigins, ","),
		AllowMethods:     "GET",
		AllowHeaders:     "Origin,Content-Type,Accept",
		AllowCredentials: false,
	}))

	s.app.Get("/health-check", s.healthCheck)
	s.initCoinRouter()
	s.initTelegramRouter()
	go s.sendNotificate()

	s.app.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": fmt.Sprintf("Sorry, endpoint %s %s not found.", ctx.Method(), ctx.Path()),
		})
	})

	s.httpListening()
}

func (s *fiberServer) httpListening() {
	url := fmt.Sprintf(":%d", s.conf.Server.Port)
	if err := s.app.Listen(url); err != nil {
		log.Fatalf("Failed to start server: %s\n", err.Error())
	}
}

func (s *fiberServer) healthCheck(ctx *fiber.Ctx) error {
	return ctx.SendString("OK")
}

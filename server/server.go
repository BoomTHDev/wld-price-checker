package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/boomthdev/wld-price-cheker/config"
	"github.com/go-co-op/gocron"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type fiberServer struct {
	app  *fiber.App
	conf *config.Config
	hub *Hub
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
		hub := NewHub()
		serverInstance = &fiberServer{
			app:  fiberApp,
			conf: conf,
			hub:  hub,
		}
		// Start the hub
		go hub.Run()
	})

	return serverInstance
}

func (s *fiberServer) setupCron() {
	cron := gocron.NewScheduler(time.UTC)
	apiURL := fmt.Sprintf("%s/health-check", os.Getenv("API_URL"))
	if apiURL == "" {
		log.Println("API_URL is not set")
		return
	}
	fmt.Println(apiURL)

	cron.Every("14m").Do(func() {
		log.Println("Running scheduled task at", time.Now().Format("2006-01-02 15:04:05"))

		client := http.Client{
			Timeout: 30 * time.Second,
		}

		resp, err := client.Get(apiURL)
		if err != nil {
			log.Println("Error running scheduled task:", err)
			return
		}

		defer resp.Body.Close()
	})

	cron.StartAsync()
}

func (s *fiberServer) Start() {
	s.app.Use(logger.New())
	s.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// WebSocket upgrade middleware
	s.app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	s.app.Get("/ws/price", websocket.New(s.handleWebSocket))
	s.app.Get("/health-check", s.healthCheck)
	s.initCoinRouter()
	s.initTelegramRouter()

	s.app.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": fmt.Sprintf("Sorry, endpoint %s %s not found.", ctx.Method(), ctx.Path()),
		})
	})
	go s.sendNotificate()
	go s.setupCron()
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

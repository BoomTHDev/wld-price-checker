package config

import (
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		Server      *Server
		CoinEnv     *CoinEnv
		TelegramEnv *TelegramEnv
	}

	Server struct {
		Port         int           `validate:"required"`
		AllowOrigins []string      `validate:"required"`
		Timeout      time.Duration `validate:"required"`
	}

	CoinEnv struct {
		APIKey string `validate:"required"`
	}

	TelegramEnv struct {
		BotToken string `validate:"required"`
		ChatID   string `validate:"required"`
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func ConfigGetting() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}

		configInstance = &Config{
			Server:      &Server{},
			CoinEnv:     &CoinEnv{},
			TelegramEnv: &TelegramEnv{},
		}

		port, err := strconv.Atoi(getEnv("SERVER_PORT", "8080"))
		if err != nil {
			panic("Invalid SERVER_PORT in .env file")
		}

		timeout, err := time.ParseDuration(getEnv("SERVER_TIMEOUT", "5s"))
		if err != nil {
			panic("Invalid SERVER_TIMEOUT in .env file")
		}

		configInstance.Server = &Server{
			Port:         port,
			AllowOrigins: strings.Split(getEnv("SERVER_ALLOW_ORIGINS", ""), ","),
			Timeout:      timeout,
		}

		configInstance.CoinEnv = &CoinEnv{
			APIKey: getEnv("COIN_API_KEY", ""),
		}

		configInstance.TelegramEnv = &TelegramEnv{
			BotToken: getEnv("TELEGRAM_BOT_TOKEN", ""),
			ChatID:   getEnv("TELEGRAM_CHAT_ID", ""),
		}

		validating := validator.New()
		if err := validating.Struct(configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}

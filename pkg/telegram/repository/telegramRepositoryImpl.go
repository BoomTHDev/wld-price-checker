package repository

import (
	"fmt"
	"net/http"
	"net/url"
)

type telegramRepositoryImpl struct {
	BotToken string
	ChatID   string
}

func NewTelegramRepository(botToken, chatID string) TelegramRepository {
	return &telegramRepositoryImpl{
		BotToken: botToken,
		ChatID:   chatID,
	}
}

func (r *telegramRepositoryImpl) SendTelegramNotification(message string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", r.BotToken)
	data := url.Values{
		"chat_id": {r.ChatID},
		"text":    {message},
	}

	resp, err := http.PostForm(apiURL, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message: %s", resp.Status)
	}

	return nil
}

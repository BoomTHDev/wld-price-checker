package repository

type TelegramRepository interface {
	SendTelegramNotification(message string) error
}

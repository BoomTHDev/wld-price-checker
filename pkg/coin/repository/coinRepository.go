package repository

type CoinRepository interface {
	GetWorldcoinPrice() (float64, error)
}

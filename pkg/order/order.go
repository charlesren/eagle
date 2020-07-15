package order

import "time"

type Order struct {
	OrderID           string
	StockName         string
	StockCode         string
	OrderType         orderType
	TransactionPrice  float64
	TransactionVolume int64
	TransactionTime   time.Time
	TransactionAmount float64
	ShareholderCode   string
}

type orderType string

const (
	Buying  orderType = "Buying"
	Selling orderType = "Selling"
)

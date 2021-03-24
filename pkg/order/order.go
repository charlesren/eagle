package order

import "time"

type Order struct {
	StockCode           string
	StockName           string
	OrderType           string
	TransactionDate     time.Time
	TransactionDateTime time.Time
	TransactionPrice    float64
	TransactionVolume   int64
	TransactionAmount   float64
	TransactionID       string
	OrderID             string
	ShareholderCode     string
}


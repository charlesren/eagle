package order

import "time"

type Order struct {
	OrderID         string
	StockName       string
	StockCode       string
	Flag            string
	Price           float64
	Number          int64
	Time            time.Time
	ShareholderCode string
}

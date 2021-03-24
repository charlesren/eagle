package order

type Order struct {
	StockCode           string
	StockName           string
	OrderType           string
	TransactionDate     string
	TransactionDateTime string
	TransactionPrice    float64
	TransactionVolume   int64
	TransactionAmount   float64
	TransactionID       string
	OrderID             string
	ShareholderCode     string
}

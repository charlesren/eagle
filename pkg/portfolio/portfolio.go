package portfolio

// Capital ...
type Capital struct{
	TotalCapital     float64
	TotalMarketValue float64
	AvailableCapital float64
	TotalPaperProfit float64
}
//StockCode is code of stock
type StockCode string

// Interest ...
type Interest struct {
	StockName         string
	MarketValue       string
	MarketPrice       float64
	CostSPrice        float64
	TotalQuantity     int64
	AvailableQuantity int64
	PaperProfit       float64
	PaperProfitRate   float64
	TodayProfit       float64
	TodayProfitRate   float64
	PositionRate   float64
}

// OpenInterests 
type OpenInterests map(StcokCode)Interest
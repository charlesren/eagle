package capital

//CapitalStock ...
type CapitalOpenInterests struct {
	TotalCapital     string
	TotalMarketValue string
	AvailableCapital string
	TotalPaperProfit float64
	OpenInterest    map(StcokCode)Detail
}
//StockCode is code of stock
type StockCode string

//Detail ...
type Detail struct {
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

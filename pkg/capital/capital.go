package capital

//CapitalStock ...
type CapitalOpenInterests struct {
	TotalCapital     string
	TotalMarketValue string
	AvailableCapital string
	OpenInterests    map(StcokCode)Detail
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
	StockholderCode   string
}

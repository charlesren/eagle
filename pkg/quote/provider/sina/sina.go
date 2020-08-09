package sina

import quotesina "github.com/charlesren/sina"

//QuoteMetrics is realtime metrics of stock
type QuoteMetrics struct {
	High float64
	Low  float64
}

//Get func fetch lastest data from sina
func Get(stockCode string) QuoteMetrics {
	var sqm QuoteMetrics
	tmpData := quotesina.GetData(stockCode)
	sqm.High = tmpData.High
	sqm.Low = tmpData.Low
	return sqm
}

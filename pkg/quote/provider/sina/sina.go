package sina

import quotesina "github.com/charlesren/sina"

//QuoteMetrics is realtime metrics of stock
type QuoteMetrics struct {
	Open float64
	High float64
	Low  float64
	Date string
}

//Get func fetch lastest data from sina
func Get(stockCode string) QuoteMetrics {
	var sqm QuoteMetrics
	tmpData := quotesina.GetData(stockCode)
	sqm.Open = tmpData.Open
	sqm.High = tmpData.High
	sqm.Low = tmpData.Low
	sqm.Date = tmpData.Date
	return sqm
}

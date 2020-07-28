package app

import quotesina "github.com/charlesren/sina"

//SinaQuoteMetrics is realtime metrics of stock
type SinaQuoteMetrics struct {
	Open float64
	High float64
	Low  float64
	Date string
}

//Get func fetch lastest data from sina
func Get(stockCode string) SinaQuoteMetrics {
	var sqm SinaQuoteMetrics
	tmpData := quotesina.GetData(stockCode)
	sqm.Open = tmpData.Open
	sqm.High = tmpData.High
	sqm.Low = tmpData.Low
	sqm.Date = tmpData.Date
	return sqm
}

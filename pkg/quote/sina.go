package sina

import quotesina "github.com/charlesren/sina"

//SinaQuoteMetrics is realtime metrics of stock
type SinaQuoteMetrics struct {
	Open     float64
	Now      float64
	High     float64
	Low      float64
	Turnover int
	Volume   float64
	Date     string
	Time     string
}

func Get(stockCode string) SinaQuoteMetrics {
	var sqm SinaQuoteMetrics
	tmpData := quotesina.GetData(stcokCode)
	sqm.Open = tmpData.Open
	sqm.Now = tmpData.Now
	sqm.High = tmpData.High
	sqm.Low = tmpData.Low
	sqm.Turnover = tmpData.Turnover
	sqm.Volume = tmpData.Volume
	sqm.Date = tmpData.Date
	sqm.Time = tmpData.Time
	return sqm
}

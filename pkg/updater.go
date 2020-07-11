package updater
type BookedStocks interface{
	AddStock([]string) BookedStocks
	DeleteStock([]string) BookedStocks
    AddMetricsSet(StockMetrics) BookedStocks
    DeleteMetricsSet(StockMetrics) BookedStocks
}
type StockList struct(
	Name map[string]struct{}
	MetricsSet []StockMetrics
)
type StockMetrics interface{
	GetMetrics() struct{} bool
	}
type MovingAverageMetrics struct{
	20Day string
	60Day string
	250Day string
}
func（MovingAeragetMetrics)   GetMetrics struct{} bool {
	return  MovingAverageMetrics
}
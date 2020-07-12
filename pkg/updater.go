package updater
type StockSet interface{
	AddStock(string) StockSet
	DeleteStock(string) StockSet
}
type StockList struct(
	Name map[string]struct{}
	Lock sync.Mutex
)
type DefaultMetrics struct{
	Open       float64
	Close      float64
	Now        float64
	High       float64
	Low        float64
	Volume     float64
}
func(sl *StockList) AddStock(stock string) {
sl.Lock.Lock
defer sl.Lock.UnLock
if _,ok := sl.Name[stock] ; ok {
fmt.Println(stock,"already exists!")
}
else {
	sl.Name[stock] = struct{}
}
}
func(sl *StockList) DeleteStock(stock string) {
sl.Lock.Lock
defer sl.Lock.UnLock
delete（sl.Name，stock)
}
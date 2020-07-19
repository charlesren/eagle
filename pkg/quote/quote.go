package quote

import (
	"fmt"
	"sync"
)
type Quote interface(
	Get()
)
//StockSet is a set of stocks
type StockSet interface {
	AddStock(string) StockSet
	DeleteStock(string) StockSet
}

//StockList implement interface StockSet
type StockList struct {
	Name map[string]struct{}
	Lock sync.Mutex
}

//DefaultMetrics is realtime metrics of stock
type DefaultMetrics struct {
	Open   float64
	Close  float64
	Now    float64
	High   float64
	Low    float64
	Volume float64
}

func (sl *StockList) AddStock(stock string) {
	sl.Lock.Lock()
	defer sl.Lock.Unlock()
	if _, ok := sl.Name[stock]; ok {
		fmt.Println(stock, "already exists!")
	} else {
		sl.Name[stock] = struct{}{}
	}
}
func (sl *StockList) DeleteStock(stock string) {
	sl.Lock.Lock()
	defer sl.Lock.Unlock()
	delete(sl.Name, stock)
}

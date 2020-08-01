package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/charlesren/eagle/pkg/quote/provider/sina"
)

func main() {
	stockName := os.Args[1]
	if stockName == "" {
		fmt.Println("stock code is nil")
		return
	}
	interval, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("wrong interval")
		return
	}
	var data sina.QuoteMetrics
	for {
		data = sina.Get(stockName)
		fmt.Println(data)
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

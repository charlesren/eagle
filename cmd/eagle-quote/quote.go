package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/charlesren/eagle/cmd/eagle-quote/app"
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
	}
	var data app.SinaQuoteMetrics
	for {
		data = app.Get(stockName)
		fmt.Println(data)
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

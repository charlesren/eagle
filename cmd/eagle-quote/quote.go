package main

import (
	"fmt"

	"github.com/charlesren/eagle/cmd/eagle-quote/app"
)

func main() {
	var data app.SinaQuoteMetrics
	data = app.Get("000027.SZ")
	fmt.Println(data)
}

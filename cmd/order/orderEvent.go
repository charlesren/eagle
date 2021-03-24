package main

import (
	"context"
	"encoding/csv"
	//"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/charlesren/eagle/pkg/order"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	cehttp "github.com/cloudevents/sdk-go/v2/protocol/http"
)

func main() {
	var orders []order.Order
	f, err := os.Open("tradeOrder.csv")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)

	for {
		o := order.Order{}
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		o.StockCode = line[0]
		o.StockName = line[1]
		o.OrderType = line[4]
		td, _ := time.Parse("20060102", line[2])
		o.TransactionDate = td.Format("2006-01-02")
		o.TransactionDateTime = o.TransactionDate + " " + line[3]
		o.TransactionPrice, _ = strconv.ParseFloat(line[5], 64)
		o.TransactionVolume, _ = strconv.ParseInt(line[6], 10, 64)
		o.TransactionAmount, _ = strconv.ParseFloat(line[7], 64)
		o.TransactionID = line[8]
		o.OrderID = line[9]
		o.ShareholderCode = line[10]
		orders = append(orders, o)
	}
	fmt.Println(orders)
	/*
		jsonOd, err := json.Marshal(orders)
		if err != nil {
			log.Fatalf("Create order JSON failed, %v\n", err.Error)
		}
		fmt.Println(string(jsonOd))
	*/
	ctx := cloudevents.ContextWithTarget(context.Background(), "http://localhost:8080/")

	p, err := cloudevents.NewHTTP()
	if err != nil {
		log.Fatalf("failed to create protocol: %s", err.Error())
	}

	c, err := cloudevents.NewClient(p, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	e := cloudevents.NewEvent()
	e.SetType("com.cloudevents.sample.sent")
	e.SetSource("https://github.com/cloudevents/sdk-go/v2/samples/httpb/sender")
	e.SetData(cloudevents.ApplicationJSON, orders)
	res := c.Send(ctx, e)
	if cloudevents.IsUndelivered(res) {
		log.Printf("Failed to send: %v", res)
	} else {
		var httpResult *cehttp.Result
		cloudevents.ResultAs(res, &httpResult)
		log.Printf("Sent with status code %d", httpResult.StatusCode)
	}
}

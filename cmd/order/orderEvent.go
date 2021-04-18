package main

import (
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
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

func main() {
	const orderSub = "orderSub"
	// Connect Options.
	opts := []nats.Option{}
	opts = append(opts, nats.Name("eagle"))

	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	subj := orderSub

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

		e := cloudevents.NewEvent()
		e.SetID(uuid.New().String())
		e.SetType("eagle.order.sent")
		e.SetTime(time.Now())
		e.SetSource("https://github.com/charlesren/eagle/pkg/order/orderEnevt")
		e.SetData(cloudevents.ApplicationJSON, o)
		fmt.Println(e)
		msg := []byte(e.String())
		nc.Publish(subj, msg)
		nc.Flush()
		if err := nc.LastError(); err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Published [%s] : '%s'\n", subj, msg)
		}
	}
}

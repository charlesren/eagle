package main

import (
	//"context"
	"fmt"
	"os"

	"encoding/csv"
	//"github.com/charlesren/eagle/pkg/order"
	"io"
	//cloudevents "github.com/cloudevents/sdk-go/v2"
	//cehttp "github.com/cloudevents/sdk-go/v2/protocol/http"
	//"log"
)

func main() {
	tradeFile := "tradeOrder.csv"
	//orders := []order.Order{}
	f, err := os.Open(tradeFile)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	reader := csv.NewReader(f)

	for {
		//o := order.Order{}
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println(line)
	}

	/*
		var targetURL string = "http://localhost:8080/"
		ctx := cloudevents.ContextWithTarget(context.Background(), targetURL)

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
		e.SetData(cloudevents.ApplicationJSON, map[string]interface{}{
			"id":      i,
			"message": "Hello, World!",
		})

		res := c.Send(ctx, e)
		if cloudevents.IsUndelivered(res) {
			log.Printf("Failed to send: %v", res)
		} else {
			var httpResult *cehttp.Result
			cloudevents.ResultAs(res, &httpResult)
			log.Printf("Sent %d with status code %d", i, httpResult.StatusCode)
		}
	*/
}

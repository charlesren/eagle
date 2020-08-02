package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	quotev1 "github.com/charlesren/eagle/pkg/api/quote/v1"
	"github.com/charlesren/eagle/pkg/quote/provider/sina"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

var (
	scheme = runtime.NewScheme()
)

func init() {
	_ = quotev1.AddToScheme(scheme)
}
func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	namespace := "default"
	fmt.Println(namespace)

	// creates the clientset
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	fmt.Println(client)
	drGVR := schema.GroupVersionResource{
		Group:    "quote.eagle.io",
		Version:  "v1",
		Resource: "dailyranges",
	}
	fmt.Println(drGVR)

	dr, err := client.Resource(drGVR).Namespace("default").Get("sh600519", v1.GetOptions{})
	fmt.Println(dr)
	fmt.Println(quotev1.GroupVersion)
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
	//for {
	data = sina.Get(stockName)
	fmt.Println(data)
	time.Sleep(time.Duration(interval) * time.Second)
	//}
}

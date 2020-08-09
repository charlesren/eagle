package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	quotev1 "github.com/charlesren/eagle/pkg/api/quote/v1"
	"github.com/charlesren/eagle/pkg/quote/provider/sina"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeconfig string

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "path to kubernetes config file")
	flag.Parse()
}
func main() {
	var config *rest.Config
	var err error
	if kubeconfig == "" {
		log.Printf("using in-cluster configuration")
		config, err = rest.InClusterConfig()
	} else {
		log.Printf("using configuration from '%s'", kubeconfig)
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}

	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	drGVR := schema.GroupVersionResource{
		Group:    "quote.eagle.io",
		Version:  "v1",
		Resource: "dailyranges",
	}
	drGVRClient := client.Resource(drGVR)
	//list drs
	drlistObj, err := drGVRClient.List(metav1.ListOptions{})
	if err != nil {
		fmt.Println("list dr error")
		log.Fatal(err)
	}
	fmt.Println("drlistObj is ", drlistObj)

	drList := &quotev1.DailyrangeList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(drlistObj.UnstructuredContent(), drList)
	if err != nil {
		panic(err)
	}
	for _, v := range drList.Items {
		fmt.Println("dr : ", v.Spec)
	}

	//get dr
	drObj, err := drGVRClient.Namespace("default").Get(os.Args[1], metav1.GetOptions{})
	if err != nil {
		fmt.Println("get dr error")
		log.Fatal(err)
	}
	fmt.Println("drObj is ", drlistObj)
	dr := &quotev1.Dailyrange{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(drObj.UnstructuredContent(), dr)
	if err != nil {
		panic(err)
	}
	fmt.Println("dr : ", dr.Spec)

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

package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/nats-io/nats.go"
)

func printMsg(m *nats.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'", i, m.Subject, string(m.Data))
}
func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to:%s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatalf("Exiting: %v", nc.LastError())
	}))
	return opts
}
func main() {
	const orderSub = "orderSub"
	// Connect Options.
	opts := []nats.Option{}
	opts = append(opts, nats.Name("eagle"))
	opts = setupConnOptions(opts)
	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL, opts...)
	if err != nil {
		log.Fatal(err)
	}
	//defer nc.Close()
	subj := orderSub
	i := 0
	nc.Subscribe(subj, func(msg *nats.Msg) {
		i += 1
		fmt.Println(i)
		printMsg(msg, i)
	})
	nc.Flush()
	//nc.Publish(subj, []byte("1"))

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on [%s]", subj)
	runtime.Goexit()

}

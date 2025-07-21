package utils

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

func CheckStan(natsServer string) bool {
	opts := []nats.Option{
		nats.Name("stan-check"),
		nats.Timeout(2 * time.Second),
		nats.ErrorHandler(func(nc *nats.Conn, sub *nats.Subscription, err error) {
			log.Printf("NATS async error: %v\n", err)
		}),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			log.Printf("Disconnected due to error: %v\n", err)
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			log.Println("Reconnected to NATS!")
		}),
	}

	nc, err := nats.Connect(natsServer, opts...)
	if err != nil {
		log.Println("Couldn't conn to server:", err)
		return false
	}
	defer nc.Close()
	// assumes "test-cluster" is the ID, todo: make dynamic
	sc, err := stan.Connect("test-cluster", "stan-check", stan.NatsConn(nc), stan.ConnectWait(2*time.Second))
	if err != nil {
		log.Println("No stans: ", err)
		return false
	}
	defer sc.Close()
	log.Println("stan loona or something (Connected, STAN exists)")

	return true
}

package main

import (
	"etum-dev/deez/utils"
	"flag"
	"fmt"
)

var (
	server = flag.String("server", "127.0.0.1:4222", "Target NATS server address (e.g., 11.11.11.11:4222 )")
)

var flagserver string

// brute streamname

// prio: brute user
func ParseFlags() {

}

func main() {
	flag.Parse()

	if utils.CheckStan(*server) {
		fmt.Println("STAN enabled")
	} else {
		fmt.Println("STAN disabled or needs auth")
	}
	utils.BruteHardcoded(*server)
}

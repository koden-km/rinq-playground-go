package main

import (
	"fmt"

	"github.com/rinq/rinq-go/src/rinq/amqp"
)

func main() {
	fmt.Println("yo dawg")

	peer, err := amqp.DialEnv()
	if err != nil {
		panic(err)
	}

	peer.GracefulStop()
}

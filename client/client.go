package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/rinq/rinq-go/src/rinq"
	"github.com/rinq/rinq-go/src/rinq/amqp"
)

func main() {
	fmt.Println("Client...")

	rand.Seed(time.Now().UnixNano())

	peer, err := amqp.DialEnv()
	if err != nil {
		panic(err)
	}

	ses := peer.Session()

	fmt.Println("Calling command 1...")
	out := rinq.NewPayload(123)
	in, err := ses.Call(context.Background(), "yolo", "getFunky", out)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Got: ", in)
	}

	fmt.Println("Calling command 2...")
	out = rinq.NewPayload(456)
	in, err = ses.Call(context.Background(), "yolo", "gotJazz", out)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Got: ", in)
	}

	peer.GracefulStop()
}

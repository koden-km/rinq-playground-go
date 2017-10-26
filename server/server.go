package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/rinq/rinq-go/src/rinq"
	"github.com/rinq/rinq-go/src/rinq/amqp"
)

func main() {
	fmt.Println("Server...")

	rand.Seed(time.Now().UnixNano())

	peer, err := amqp.DialEnv()
	if err != nil {
		panic(err)
	}

	ch := make(chan os.Signal)

	signal.Notify(ch, os.Interrupt)

	peer.Listen("yolo", fooHandler)

	select {
	case <-peer.Done():

	case <-ch:
		peer.GracefulStop()
		<-peer.Done()
	}

	err = peer.Err()
	if err != nil {
		panic(err)
	}
}

func fooHandler(ctx context.Context, req rinq.Request, res rinq.Response) {
	fmt.Println("Handle: ", req.Command)
	defer req.Payload.Clone()

	if req.Command == "getFunky" {
		res.Done(req.Payload)
	} else {
		// out := rinq.NewPayload("Bad command or filename. (A)bort, (R)etry or (I)gnore?")
		// defer out.Close()
		// res.Done(out)

		UnknownCommand := "unknown.command"
		res.Fail(UnknownCommand, "YOU ARE RONG!")
	}
}

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rinq/rinq-go/src/rinq/ident"
)

var history = map[uint16]struct{}{}

func main() {
	time.Now()
	// rand.Seed(time.Now().UnixNano())
	// rand.Seed(123)
	// This seed had failed during regular make tests
	rand.Seed(1509506564)

	count := 0
	for {
		id := ident.NewPeerID()
		count++

		if _, ok := history[id.Rand]; ok {
			panic(fmt.Sprintf("repeated %s after %d iterations", id, count))
		}

		history[id.Rand] = struct{}{}
	}
}

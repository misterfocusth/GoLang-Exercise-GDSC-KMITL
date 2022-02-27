package exercise

import (
	"context"
	"fmt"
)

func send(ch chan<- int, cancel context.CancelFunc) {
	defer cancel()
	for i := 0; i < 100; i++ {
		ch <- 1
	}
}

func receive(ch <-chan int, ctx context.Context) {
	var value int
	for {
		select {
		case val := <-ch:
			value += val
		case <-ctx.Done():
			fmt.Println(value)
			return
		}
	}
}

func Exercise7() {
	ch := make(chan int)
	bgCtx := context.Background()
	ctx, cancel := context.WithCancel(bgCtx)

	go send(ch, cancel)
	receive(ch, ctx)
}

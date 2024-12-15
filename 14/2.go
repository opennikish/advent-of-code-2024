package main

import (
	"adventofcode2024/lib"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	input, err := lib.GetInput(14)
	lib.Check(err)
	_ = input

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	c := time.Tick(500 * time.Millisecond)
	go func() {
		for {
			select {
			case v := <-c:
				clear()
				fmt.Println("tick", v)
			case <-ctx.Done():
				fmt.Println("Bye")
				return
			}
		}
	}()
	<-ctx.Done()
}

func clear() {
	fmt.Print("\033[H\033[2J")
}

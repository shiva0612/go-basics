package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// normal()
	advanced()
}

func advanced() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer func() {
		fmt.Println("end advanced")
		cancel()
	}()
	nch := make(chan int, 1)
	go func() {
		time.Sleep(5 * time.Second)
		nch <- 1
	}()

	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("ctrl+c is pressed")
	case <-nch:
		fmt.Println("from normal channel")
	}
	fmt.Println("main end")

}
func normal() {
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGINT)

	nch := make(chan int, 1)
	go func() {
		time.Sleep(5 * time.Second)
		nch <- 1
	}()

	time.Sleep(1 * time.Second)
	select {
	case <-sigch:
		fmt.Println("ctrl+c is pressed")
	case <-nch:
		fmt.Println("from normal channel")
	}
	fmt.Println("main end")
}

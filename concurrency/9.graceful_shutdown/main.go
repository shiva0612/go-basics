package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
SIGHUP	1	Hang up detected on controlling terminal or death of controlling process (closing the termianl)
SIGINT	2	Issued if the user sends an interrupt signal (Ctrl + C)
SIGQUIT	3	Issued if the user sends a quit signal (Ctrl + D)
SIGKILL	9	If a process gets this signal it must quit immediately and will not perform any clean-up operations
SIGTERM	15 Software termination signal (sent by kill by default)

os.Interrupt, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
*/
func main() {
	// normal()
	advanced()
}

func advanced() {
	//you can use this context in all goroutines spun from main, and cancel them as ctrl+c is pressed using context
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

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	for {
		select {
		case <-c:
			fmt.Println("interupt")
			return
		default:
			fmt.Print(".")

		}
		time.Sleep(500 * time.Millisecond)
	}
}

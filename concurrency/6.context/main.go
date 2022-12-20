package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg1 sync.WaitGroup
var wg2 sync.WaitGroup
var wgm sync.WaitGroup
var ch = make(chan int)

func main() {

	c := context.Background()
	ctx, cancel := context.WithCancel(c)

	wgm.Add(2)
	go func() {
		defer wgm.Done()
		wg1.Add(1)
		go test1(cancel)
		wg1.Wait()
	}()

	go func() {
		defer wgm.Done()
		wg2.Add(1)
		go test2(cancel)
		wg2.Wait()
	}()

	go wgmwait()

	select {
	case <-ctx.Done():
		fmt.Println("ctx cancelled")
	case <-ch:
		fmt.Println("all wg done")
	}
	fmt.Println("main done")

}
func wgmwait() {
	wgm.Wait()
	ch <- 1
}
func test1(cancel context.CancelFunc) {
	defer wg1.Done()
	for i := 0; i < 4; i++ {
		fmt.Println("1")
		if i == 1 {
			cancel()
		}
		time.Sleep(100 * time.Millisecond)
	}

}

func test2(cancel context.CancelFunc) {
	defer wg2.Done()
	for i := 0; i < 4; i++ {
		fmt.Println("2")
		time.Sleep(100 * time.Millisecond)
	}
}

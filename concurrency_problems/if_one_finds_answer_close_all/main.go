package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	wg           *sync.WaitGroup
	capabilities = []int{3, 4, 5, 6, 7}
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg = new(sync.WaitGroup)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go work(ctx, cancel, i)
	}

	//work is no longer needed - please stop all workers (no leakage)
	// time.Sleep(2 * time.Second)
	// cancel()
	wg.Wait()
}

func work(ctx context.Context, cancel context.CancelFunc, worker int) {

	defer wg.Done()

	done := make(chan bool)
	go func(worker int) {
		duration := time.Duration(capabilities[worker])
		time.Sleep(duration * time.Second)
		done <- true
	}(worker)

	select {
	case <-ctx.Done(): //if parent no longer needs it
		fmt.Println("someone (parent/a fellow)found answer ")
	case <-done:
		cancel() //if one finds answers no one needs to work
	}
}

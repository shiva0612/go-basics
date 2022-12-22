package main

import (
	"context"
	"fmt"
	"time"
)

var (
	sleep_time = []int{1, 2, 3, 4}
)

func main() {
	create_n_goroutines_wait_until_one_is_done_and_exit(4)
}

// if you want to pass channel to goroutines(you can)
func create_n_goroutines_wait_until_one_is_done_and_exit(n int) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < n; i++ {
		go work(i, cancel)
	}
	<-ctx.Done()
}

func work(i int, cancel context.CancelFunc) {
	defer cancel()
	time.Sleep(time.Duration(sleep_time[i]) * time.Second)
	fmt.Println("DONE by index: ", i)
}

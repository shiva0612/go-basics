package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x = int32(0)
var wg *sync.WaitGroup

func main() {
	wg = &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		// go incr_without_atomics()
		go incr_with_atomics()
	}
	wg.Wait()
	fmt.Println("x = ", x)
}

func incr_without_atomics() {
	defer wg.Done()
	x += 1
}

func incr_with_atomics() {
	defer wg.Done()
	atomic.AddInt32(&x, 1)
}

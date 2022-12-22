package main

import (
	"fmt"
	"sync"
)

var x = 0
var wg *sync.WaitGroup
var m = sync.Mutex{}

func main() {
	wg = &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		// go incr_without_mutex()
		go incr_with_mutex()
	}
	wg.Wait()
	fmt.Println("x = ", x)
}

func incr_without_mutex() {
	defer wg.Done()
	x += 1
}

func incr_with_mutex() {
	defer wg.Done()
	m.Lock()
	x += 1
	m.Unlock()
}

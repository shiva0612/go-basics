package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func f(v *uint32, wg *sync.WaitGroup) {
	for i := 0; i < 3000; i++ {
		atomic.AddUint32(v, 1)
	}
	wg.Done()
}

func main() {
	var v uint32 = 42
	var wg sync.WaitGroup
	wg.Add(2)
	go f(&v, &wg)
	go f(&v, &wg)
	wg.Wait()

	fmt.Println(v)
}

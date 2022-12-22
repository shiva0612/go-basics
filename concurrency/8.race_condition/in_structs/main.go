package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Count int
	sync.Mutex
}

var (
	p  = &Person{}
	wg sync.WaitGroup
)

func main() {

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incr()
	}
	wg.Wait()
	fmt.Println("count = ", p.Count)
}
func incr() {

	defer wg.Done()
	p.Lock()
	p.Count += 1
	p.Unlock()
}

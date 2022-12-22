package main

import (
	"fmt"
	"sync"
)

var (
	wg   = &sync.WaitGroup{}
	once = &sync.Once{}
)

func sync_once() {
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go work()
	}
	wg.Wait()
	fmt.Println("main DONE")

}

func work() {
	defer wg.Done()
	once.Do(func() {
		fmt.Println("in work")
	})
}

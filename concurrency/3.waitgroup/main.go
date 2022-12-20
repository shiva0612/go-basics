package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go for1()
	go for2()

	wg.Wait()
	fmt.Println("main end")
}

func for1() {
	defer wg.Done()
	for i := 0; i < 5; i++ {

		fmt.Println(".")
		time.Sleep(100 * time.Millisecond)

	}
}
func for2() {
	defer wg.Done()
	for i := 0; i < 5; i++ {

		fmt.Println("*")
		time.Sleep(100 * time.Millisecond)

	}
}

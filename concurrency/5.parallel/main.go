package main

import (
	"fmt"
	"sync"
	"time"
)

var wg1 sync.WaitGroup
var wg2 sync.WaitGroup
var wgm sync.WaitGroup

func main() {

	wgm.Add(2)
	go func() {
		defer wgm.Done()
		wg1.Add(1)
		go test1()
		wg1.Wait()
	}()

	go func() {
		defer wgm.Done()
		wg2.Add(1)
		go test2()
		wg2.Wait()
	}()
	wgm.Wait()

}
func test1() {
	defer wg1.Done()
	for i := 0; i < 4; i++ {
		fmt.Println("1")
		time.Sleep(100 * time.Millisecond)
	}

}

func test2() {
	defer wg2.Done()
	for i := 0; i < 4; i++ {
		fmt.Println("2")
		time.Sleep(100 * time.Millisecond)
	}
}

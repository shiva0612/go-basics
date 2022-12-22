package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan int)
	c2 := make(chan int, 1)

	defer close(c1)
	defer close(c2)

	select_(c1)
	select_(c2)
	fmt.Println("main DONE")

}

func select_(ch chan int) {
	select {
	case <-time.After(time.Second):
		fmt.Println("waited for second")
	case ch <- 1:
		fmt.Println("had put values")
	default:
		fmt.Println("default")
	}
	fmt.Println("--select DONE--")

}

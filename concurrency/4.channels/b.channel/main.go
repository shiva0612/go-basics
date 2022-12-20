package main

import (
	"fmt"
	"time"
)

var bufCh = make(chan int, 2)
var Ch = make(chan int) //also works

func main() {
	usingBufCh()

}

func usingBufCh() {
	go for1()
	go for2()
	l := 0
	for a := range bufCh {
		l += 1
		if l == 2 {
			close(bufCh)
		}
		fmt.Println("a=", a)
	}
	fmt.Println("main end")
}

func for1() {
	for i := 0; i < 5; i++ {

		fmt.Println(".")
		time.Sleep(100 * time.Millisecond)

	}
	bufCh <- 1
}
func for2() {
	for i := 0; i < 5; i++ {

		fmt.Println("*")
		time.Sleep(100 * time.Millisecond)

	}
	bufCh <- 2
}

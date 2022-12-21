package main

import "fmt"

func UnBuf() {
	ch := make(chan int)
	go UnBufProduce(ch)
	// UnBufConsume(ch)
	UnBufConsumeForRange(ch)

}

func UnBufConsume(ch chan int) {
	for {
		n, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("n = ", n)
	}
}
func UnBufConsumeForRange(ch chan int) {
	for n := range ch {
		fmt.Println("n = ", n)
	}
}

func UnBufProduce(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

package main

import "fmt"

func UnBuf() {
	ch := make(chan int)
	go UnBufProduce(ch)
	// UnBufConsume(ch)
	// UnBufConsumeContinuous(ch)
	UnBufConsumeForRange(ch)

}

func UnBufConsumeContinuous(ch chan int) {
	i := 0
	for {
		n, _ := <-ch
		if i == 8 {
			break
		}
		fmt.Println("n = ", n)
		i++
	}
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

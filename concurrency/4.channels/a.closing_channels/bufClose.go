package main

import "fmt"

func BufEg() {
	ch := make(chan int, 3)
	BufProduce(ch)
	// BufConsume(ch)
	BufConsumeContinuous(ch)
	// BufConsumeForRange(ch)

}

func BufConsumeContinuous(ch chan int) {
	i := 0
	for {
		n, _ := <-ch
		if i == 5 {
			break
		}
		fmt.Println("n = ", n)
		i++
	}
}

func BufConsume(ch chan int) {
	for {
		n, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("n = ", n)
	}
}
func BufConsumeForRange(ch chan int) {
	for n := range ch {
		fmt.Println("n = ", n)
	}
}

func BufProduce(ch chan int) {
	for i := 0; i < cap(ch); i++ {
		ch <- i
	}
	close(ch)
}

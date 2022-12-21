package main

import "fmt"

func BufEg() {
	ch := make(chan int, 3)
	BufProduce(ch)
	BufConsume(ch)
	// BufConsumeForRange(ch)

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

// for range will keep on reading from channel until closed
// once closed
// if buffered channel:only read the remaining data and exit
// if unbuffered channel: will exit immediately (only 1 read and 1 write at a time happens this unbuffered channel)
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

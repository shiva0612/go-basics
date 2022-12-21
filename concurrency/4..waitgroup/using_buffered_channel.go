package main

import (
	"fmt"
	"time"
)

var (
	bufch chan int
)

func using_buffered_channel() {

	bufch = make(chan int, 2)
	go for11()
	go for22()

	buf_wait(2)
	fmt.Println("using_buffered_channel DONE")

}

func buf_wait(no_of_goroutines int) {
	go_routines_count := 0
	for {
		_, open := <-bufch
		if !open {
			break
		}
		go_routines_count += 1
		if go_routines_count == no_of_goroutines {
			close(bufch)
		}
	}
}

func for11() {
	defer func() {
		bufch <- 1
	}()
	for i := 0; i < 5; i++ {

		fmt.Print(".")
		time.Sleep(100 * time.Millisecond)

	}
}
func for22() {
	defer func() {
		bufch <- 1
	}()
	for i := 0; i < 5; i++ {

		fmt.Print("*")
		time.Sleep(100 * time.Millisecond)

	}
}

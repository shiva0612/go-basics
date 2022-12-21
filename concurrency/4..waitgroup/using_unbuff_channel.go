package main

import (
	"fmt"
	"time"
)

var (
	unbufch chan int
)

func using_unbuf_channel() {
	unbufch = make(chan int)
	go for111()
	go for222()

	unbuf_wait(2)
	fmt.Println("using_unbuf_channel DONE")

}
func unbuf_wait(no_of_goroutines int) {
	go_routines_done := 0
	for {
		_, open := <-unbufch
		if !open {
			break
		}
		go_routines_done += 1
		if go_routines_done == no_of_goroutines {
			close(unbufch)
		}
	}
}
func for111() {
	defer func() {
		unbufch <- 1
	}()

	for i := 0; i < 5; i++ {

		fmt.Print(".")
		time.Sleep(100 * time.Millisecond)

	}
}
func for222() {
	defer func() {
		unbufch <- 1
	}()

	for i := 0; i < 5; i++ {

		fmt.Print("*")
		time.Sleep(100 * time.Millisecond)

	}
}

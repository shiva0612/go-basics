package main

import (
	"fmt"
	"time"
)

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

// here since channel is getting populated in another goroutine
// in the main goroutine, in case statement, it will see if the channel has value if not, it will immediately
// go to another channel or default without waiting

func main() {

	ch := make(chan string)
	go process(ch)
	select {
	case v := <-ch:
		fmt.Println("received value: ", v)
		return
	default:
		fmt.Println("no value received")
	}
	fmt.Println("now checking for select")
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received")
		}
	}

}

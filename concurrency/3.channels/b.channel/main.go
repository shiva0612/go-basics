package main

import "fmt"

var bufCh = make(chan int, 2)
var Ch = make(chan int) //also works

func main() {
	basic_channel()
}

func basic_channel() {

	c := make(chan int)
	go normal(c)
	go write_only(c)
	go read_only(c)
}

// normal passing of channel
func normal(c chan int) {
	fmt.Println("chan c can  be used for read/write")

}

// passing channel as writeonly channel
func write_only(c chan<- int) {
	fmt.Println("chan c can only be used as write only channel")
	c <- 2
}

// passing channel as readonly channel
func read_only(c <-chan int) {
	fmt.Println("chan c can only be used as read only channel")
	<-c
}

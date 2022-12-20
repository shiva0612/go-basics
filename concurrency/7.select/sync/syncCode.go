package main

import (
	"fmt"
	"time"
)

func giveChan() chan int {
	time.Sleep(1 * time.Second)
	c := make(chan int, 1)
	c <- 1
	return c
}

func giveChan2() chan int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from ", r)
		}
	}()
	time.Sleep(1 * time.Second)
	c := make(chan int, 1)
	c <- 1
	panic("nothing")
	return c
}

// since while requesting we are calling in a sync way -> it will wait till func returns something
// it goes to default only when the func gives some error

func main() {
	select {
	case <-giveChan():
		fmt.Println("got channel")
	default:
		fmt.Println("in default")
	}

	i := 0
	for {

		select {
		case <-giveChan():
			fmt.Println("got channel")
		default:
			fmt.Println("in default")
		}
		i++
		if i > 2 {
			break
		}
	}

	fmt.Println("now we see default values since we get error in func call")

	select {
	case <-giveChan2():
		fmt.Println("got channel")
	default:
		fmt.Println("in default")
	}

	i = 0
	for {

		select {
		case <-giveChan2():
			fmt.Println("got channel")
		default:
			fmt.Println("in default")
		}
		i++
		if i > 2 {
			break
		}
	}
}

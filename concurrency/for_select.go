package main

import (
	"fmt"
	"time"
)

var (
	timer *time.Timer
	ch    chan int
)

func minit() {
	timer = time.NewTimer(time.Second)
	ch = make(chan int, 10)
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- 1
	}()
}
func main() {

	for_select_proper()
	for_select_return()
	for_select_break()
}

func for_select_break() {
	minit()
	for {
		time.Sleep(100 * time.Millisecond)
		select {
		case <-timer.C:
			fmt.Println("for_select_break: time...")
		case <-ch:
			fmt.Println("for_select_break: my ch...")
			break
		default:
			fmt.Println("for_select_break: default ...")
		}
	}
	fmt.Println("after for")

}

func for_select_return() {
	minit()
	for {
		time.Sleep(100 * time.Millisecond)
		select {
		case <-timer.C:
			fmt.Println("for_select_return: time...")
		case <-ch:
			fmt.Println("for_select_return: my ch...")
			return
		default:
			fmt.Println("for_select_return: default ...")
		}
	}
	fmt.Println("after for")
}

/*
u can use labeled break for:
composite for loops (for for for)
for select
switch
*/
func for_select_proper() {
	minit()

outerFor:
	for {
		time.Sleep(100 * time.Millisecond)
		select {
		case <-timer.C:
			fmt.Println("for_select_proper: time...")
		case <-ch:
			fmt.Println("for_select_proper: my ch...")
			break outerFor
		default:
			fmt.Println("for_select_proper: default ...")
		}
	}
	fmt.Println("for_select_proper: after for")

}

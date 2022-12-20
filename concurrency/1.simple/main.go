package main

import (
	"fmt"
	"time"
)

func main() {
	go for1()
	go for2()
	fmt.Println("main end")

}

func for1() {
	for {
		fmt.Println(".")
		time.Sleep(100 * time.Millisecond)
	}
}
func for2() {
	for {
		fmt.Println("*")
		time.Sleep(100 * time.Millisecond)
	}
}

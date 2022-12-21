package main

import (
	"fmt"
	"time"
)

func waiting_using_waitgroups() {
	wg.Add(2)
	go for1()
	go for2()

	wg.Wait()
	fmt.Println("waiting_using_waitgroups DONE")

}

func for1() {
	defer wg.Done()
	for i := 0; i < 5; i++ {

		fmt.Print(".")
		time.Sleep(100 * time.Millisecond)

	}
}
func for2() {
	defer wg.Done()
	for i := 0; i < 5; i++ {

		fmt.Print("*")
		time.Sleep(100 * time.Millisecond)

	}
}

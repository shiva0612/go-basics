package main

import (
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

/*
go routine above the limit will be ignored no error will be thrown
*/

func over_the_limit() {
	eg := &errgroup.Group{}
	eg.SetLimit(2)
	eg.TryGo(w1)
	eg.TryGo(w2)
	eg.TryGo(w3)

	if err := eg.Wait(); err != nil {
		fmt.Println("erro = ", err.Error())
	}
	fmt.Println("main DONE")
}

func w1() error {
	time.Sleep(3 * time.Millisecond)
	fmt.Println("in w1")
	return nil
}
func w2() error {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("in w2")
	return nil
}
func w3() error {
	time.Sleep(5 * time.Millisecond)
	fmt.Println("in w3")
	return nil
}

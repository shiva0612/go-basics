package main

import (
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

/*
eg.wait() just for all go routines to be completed and returns first non-nil error
*/

func one_go_returns_err() {
	eg := &errgroup.Group{}
	eg.SetLimit(3)
	eg.Go(w11)
	eg.Go(w22)
	eg.Go(w33)

	if err := eg.Wait(); err != nil {
		fmt.Println("erro = ", err.Error())
	}
	fmt.Println("main DONE")
}

func w11() error {
	time.Sleep(3 * time.Millisecond)
	fmt.Println("in w1")
	return nil
}
func w22() error {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("in w2")
	return nil
}
func w33() error {
	time.Sleep(5 * time.Millisecond)
	fmt.Println("in w3")
	return fmt.Errorf("im error")
}

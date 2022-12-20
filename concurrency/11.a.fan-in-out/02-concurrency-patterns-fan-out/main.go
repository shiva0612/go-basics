package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
)

var wg = &sync.WaitGroup{}

// What is the Fan-Out concurrency pattern?
// Breakup of one channel into multiple ones by distributing each value.
func main() {
	ch1, err := readCSV("file1.csv")
	if err != nil {
		panic(fmt.Errorf("Could not read file1 %v\n", err))
	}

	//-

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int, ch1 <-chan []string) {
			defer wg.Done()
			for v := range ch1 {
				fmt.Println("got ", v, " in", i)
			}
		}(i, ch1)
	}
	wg.Wait()

	fmt.Println("All completed, exiting")
}

func breakup(worker string, ch <-chan []string) chan struct{} {
	chE := make(chan struct{})

	go func() {
		for v := range ch {
			fmt.Println(worker, v)
		}

		close(chE)
	}()

	return chE
}

func readCSV(file string) (<-chan []string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("opening file %v\n", err)
	}

	ch := make(chan []string)

	cr := csv.NewReader(f)

	go func() {
		for {
			record, err := cr.Read()
			if err == io.EOF {
				close(ch)

				return
			}

			ch <- record
		}
	}()

	return ch, nil
}

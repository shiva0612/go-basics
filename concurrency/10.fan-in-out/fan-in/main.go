package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
)

// What is the Fan-In concurrency pattern?
// Consolidation of multiple channels into one channel by multiplexing each recieved value.
func main() {
	ch1, err := readCSV("file1.csv")
	if err != nil {
		panic(fmt.Errorf("Could not read file1 %v\n", err))
	}

	ch2, err := readCSV("file2.csv")
	if err != nil {
		panic(fmt.Errorf("Could not read file2 %v\n", err))
	}

	chM := merge_channels(ch1, ch2)

	for v := range chM {
		fmt.Println(v)
	}

	fmt.Println("All completed, exiting")
}
func merge_channels(cs ...<-chan []string) <-chan []string {
	merged_ch := make(chan []string)

	go func() {
		defer close(merged_ch)
		wg := &sync.WaitGroup{}
		wg.Add(len(cs))
		for i := 0; i < len(cs); i++ {
			go func(ch <-chan []string) {
				for row := range ch {
					merged_ch <- row
				}
			}(cs[i])
		}
		wg.Wait()
	}()

	return merged_ch
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

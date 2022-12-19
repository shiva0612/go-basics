package main

import (
	"fmt"
	"sync"
)

func main() {
	mfunc := func() {
		fmt.Println("executed")
	}

	once := &sync.Once{}

	wg := &sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			once.Do(mfunc)
		}()
	}
	wg.Wait()
	once.Do(mfunc)
	fmt.Println("DONE...")
}

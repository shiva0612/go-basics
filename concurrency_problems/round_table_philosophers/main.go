package main

import (
	"hash/fnv"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

// Number of philosophers is simply the length of this list.
var ph = []string{"Mark", "Russell", "Rocky", "Haris", "Root"}

const hunger = 3                // Number of times each philosopher eats
const think = time.Second / 100 // Mean think time
const eat = time.Second / 100   // Mean eat time

var fmt = log.New(os.Stdout, "", 0)

var dining sync.WaitGroup

func diningProblem(phName string, left, right *sync.Mutex) {
	fmt.Println(phName + " dined")
	h := fnv.New64a()
	h.Write([]byte(phName))
	rg := rand.New(rand.NewSource(int64(h.Sum64())))
	rSleep := func(t time.Duration) {
		time.Sleep(t/2 + time.Duration(rg.Int63n(int64(t))))
	}

	for h := hunger; h > 0; h-- {

		//locks forks and eat
		right.Lock()
		left.Lock()
		fmt.Println(phName + "eating")
		rSleep(eat)

		//unlocks forks and thiks
		right.Unlock()
		left.Unlock()
		fmt.Println(phName + "thinking")
		rSleep(think)
	}

	fmt.Println(phName, " satisfied and left")
	dining.Done()
}

func main() {

	dining.Add(5)

	f0 := &sync.Mutex{}
	fLeft := f0
	for i := 1; i < len(ph); i++ {
		fRight := &sync.Mutex{}
		go diningProblem(ph[i], fLeft, fRight)
		fLeft = fRight
	}
	go diningProblem(ph[0], f0, fLeft)

	dining.Wait()
}

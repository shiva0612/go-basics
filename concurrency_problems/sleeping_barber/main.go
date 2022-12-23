package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	sleeping = iota
	checking
	cutting
)

var stateLog = map[int]string{
	0: "Sleeping",
	1: "Checking",
	2: "Cutting",
}
var wg *sync.WaitGroup // Amount of potentional customers

type Barber struct {
	name string
	sync.Mutex
	state    int // Sleeping/Checking/Cutting
	customer *Customer
}

type Customer struct {
	name string
}

func (c *Customer) String() string {
	return fmt.Sprintf("%p", c)[7:]
}

func NewBarber() (b *Barber) {
	return &Barber{
		name:  "Sam",
		state: sleeping,
	}
}

// Barber goroutine
// Checks for customers
// Sleeps - wait for wakers to wake him up
func barber(b *Barber, wr chan *Customer, wakers chan *Customer) {
	for {
		b.Lock()
		// defer b.Unlock()
		b.state = checking
		b.customer = nil

		// checking the waiting room
		fmt.Printf("Checking waiting room: %d\n", len(wr))
		time.Sleep(time.Millisecond * 100)
		select {
		case c := <-wr:
			HairCut(c, b)
			// b.Unlock()
		default: // Waiting room is empty
			fmt.Printf("Barber goes to sleep \n")
			b.state = sleeping
			b.customer = nil
			b.Unlock() //unlock so that waker can wake him and and acquire barber lock so that no one else does before
			c := <-wakers
			b.Lock()
			fmt.Printf("Woken by %s\n", c)
			HairCut(c, b)
			// b.Unlock()
		}
	}
}

func HairCut(c *Customer, b *Barber) {
	b.state = cutting
	b.customer = c
	fmt.Printf("Cutting  %s hair\n", c)
	b.Unlock() // while cutting hair, barber is unlocked since, customers who came after sees him cut hair and waits in waiting room

	time.Sleep(time.Millisecond * 100)

	b.Lock() //we lock - make customer nil - unlock
	wg.Done()
	b.customer = nil
	b.Unlock()
}

// customer goroutine
// just fizzles out if it's full, otherwise the customer
// is passed along to the channel handling it's haircut etc
func customer(c *Customer, b *Barber, wr chan<- *Customer, wakers chan<- *Customer) {
	// arrive
	time.Sleep(time.Millisecond * 50)
	// lock the barber to check on status of barber
	b.Lock()
	fmt.Printf("Customer %s checks %s barber | room: %d, w %d - customer: %s\n",
		c, stateLog[b.state], len(wr), len(wakers), b.customer)
	switch b.state {
	case sleeping:
		//customer is put in waker list of size 1, if size if full
		select {
		case wakers <- c:
			//customer is put in waiting list,
			//if waiting list is also full, he cannot cut his hair today (bad luck)
		default:
			select {
			case wr <- c:
			default:
				wg.Done()
			}
		}
	case cutting:
		select {
		case wr <- c:
		default: // Full waiting room, leave shop
			wg.Done()
		}
	case checking:
		panic("Customer shouldn't check for the Barber when Barber is Checking the waiting room")
	}
	b.Unlock()
}

func main() {
	b := NewBarber()
	b.name = "Rocky"
	WaitingRoom := make(chan *Customer, 5) // 5 chairs
	Wakers := make(chan *Customer, 1)      // Only one waker at a time
	go barber(b, WaitingRoom, Wakers)

	time.Sleep(time.Millisecond * 100)
	wg = new(sync.WaitGroup)
	n := 10
	wg.Add(10)
	// Spawn customers
	for i := 0; i < n; i++ {
		time.Sleep(time.Millisecond * 50)
		c := new(Customer)
		go customer(c, b, WaitingRoom, Wakers)
	}

	wg.Wait()
	fmt.Println("No more customers for the day")
}

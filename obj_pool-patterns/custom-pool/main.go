package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

/*
 */
type Mpool struct {
	pool        *sync.Pool
	Active      int
	MaxActive   int
	ConnCreated int
	Capacity    int
	m           *sync.Mutex
}

type DBConn struct {
	Name string
}

func (db *DBConn) performTask() {
	fmt.Printf("performing task...%v\n", db)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("DONE")
}

func (mp *Mpool) init(capacity, intial, maxActive int) {
	mp.MaxActive = maxActive
	mp.pool = &sync.Pool{
		New: func() any {
			mp.ConnCreated += 1
			return &DBConn{Name: "default"}
		},
	}
	mp.m = &sync.Mutex{}

	for i := 0; i < intial; i++ {
		mp.pool.Put(&DBConn{Name: "initial"})
	}
	mp.ConnCreated = intial
}

func (mp *Mpool) Get() *DBConn {

	mp.m.Lock()
	defer mp.m.Unlock()

	if mp.Active >= mp.MaxActive {
		log.Println("max active connection reached")
		return nil
	}

	dbConn := mp.pool.Get().(*DBConn)
	mp.Active += 1
	return dbConn
}

func (mp *Mpool) IdleConn() int {
	mp.m.Lock()
	defer mp.m.Unlock()
	idleConn := mp.ConnCreated - mp.Active
	return idleConn
}

func (mp *Mpool) Put(dbConn *DBConn) {
	mp.m.Lock()
	defer mp.m.Unlock()

	mp.pool.Put(dbConn)
	mp.Active -= 1

}

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(20)

	mp := Mpool{}
	mp.init(10, 5, 8)

	for i := 0; i < 20; i++ {
		go func() {
			defer wg.Done()

			//max time for acquiring the connection
			timer := time.NewTimer(5 * time.Second)

			ch := make(chan *DBConn)
			go func() {
				for {
					//request for connection every 500 millisecond
					time.Sleep(500 * time.Millisecond)
					conn := mp.Get()
					if conn != nil {
						ch <- conn
						break
					}
				}
			}()

			for {
				select {
				case <-timer.C:
					fmt.Println("max wait time for acquiring conn obj reached...")
					return
				case con := <-ch:
					con.performTask()
					mp.Put(con)
					return
				}
			}

		}()
	}

	wg.Wait()
	fmt.Printf("conCreated: %v active:%v maxActive:%v idle:%v\n", mp.ConnCreated, mp.Active, mp.MaxActive, mp.IdleConn())

	fmt.Println("main DONE...")

}

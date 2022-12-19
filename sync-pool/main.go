package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type P struct {
	Name string
}

type S struct {
	Name string
}

func main() {
	pool := sync.Pool{
		New: func() any { return P{"default-p"} },
	}

	/*

		--------------
		get			put

		if no p=object present in the pool, it will return object returned from New function defined in sync.pool
		we cannot define the length of the pool
		it is kinda unlimited

	*/
	pool.Put(S{Name: "shiva"})
	printjson(pool.Get().(S))
	printjson(pool.Get().(P))

}

func printjson(cls any) {
	b, _ := json.MarshalIndent(cls, " ", "  ")
	fmt.Println(string(b))
}

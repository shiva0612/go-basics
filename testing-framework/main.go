package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	age  int
}

func main() {
	p1 := Person{"shiva", 23}
	p2 := Person{"shiva", 24}

	fmt.Println(p1 == p2)
	fmt.Println(reflect.DeepEqual(p1, p2))
}

package main

import "fmt"

type daytype int

const (
	weekend daytype = iota
	week
)

func (d daytype) String() string {
	switch d {
	case weekend:
		return "weekend"
	case week:
		return "week"
	}
	return ""
}

func enum_test() {
	// fmt.Println(week.String())
	fmt.Println(week)
}

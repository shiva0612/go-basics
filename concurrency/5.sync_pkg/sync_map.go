package main

import (
	"fmt"
	"sync"
)

func sync_map() {
	m := sync.Map{}
	m.LoadOrStore("name", "shiva")
	m.Store("age", "23")
	m.Range(func(key, value any) bool {
		fmt.Println(key, value)
		// if key.(string) == "name" {
		// 	return false
		// }
		return true
	})
	m.Delete("name")
}

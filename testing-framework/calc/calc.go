package calc

import "fmt"

func Add(a, b int) int {
	return a + b
}

func LimitAdd(a, b int) int {
	result := a + b
	if result > 10 {
		fmt.Println("heavy duty")
	}
	return result
}

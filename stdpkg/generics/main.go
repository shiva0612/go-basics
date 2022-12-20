package main

import "fmt"

type Number interface {
	int32 | int64 | float32 | float64
}

func genericFunc2[N Number](ip N) {
	fmt.Println(ip)
}

func genericFunc1[n int32 | float32 | int64](ip n) n {
	fmt.Println(ip)
	return ip
}

func main() {

	genericFunc2(int32(1))
	genericFunc2(1.1)

	fmt.Println(genericFunc1[float32](1.1)) //since 64 bit process defualt int and float are float64 and int64

}

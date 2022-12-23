package main

import "fmt"

type Number interface {
	int32 | int64 | float32 | float64
}

func genericFunc2[N Number](ip N) {
	fmt.Printf("%T - %v\n", ip, ip)
}

func genericFunc1[NN int32 | float32 | int64](ip NN) NN {
	fmt.Println(ip)
	return ip
}

func main() {

	genericFunc2(int32(1))
	genericFunc2(1.1)
	// genericFunc2("1") //wont compile

	fmt.Println(genericFunc1[float32](1.1)) //since 64 bit process defualt int and float are float64 and int64

}

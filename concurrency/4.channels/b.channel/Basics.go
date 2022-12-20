package main

func basic_channel() {

	c := make(chan int)
	go test(c)
	go write_test(c)
	go read_test(c)
}

func test(c chan int) {

}

func write_test(c chan<- int) {
	c <- 2
}

func read_test(c <-chan int) {
	<-c
}

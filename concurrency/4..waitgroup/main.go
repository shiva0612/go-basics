package main

import (
	"sync"
)

var wg sync.WaitGroup

func main() {
	waiting_using_waitgroups()
	using_buffered_channel()
	using_unbuf_channel()
}

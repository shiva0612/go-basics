package main

import (
	"context"
	"fmt"
)

/*

if cancelled -> ctx.Done() channel is closed -> ctx.Err() is no nil (mostly context cancelled)
if not cancelled -> ctx.Done() channel is still open -> ctx.Err() is nil


*/

/*
create context with value
then, use that as parent context
create context with deadline ...

values are still available in the context after cancellation as well
*/
func main() {

	c1 := context.WithValue(context.Background(), "name", "shiva")
	c2, cancel := context.WithCancel(c1)
	fmt.Println(c2.Value("name"))
	fmt.Println(c2.Err())
	cancel()
	fmt.Println(c2.Err())
	fmt.Println(c2.Value("name"))
}

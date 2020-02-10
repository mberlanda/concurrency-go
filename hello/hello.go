package hello

import "fmt"

// Hello illustrates single channel, single goroutine, one write, one read
func Hello() {
	// create new channel of type int
	ch := make(chan int)

	// start new anonymous goroutine
	go func() {
		// send 42 to channel
		ch <- 42
	}()
	// read from channel
	i := <-ch
	fmt.Printf("%d\n", i)
}

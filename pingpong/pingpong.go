package pingpong

import (
	"fmt"
	"time"
)

// PingPong implements a channel as a table of the ping-pong game
func PingPong() {
	var Ball int
	table := make(chan int)
	go player(table)
	go player(table)

	table <- Ball
	time.Sleep(1 * time.Second)
	b := <-table

	fmt.Printf("Last ball played: %d\n", b)
}

func player(table chan int) {
	for {
		ball := <-table
		fmt.Printf("Received ball: %d\n", ball)
		ball++
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}

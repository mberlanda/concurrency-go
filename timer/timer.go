package timer

import "time"

func timer(d time.Duration) <-chan int {
	c := make(chan int)
	go func() {
		time.Sleep(d)
		c <- 1
	}()
	return c
}

/*
Timer creates a channel, start goroutine which writes to this channel
after given duration and returns this channel to the caller of your func
*/
func Timer() {
	for i := 0; i < 24; i++ {
		c := timer(1 * time.Second)
		<-c
	}
}

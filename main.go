package main

import (
	"github.com/mberlanda/concurrency-go/workers"
)

func main() {
	// hello.Hello()
	// timer.Timer()
	// pingpong.PingPong()
	// fanin.FanIn()
	workers.LaunchPool()
}

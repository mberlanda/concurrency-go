package main

import (
	"github.com/mberlanda/concurrency-go/pingpong"
	"github.com/mberlanda/concurrency-go/fanin"
)

func main() {
	// hello.Hello()
	// timer.Timer()
	pingpong.PingPong()
	fanin.FanIn()
}

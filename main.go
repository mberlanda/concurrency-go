package main

import (
	"github.com/mberlanda/concurrency-go/hello"
	"github.com/mberlanda/concurrency-go/timer"
	"github.com/mberlanda/concurrency-go/pingpong"
)

func main() {
	hello.Hello()
	timer.Timer()
	pingpong.PingPong()
}

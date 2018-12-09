package main

import (
	"fmt"
	"time"
)

func ping(ch chan int) {
	for {
		<-ch
		fmt.Println("Ping")
		ch <- 0
	}
}

func pong(ch chan int) {
	for {
		<-ch
		fmt.Println("Pong")
		ch <- 0
	}
}

func main() {
	pingPong := make(chan int)
	go ping(pingPong)
	go pong(pingPong)
	pingPong <- 0
	time.Sleep(1 * time.Second)
}

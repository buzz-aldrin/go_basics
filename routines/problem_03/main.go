package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	boom := time.NewTicker(time.Second * 10)

	for {
		select {
		case <-ticker.C:
			fmt.Println("tick")
		case <-boom.C:
			fmt.Println("BOOM")
			return
		default:
			fmt.Println("blocked")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

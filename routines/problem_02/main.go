package main

import "fmt"

func fib(result, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case result <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quiting")
			return
		}
	}

}

func main() {
	result := make(chan int, 10)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-result)
		}
		quit <- 1
	}()

	fib(result, quit)
}

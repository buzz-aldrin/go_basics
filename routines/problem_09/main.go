package main

import (
	"fmt"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

func gen(arr ...int) <-chan int {
	pro := make(chan int)
	go func() {
		for _, v := range arr {
			pro <- v
		}
		close(pro)
	}()
	return pro
}

func sqr(pro <-chan int) <-chan int {
	con := make(chan int)
	go func() {
		for i := range pro {
			con <- i * i
		}
		close(con)
	}()
	return con
}

func merge(done chan struct{}, cons ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(con <-chan int) {
		defer wg.Done()
		for {
			select {
			case i := <-con:
				out <- i
			case _, ok := <-done:
				if !ok {
					return
				}
			}
		}
	}

	for _, con := range cons {
		wg.Add(1)
		go output(con)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	pro := gen(1, 2, 3, 4, 5, 6)
	done := make(chan struct{})
	out := merge(done, sqr(pro), sqr(pro))
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	close(done)
	time.Sleep(time.Second * 5)
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())

	filepath.Walk()
}

package main

import (
	"fmt"
	"sync"
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

func merge(cons ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(con <-chan int) {
		for i := range con {
			out <- i
		}
		wg.Done()
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
	for i := range merge(sqr(pro), sqr(pro), sqr(pro)) {
		fmt.Println(i)
	}
}

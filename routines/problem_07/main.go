package main

import "fmt"

func gen(arr ...int) <-chan int {
	in := make(chan int)
	go func() {
		for _, v := range arr {
			in <- v
		}
		close(in)
	}()
	return in
}

func sqr(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			out <- i * i
		}
		close(out)
	}()
	return out
}

func main() {
	for i := range sqr(gen(1, 2, 3, 4, 5, 6)) {
		fmt.Println(i)
	}
}

package main

import "fmt"

func add(arr []int, result chan int) {
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	result <- sum
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := make(chan int)

	go add(arr[:len(arr)/2], result)
	go add(arr[len(arr)/2:], result)

	x, y := <-result, <-result

	fmt.Println("sum: ", x+y)
}

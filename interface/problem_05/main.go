package main

import "fmt"

func main() {
	var i interface{}
	i = "hello"
	s := i.(string)
	fmt.Println(s)

	i = float64(0)
	f, ok := i.(float64)
	if ok {
		fmt.Println(f)
	}

	s = i.(string) // panic
	fmt.Println(s)
}

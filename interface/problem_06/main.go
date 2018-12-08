package main

import "fmt"

func do(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Printf("val:%d type:%T\n", i.(int), i)
	case bool:
		fmt.Printf("val:%t type:%T\n", i.(bool), i)
	case string:
		fmt.Printf("val:%s type:%T\n", i.(string), i)
	default:
		fmt.Printf("unknown type %T\n", i)
	}
}
func main() {
	var i interface{}
	i = 1
	do(i)
	i = true
	do(i)
	i = "hello"
	do(i)
	i = float64(1.2)
	do(i)
}

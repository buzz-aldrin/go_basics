package main

import (
	"fmt"
)

type Abser interface {
	Abs() float64
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var a Abser
	describe(a)
	a.Abs()
}

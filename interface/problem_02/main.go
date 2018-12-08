package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) String() string {
	return fmt.Sprintf("X:%.02f , Y:%.02f", v.X, v.Y)
}

func main() {
	var a Abser
	v := Vertex{X: 2, Y: 3}

	a = &v
	fmt.Println(a.String())
}

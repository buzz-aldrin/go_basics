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
	if v == nil {
		return float64(0)
	}

	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) String() string {
	return fmt.Sprintf("X:.02%f , Y:.02%f", v.X, v.Y)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var a Abser
	var v *Vertex
	a.Abs()
	a = v
	describe(a)
	fmt.Println(a.Abs())
	v = &Vertex{X: 2, Y: 3}
	a = v
	fmt.Println(a.Abs())
	describe(a)
}

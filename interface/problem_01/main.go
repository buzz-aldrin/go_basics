package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (mf MyFloat) Abs() float64 {
	if mf < 0 {
		return float64(-mf)
	}
	return float64(mf)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	mf := MyFloat(math.Sqrt2)
	v := Vertex{X: 2, Y: 3}

	a = mf
	fmt.Println(a.Abs())
	a = &v
	fmt.Println(a.Abs())

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = v
	fmt.Println(a.Abs())
}

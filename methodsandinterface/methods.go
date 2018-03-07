package methodsandinterface

import (
	"fmt"
	"math"
)

//Vertex needs comment
type Vertex struct {
	X, Y float64
}

//this is how you do it as a method
func (v Vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//ReturnMethod needs comment
func ReturnMethod() {
	v := Vertex{3, 4}
	fmt.Println(v.abs())
}

// this can be done as a normal function too
func absFunction(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//ReturnFunction comment
func ReturnFunction() {
	v := Vertex{3, 4}
	fmt.Println(absFunction(v))
}

type MyFloat float64

func (f MyFloat) absoluteFloat() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//ReturnFloat comment
func ReturnFloat() {
	f := MyFloat(-math.Sqrt(2))
	fmt.Println(f.absoluteFloat())
}

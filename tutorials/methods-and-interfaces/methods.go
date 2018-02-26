package methodsAndInterfaces

import (
	"fmt"
	"math"
)

//Vertex needs comment
type Vertex struct {
	X, Y float64
}

//Abs needs comment too! for export
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func ReturnMethod() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

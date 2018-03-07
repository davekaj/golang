package methodsandinterface

import (
	"fmt"
	"math"
)

func (v Vertex) abs2() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//the biggest thing to note here is the * infront of *Vertex. if this was
//not there, it wouldnt alter it by 10, and it would return 5 instead of 50
func (v *Vertex) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//PointerReceiver comment
func PointerReceiver() {
	v := Vertex{3, 4}
	v.scale(10)
	fmt.Println(v.abs2())
}

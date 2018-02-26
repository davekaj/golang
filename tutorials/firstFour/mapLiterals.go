package oneToFour

import (
	"fmt"
)

//Vertex should get imported from my other package

//map Literals are like struct literals. to me it just looks like these are literally typed, and thats all that is special about them

func mapLiterals() {
	var m = map[string]Vertex{
		"davestest": Vertex{
			23.2, 55.4,
		},
		"toronto_city": Vertex{
			5.4, -23.4,
		},
	}
	fmt.Println(m)

}

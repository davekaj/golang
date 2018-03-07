package firstfour

import (
	"fmt"
)

//Mutations comment
func Mutations() {

	//make a maping of stinrg to int
	//update ojne key value, and print
	//overwrite that, and print
	//delete it, and print
	// test if it is present, and print

	var m = make(map[string]int)

	m["answer"] = 42
	fmt.Println(m["answer"])

	m["answer"] = 48
	fmt.Println(m["answer"])

	delete(m, "answer")
	fmt.Println(m["answer"])

	v, ok := m["answer"]
	fmt.Println("The value: ", v, " is present? : ", ok)

}

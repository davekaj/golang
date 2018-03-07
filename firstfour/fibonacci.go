package firstfour

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
// func fibonacci() func(int) int {
// 	firstNumber := 0
// 	secondNumber := 1
// 	return func(f int) int {
// 		if f == 0 {
// 			return 0
// 		}
// 		if f == 1 {
// 			return 1
// 		}
// 		newNumber := secondNumber + firstNumber
// 		firstNumber = secondNumber
// 		secondNumber = newNumber
// 		return newNumber
// 	}
// }

//but this one skips the first two values !
//main thing i elarned is keeping similar assignments vars in the same line
func simpleFibo() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

//ReturnFib comment
func ReturnFib() {
	f := simpleFibo()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}

//we need to give it some recursion

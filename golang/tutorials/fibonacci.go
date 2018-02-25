package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	firstNumber := 0
	secondNumber := 1
	return func(f int) int {
		if f == 0 {
			return 0
		}
		if f == 1 {
			return 1
		}
		newNumber := secondNumber + firstNumber
		firstNumber = secondNumber
		secondNumber = newNumber
		return newNumber
	}
}

func returnFib() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f(i))
	}
}

//we need to give it some recursion

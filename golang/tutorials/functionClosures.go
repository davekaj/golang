package main

import "fmt"

func equation() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func testClosure() {
	pos, neg := equation(), equation()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

/*
the function equation() returns another function, which is the one func(x int)
this returned function CLOSES OVER the  x variable

then we assign positive and negative to the same equation, we now have two seperate instances of the returned func(x int), and they will act independantly
they each have their own closed x value which will be modifyied by whatever each one is passed

*/

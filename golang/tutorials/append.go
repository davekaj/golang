package main

import (
	"fmt"
)

func appendTest() {

	var s []int
	printAppendSlice(s)

	s = append(s, 0)
	printAppendSlice(s)

	s = append(s, 1)
	printAppendSlice(s)

	s = append(s, 2, 3)
	printAppendSlice(s)

	s = append(s, 55, 698)
	printAppendSlice(s)

	//it appears that everytime we double, we add double the capactity, then fill up, and then we will double again, i.e. 2, 4, 8, 16
	s = append(s, 4444, 111111, 9)
	printAppendSlice(s)

}

func printAppendSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

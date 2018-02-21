package main

import (
	"fmt"
)

//range returns the index and the value of a loop. you can toss one of the values away by using underscore

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
var kanyeWest = []string{"kanye", "is", "a", "good", "rapper"}

func testRange() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

}

func rangeWithUnderscore() {
	for _, v := range kanyeWest {
		fmt.Printf("Line 1: %s\n", v)
	}
}

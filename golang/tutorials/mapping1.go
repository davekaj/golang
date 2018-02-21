package main

import "fmt"

//Vertex - giving it a comment, because my linter requires it!
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

//make a map, fill it it, print it
func map1() {
	m = make(map[string]Vertex)
	m["DAVES-MAP"] = Vertex{54.44, 80.33}
	fmt.Println(m["DAVES-MAP"])

}

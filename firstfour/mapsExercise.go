package firstfour

import (
	"fmt"
	"strings"
)

func wordCount(s string) map[string]int {
	var stringSplit = strings.Fields(s)

	mappy := make(map[string]int)

	for _, v := range stringSplit {
		if _, ok := mappy[v]; ok == true {
			mappy[v]++
		} else {
			mappy[v] = 1
		}
	}
	return mappy
}

//CallWordCount comment
func CallWordCount() {
	var theWords = wordCount("I am learning go! Yes I am!")
	fmt.Println(theWords)
}

package main

import (
	"fmt"
)

// count the number of occurence of each 8 bit sequence in a given string
// return the result as
func countOccurences(s string) [256]uint {
	var occurences [256]uint

	for _, value := range s {
		occurences[value] += 1
	}
	return occurences
}

func main() {
	for k, v := range countOccurences("heeeeeeeeeeeeee") {
		if v > 0 {
			fmt.Println("%c : %u", k, v)
		}
	}
}

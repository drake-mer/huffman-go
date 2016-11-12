package counter

import (
	"fmt"
	"log"
	"os"
)

// count the number of occurence of each 8 bit sequence in a given string
// return the result as an array of unsigned integers
// with the given setup, the actual limitation for this code to work is
// that the size (in bytes) of the string does not exceed the capacity of
// an unsigned int on the system (2**32 ~ 1 Gb on a 32 bytes system)
//
func CountOccurences(data []byte) [256]uint {
	var occurences [256]uint

	for _, value := range data {
		occurences[value] += 1
	}
	return occurences
}

func MakeDict(dict [256]uint, fname string) {
	// write a file containing the frequency count for each possible value of
	// a byte
	dict_name := fname + ".dict"
	file, err := os.Create(dict_name)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range dict {
		fmt.Fprintln(file, v)
	}
}

package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/elijahbal/huffman/counter"
)

const Separator = " "

func MakeDict(Dict counter.Dict, fName string) {
	// write a file containing the frequency count for each possible value of
	// a byte

	// here you see, we are creating the file and we are writing into it in
	// a straightforward pattern. But this package content is devoted to one thing
	// count occurences ; it should not do IO.

	// A better design would be to pass a file pointer to this function instead
	// of a string

	dictName := fName + ".dict"
	file, err := os.Create(dictName)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range Dict {
		fmt.Fprintln(file, k, Separator, v)
	}
}

func ReadDict(fname string) counter.Dict {
	// this method takes as an input a standard filename and output
	// a dictionary
	FP, err := os.Open(fname + ".dict")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(FP)
	var Dictionary counter.Dict
	var Line string
	var Items []string
	for scanner.Scan() {
		Line = scanner.Text()
		Items = strings.SplitAfterN(Line, Separator, 1)
		Dictionary[byte(Items[0][0])], err = strconv.ParseUint(Items[1], 10, 64)
	}
	return Dictionary
}

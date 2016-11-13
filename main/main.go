package main

import (
	"log"
	"os"

	"github.com/elijahbal/huffman/counter"
)

func main() {
	file, err := os.Open(os.Args[1])
	data := make([]byte, 100000)
	if err != nil {
		log.Fatal(err)
	}
	count, err := file.Read(data)
	dictionary := counter.CountOccurences(data[:count])
	counter.MakeDict(dictionary, os.Args[1])
}

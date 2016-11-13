package counter

import (
	"sort"
)

type Pair struct {
	Key   byte
	Value uint64
}

// A slice of pairs that implements sort.Interface to sort by values
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

type Dict map[byte]uint64

// count the number of occurence of each 8 bit sequence in a given string
// return the result as an array of unsigned integers
// with the given setup, the actual limitation for this code to work is
// that the size (in bytes) of the string does not exceed the capacity of
// an unsigned int on the system (2**32 ~ 1 Gb on a 32 bytes system)
//
func CountOccurences(data []byte) Dict {
	var occurences Dict

	for _, value := range data {
		v, k := occurences[value]
		if k {
			occurences[value] = v + 1
		} else {
			occurences[value] = 1
		}
	}
	return occurences
}

func SortDict(data Dict) PairList {
	Result := make(PairList, len(data))
	for k, v := range data {
		Result = append(Result, Pair{k, v})
	}
	sort.Sort(Result)
	return Result
}

func MakeDict(data []byte) PairList {
	rawDict := CountOccurences(data)
	return SortDict(rawDict)
}

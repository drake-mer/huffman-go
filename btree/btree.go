// btree.go
package btree

import (
	"fmt"

	"github.com/elijahbal/huffman/counter"
)

type nodeContent byte

type bNode struct {
	left    *bNode
	right   *bNode
	content nodeContent
}

func mapTreeToFrequencyMap(input counter.PairList) bNode {
	var initTree bNode
}

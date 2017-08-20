package btree

import (
    "testing"
    "fmt"
)

func TestTreeBuild(t * testing.T){
    my := BinaryTree {
        left: nil,
        right: nil,
        content: 255,
    }
    length := TreeSize(my)
    if length != 0 {
        t.Error("length of empty tree is not zero")
    }
}


func TestTreeSize1(t * testing.T){
    my := BinaryTree {
        left: &BinaryTree{
            left: nil,
            right: nil,
            content: 0,
        },
        right: nil,
        content: 255,
    }
    length := TreeSize(my)
    if length != 1 {
        t.Error("length of empty tree is not zero")
    }
}


func BuildRightNodes(t* testing.T){
    length:=10
    my:=BinaryTree{content:0}
    for i:=0;i<10;i++ {
        addRightNode(my, 0)
    }
    treeLength := TreeSize(my)
    if length != treeLength {
        analyzeString := fmt.Sprint("number of nodes added : %u, length of tree : %u", length, treeLength)
        t.Error(analyzeString)
    }
}




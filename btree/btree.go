package btree

// simply define the types used to build the
// corresponding binary tree
// 
// this type definition is not very interesting
// because it works only on byte, and not on 
// words of varying length.
//
// It is nevertheless 

func max(a, b int) int {
    if a>b {
        return a
    }
    return b
}

type BinaryTree struct {
    left * BinaryTree  // pointer to a binary tree
    right * BinaryTree // pointer to a binary tree
    content byte  // content of the node
}

func addNode(btree BinaryTree, nodeContent byte, left bool){
    var rightNode * BinaryTree
    if left {
        rightNode = btree.right
    } else {
        rightNode = btree.left
    }
    for rightNode!=nil {
        btree = *rightNode
        if left {
            rightNode = btree.right
        } else {
            rightNode = btree.left
        }
    }
    newNodePointer := &BinaryTree{
        left: nil,
        right: nil,
        content: nodeContent,
    }
    if left {
        btree.left = newNodePointer
    } else {
        btree.right = newNodePointer
    }
}

func addLeftNode(btree BinaryTree, nodeContent byte){
    addNode(btree, nodeContent, true)
}

func addRightNode(btree BinaryTree, nodeContent byte){
    addNode(btree, nodeContent, false)
}

func TreeToMap(btree BinaryTree) map[byte]int {
    counter := make(map[byte]int)
    var f func(x BinaryTree) 
    f = func (btree BinaryTree) {
        counter[btree.content] = counter[btree.content] + 1
        if btree.left != nil {
            f(*btree.left)
        }
        if btree.right != nil {
            f(*btree.right)
        }
    }
    defer f(btree)
    return counter
}

func TreeSize(btree BinaryTree) int {
    var f func(_ BinaryTree) int
    f = func (btree BinaryTree) int{
        if btree.left == nil && btree.right == nil {
            return 0
        }
        if btree.left != nil && btree.right != nil {
            return (1 + max (f(*btree.left), f(*btree.right)))
        } else { 
            if btree.left != nil {
                return (1 + f(*btree.left))
            } else { 
                if btree.right != nil {
                    return (1 + f(*btree.right))
                } 
            } 
        }
        return 0
    }
    return f(btree)
}

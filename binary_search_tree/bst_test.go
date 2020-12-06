package binary_search_tree

import (
	"fmt"
	"testing"
)

func TestBST(t *testing.T) {
	node := &Node{Value: 20}
	node.Insert(15)
	node.Insert(18)
	node.Insert(25)
	node.Insert(22)
	node.InOrder()
	fmt.Println()
	node.PreOrder()
	fmt.Println()
	node.PostOrder()
}

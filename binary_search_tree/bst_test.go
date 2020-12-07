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
	fmt.Println()
	fmt.Println(node.GetMin())
	fmt.Println(node.GetMax())
	fmt.Println(node.Find(21))
	fmt.Println(node.Find(25).Value)


	// 输出
	// 15 18 20 22 25
	// 20,15,18,25,22,
	// 18,15,22,25,20,
	// 15
	// 25
	// <nil>
	// 25
}

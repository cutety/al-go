package binary_search_tree

import "fmt"

type Node struct {
	Value int
	Left *Node
	Right *Node
}

func (node *Node) Insert(value int) {
	if  node.Value > value {
		if node.Left == nil {
			node.Left = &Node{Value: value}
		} else {
			node.Left.Insert(value)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{Value: value}
		} else {
			node.Right.Insert(value)
		}
	}
}

// InOrder 二叉搜索树中序遍历输出的内容是有序的
func (node *Node) InOrder() {
	if node.Left != nil {
		node.Left.InOrder()
	}
	fmt.Print(node.Value," ")
	if node.Right != nil {
		node.Right.InOrder()
	}
}


func (node *Node) PreOrder() {
	fmt.Print(node.Value, ",")
	if node.Left != nil {
		node.Left.PreOrder()
	}
	if node.Right != nil {
		node.Right.PreOrder()
	}
}

func (node *Node) PostOrder() {
	if node.Left != nil {
		node.Left.PostOrder()
	}
	if node.Right != nil {
		node.Right.PostOrder()
	}
	fmt.Print(node.Value,",")
}

package binary_search_tree

import "fmt"
// Node 二叉树
type Node struct {
	Value int
	Left *Node
	Right *Node
}
// Insert 插入
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

// PreOrder 先序遍历
func (node *Node) PreOrder() {
	fmt.Print(node.Value, ",")
	if node.Left != nil {
		node.Left.PreOrder()
	}
	if node.Right != nil {
		node.Right.PreOrder()
	}
}
// PostOrder 后序遍历
func (node *Node) PostOrder() {
	if node.Left != nil {
		node.Left.PostOrder()
	}
	if node.Right != nil {
		node.Right.PostOrder()
	}
	fmt.Print(node.Value,",")
}

// GetMin 获取最小值
func (node *Node) GetMin() int {
	cur := node
	for cur.Left != nil {
		cur = cur.Left
	}
	return cur.Value
}

// GetMax 获取最大值
func (node *Node) GetMax() int {
	if node.Right != nil {
		node.Right.GetMax()
		return node.Right.Value
	} else {
		return node.Value
	}
}

// Find 查找指定数
func (node *Node) Find(value int) *Node {
	if node.Value == value {
		return node
	} else if node.Value > value && node.Left != nil {
		return node.Left.Find(value)
	} else if node.Right != nil {
		return node.Right.Find(value)
	}
	return nil
}
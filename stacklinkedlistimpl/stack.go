package stacklinkedlistimpl

import "github.com/cutety/al-go/linkedlist"

// Stack 用链表实现栈
type Stack struct {
	Node *linkedlist.Node
}

// Push 入栈
func (stack *Stack) Push(value int) {
	newNode := &linkedlist.Node{Value: value}
	newNode.Next = stack.Node
	stack.Node = newNode
}

// Pop 出栈
func (stack *Stack) Pop() int {
	//tempNode := stack.Node
	//stack.Node = stack.Node.Next
	//return tempNode.Value

	defer func() {
		stack.Node = stack.Node.Next
	}()
	return stack.Node.Value
}

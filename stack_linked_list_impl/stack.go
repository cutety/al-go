package stack_linked_list_impl

import "github.com/cutety/al-go/linked_list"

type Stack struct {
	Node *linked_list.Node
}

func(stack *Stack) Push(value int) {
	newNode := &linked_list.Node{Value: value}
	newNode.Next = stack.Node
	stack.Node = newNode
}

func (stack *Stack) Pop() int {
	//tempNode := stack.Node
	//stack.Node = stack.Node.Next
	//return tempNode.Value

	defer func() {
		stack.Node = stack.Node.Next
	}()
	return stack.Node.Value
}
package stack_slice_implemention

import (
	"fmt"
	"testing"
)


func TestStack(t *testing.T) {
	stack := &Stack{
		top:   0,
		stack: []int{},
	}
	stack.Push(1)
	stack.Push(5)
	stack.Push(4)
	stack.Push(6)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.stack)
}

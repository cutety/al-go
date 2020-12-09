package stack_linked_list_impl

import (
	"fmt"
	"github.com/cutety/al-go/linked_list"
	"testing"
)


func TestStack(t *testing.T) {
	stack :=  &Stack{Node: &linked_list.Node{}}
	stack.Push(4)
	stack.Push(7)
	stack.Push(2)
	fmt.Println("popped value:", stack.Pop())
	fmt.Println("popped value:", stack.Pop())
	fmt.Println("popped value:", stack.Pop())
}

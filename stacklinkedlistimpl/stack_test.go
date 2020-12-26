package stacklinkedlistimpl

import (
	"fmt"
	"testing"

	"github.com/cutety/al-go/linkedlist"
)

func TestStack(t *testing.T) {
	stack := &Stack{Node: &linkedlist.Node{}}
	stack.Push(4)
	stack.Push(7)
	stack.Push(2)
	fmt.Println("popped value:", stack.Pop())
	fmt.Println("popped value:", stack.Pop())
	fmt.Println("popped value:", stack.Pop())
}

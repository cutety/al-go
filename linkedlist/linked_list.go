package linkedlist

import "fmt"

// ListNode 链表
type ListNode struct {
	Value int
	Next  *ListNode
}

func New(values ...int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	if len(values) == 1 {
		return &ListNode{
			Value: values[0],
		}
	}

	head := &ListNode{
		Value: values[0],
	}

	node := head
	for i, v := range values {
		if i > 0 {
			node.Next = &ListNode{
				Value: v,
			}

			node = node.Next
		}
	}

	return head
}

func (n *ListNode) Print() {
	for n != nil {
		fmt.Printf("%d ", n.Value)
		n = n.Next
	}

	fmt.Println()
}

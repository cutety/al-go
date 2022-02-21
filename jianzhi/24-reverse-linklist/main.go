package main

import (
	ll "github.com/cutety/al-go/linkedlist"
)

// reverseLinklistRecursion 递归的方式
func reverseLinklistRecursion(head *ll.ListNode) *ll.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := reverseLinklistRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}

func reverseLinklist(head *ll.ListNode) *ll.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ll.ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

func main() {
	head := ll.New(1, 2, 3, 4, 5)

	node := reverseLinklistRecursion(head)
	node.Print()

	head2 := ll.New(4, 5, 6, 7, 8)
	node2 := reverseLinklist(head2)

	node2.Print()
}

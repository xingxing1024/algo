package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	if head.Next == nil {
		return &ListNode{}
	}

	newHead := &ListNode{}
	cur := head.Next
	for cur != nil {
		NewNode := &ListNode{Val: cur.Val}
		NewNode.Next = newHead.Next
		newHead.Next = NewNode
		cur = cur.Next
	}
	return newHead
}

func main() {
	head := &ListNode{Val: -1}
	firstNode := &ListNode{Val: 1}
	secondNode := &ListNode{Val: 2}
	thirdNode := &ListNode{Val: 3}

	head.Next = firstNode
	firstNode.Next = secondNode
	secondNode.Next = thirdNode

	newHead := reverse(head)

	for cur := newHead.Next; cur != nil; cur = cur.Next {
		fmt.Println(cur.Val)
	}
}

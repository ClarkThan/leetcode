package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// h := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	h := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}
	n := 1
	printList(h)
	h = removeNthFromEnd(h, n)
	printList(h)
}

// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 {
		return head
	}

	slow, fast := head, head

	for n > 0 {
		fast = fast.Next
		n--
	}

	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
	fmt.Println()
}

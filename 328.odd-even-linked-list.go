package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// [1,2,3,4,5,6,7,8]
	// n1 := ListNode{Val: 2}
	// n2 := ListNode{Val: 1}
	// n3 := ListNode{Val: 3}
	// n4 := ListNode{Val: 5}
	// n5 := ListNode{Val: 6}
	// n6 := ListNode{Val: 4}
	// n7 := ListNode{Val: 7}
	// n1.Next = &n2
	// n2.Next = &n3
	// n3.Next = &n4
	// n4.Next = &n5
	// n5.Next = &n6
	// n6.Next = &n7

	n1 := ListNode{Val: 1}
	n2 := ListNode{Val: 2}
	n3 := ListNode{Val: 3}
	n4 := ListNode{Val: 4}
	n5 := ListNode{Val: 5}
	n6 := ListNode{Val: 6}
	n7 := ListNode{Val: 7}
	n8 := ListNode{Val: 8}
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	n4.Next = &n5
	n5.Next = &n6
	n6.Next = &n7
	n7.Next = &n8

	ret := oddEvenList(&n1)
	n := ret
	for n != nil {
		fmt.Printf("%d ", n.Val)
		n = n.Next
	}
	fmt.Println()
}

// https://leetcode.cn/problems/odd-even-linked-list/
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	oddCur := head.Next
	oddHead := oddCur
	for cur != nil && oddCur != nil {
		if oddCur.Next == nil {
			break
		}
		cur.Next = oddCur.Next
		cur = cur.Next
		oddCur.Next = oddCur.Next.Next
		oddCur = oddCur.Next
	}

	cur.Next = oddHead
	return head
}

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
	// var l1 *ListNode = nil
	// var l2 *ListNode = &ListNode{Val: 0, Next: nil}

	showList(mergeTwoListsRecursive(l1, l2))
}

// https://leetcode.cn/problems/merge-two-sorted-lists/?favorite=2cktkvj
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// 确定返回指向和对方
	cur, rev := list1, list2
	if list1.Val >= list2.Val {
		cur = list2
		rev = list1
	}
	ret := cur

	for cur.Next != nil {
		if cur.Next.Val < rev.Val {
			cur = cur.Next
		} else {
			// 对方比原先 cur 的下一个小
			tmp := cur.Next
			cur.Next = rev
			rev = tmp
			cur = cur.Next
		}
	}

	if rev != nil {
		cur.Next = rev
	}

	return ret
}

func mergeTwoListsRecursive(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoListsRecursive(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoListsRecursive(list1, list2.Next)
		return list2
	}
}

func showList(l *ListNode) {
	for l != nil {
		fmt.Printf("%d -> ", l.Val)
		l = l.Next
	}
	fmt.Println("nil")
}

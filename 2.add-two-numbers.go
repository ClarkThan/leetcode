package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// l1 := &ListNode{Val: 0, Next: nil}
	// l2 := &ListNode{Val: 0, Next: nil}
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 6, Next: nil}}}
	// l1 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 1, Next: nil}}}
	// l2 := &ListNode{Val: 1, Next: nil}
	l3 := addTwoNumbers(l1, l2)
	printList(l3)
}

// https://leetcode-cn.com/problems/add-two-numbers
// update in place
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := l1            // 最终以 l1 的头结点作为返回
	pre := new(ListNode) // 跟踪前一个节点
	var extra int
	for l1 != nil && l2 != nil { // l1 指向当前处理节点
		curVal := l1.Val + l2.Val + extra
		if curVal >= 10 {
			l1.Val = (curVal % 10)
			extra = 1
		} else {
			l1.Val = curVal
			extra = 0
		}

		pre = l1
		l1 = l1.Next
		l2 = l2.Next
	}

	// l2 遍历完了, 那么直接看 l1 剩余就行了
	if l2 == nil {
		for l1 != nil {
			curVal := l1.Val + extra
			if curVal < 10 { // 这里不超过 10, 后面也不可能再超过 10 了
				l1.Val = curVal
				return ret
			} else {
				l1.Val = (curVal % 10)
				extra = 1
			}
			pre = l1
			l1 = l1.Next
		}
	}

	if l1 == nil {
		pre.Next = l2
		for l2 != nil {
			curVal := l2.Val + extra
			if curVal < 10 {
				l2.Val = curVal
				return ret
			} else {
				l2.Val = (curVal % 10)
				extra = 1
			}
			pre = l2
			l2 = l2.Next
		}
	}

	// 最后还有进位, 这需要在尾巴上再添加一个节点
	if extra == 1 {
		pre.Next = new(ListNode)
		pre.Next.Val = 1
	}

	return ret
}

func addTwoNumbersCopy(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := new(ListNode)
	cur := ret
	var extra int
	for l1 != nil && l2 != nil {
		curVal := l1.Val + l2.Val + extra
		if curVal >= 10 {
			cur.Val = (curVal % 10)
			extra = 1
		} else {
			cur.Val = curVal
			extra = 0
		}

		l1 = l1.Next
		l2 = l2.Next
		if l1 != nil || l2 != nil {
			cur.Next = new(ListNode)
			cur = cur.Next
		}
	}
	printList(ret)
	fmt.Println("--------")
	var other *ListNode
	if l1 != nil {
		other = l1
	} else if l2 != nil {
		other = l2
	}
	printList(other)
	fmt.Println("---------")

	for other != nil {
		curVal := other.Val + extra
		if curVal >= 10 {
			cur.Val = (curVal % 10)
			extra = 1
		} else {
			cur.Val = curVal
			extra = 0
		}

		other = other.Next

		if other != nil {
			cur.Next = new(ListNode)
			cur = cur.Next
		}
	}

	if extra > 0 {
		if cur.Next == nil {
			cur.Next = new(ListNode)
			cur.Next.Val = extra
		} else {
			cur.Next.Val = extra
		}
	}

	return ret
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val, " -> ")
		l = l.Next
	}
	fmt.Println("nil")
}

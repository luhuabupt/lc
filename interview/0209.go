package main

import (
	"fmt"
)

func main() {
	fmt.Println(" ")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeLists(arr []*ListNode) *ListNode {
	if len(arr) == 1 {
		return arr[0]
	}
	if len(arr) == 2 {
		return mergeTwoLists(arr[0], arr[1])
	}

	l := mergeLists(arr[:len(arr)/2]) // 0 ~ n/2
	r := mergeLists(arr[len(arr)/2:]) // n/2 ~ n

	return mergeTwoLists(l, r)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p := &ListNode{}
	nl := &ListNode{0, p} // 新链表的哨兵节点

	p1, p2 := l1, l2
	for {
		// for
		// 小顶堆 // log(k)
		// k/2  // k/4 // k/8

		// 1-2 3-4 -> 1 3 |  2-3 -> 2

		// merge(arr []*ListNode)
		// 1. 0-n/2 n/2-n
		// ...
		// 2. 退出条件： len(arr) == 2

		if p1 == nil {
			p.Next = p2
			break
		}
		if p2 == nil {
			p.Next = p1
			break
		}
		if p1.Val < p2.Val {
			p.Val = p1.Val
			p1 = p1.Next
		} else {
			p.Val = p2.Val
			p2 = p2.Next
		}

		p.Next = &ListNode{}
		p = p.Next
	}

	return nl.Next
}

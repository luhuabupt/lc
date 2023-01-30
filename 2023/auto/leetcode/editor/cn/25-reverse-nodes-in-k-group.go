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

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}

	var dummy *ListNode
	var pre *ListNode
	var np *ListNode
	var th *ListNode

	for n > k {
		n -= k

		for i := 0; i < k; i++ {
			np = p.Next
			p.Next = pre
			pre = p
			p = np
		}

		pe.Next = p

	}

	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

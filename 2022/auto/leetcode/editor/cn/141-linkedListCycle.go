package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.

 */

func hasCycle(p *ListNode) bool {
	if p == nil || p.Next == nil || p.Next.Next == nil {
		return false
	}
	l, f := p.Next, p.Next.Next
	for {
		if f == nil || f.Next == nil || f.Next.Next == nil || l == nil || l.Next == nil {
			return false
		}
		if f == l {
			return true
		}
		f = f.Next.Next
		l = l.Next
	}
}

//leetcode submit region end(Prohibit modification and deletion)

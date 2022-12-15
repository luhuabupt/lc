package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(h *ListNode, x int) *ListNode {
	end := &ListNode{202, nil}
	eh := end
	ans := &ListNode{201, h}

	for p, pre := h, ans; p != nil; p = p.Next {
		if p.Val >= x {
			pre.Next = p.Next

			eh.Next = p
			p.Next = nil
			eh = p

			p = pre
		} else {
			pre = p
		}
		if p.Next == nil {
			p.Next = end.Next
			break
		}
	}

	return ans.Next
}

//leetcode submit region end(Prohibit modification and deletion)

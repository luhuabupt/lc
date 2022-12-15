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
func rotateRight(head *ListNode, k int) *ListNode {
	n := 0
	end := &ListNode{}
	for p := head; p != nil; p = p.Next {
		n++
		end = p
	}
	if n == 0 {
		return head
	}
	k %= n
	if k == 0 {
		return head
	}

	ans := &ListNode{}
	for i, p := 0, head; p != nil; p = p.Next {
		if i == n-k-1 {
			end.Next = head
			ans = p.Next
			p.Next = nil
			break
		}
		i++
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

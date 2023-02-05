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
func mergeInBetween(c *ListNode, a int, b int, d *ListNode) *ListNode {
	var ap *ListNode
	var bp *ListNode
	var de *ListNode
	for p, i := c, 0; ; i++ {
		if i == a-1 {
			ap = p
		}
		if i == b+1 {
			bp = p
			break
		}
		p = p.Next
	}

	ap.Next = d
	for p := d; p != nil; p = p.Next {
		de = p
	}
	de.Next = bp

	return c
}

//leetcode submit region end(Prohibit modification and deletion)

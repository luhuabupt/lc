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
func deleteDuplicates(head *ListNode) *ListNode {
	a := []int{}
	for p := head; p != nil; p = p.Next {
		a = append(a, p.Val)
	}

	b := []int{}
	for i, v := range a {
		if (i == 0 || a[i-1] != v) && (i == len(a)-1 || a[i+1] != v) {
			b = append(b, v)
		}
	}
	if len(b) == 0 {
		return nil
	}

	nh := &ListNode{}
	p := nh
	for i, v := range b {
		p.Val = v
		if len(b) != 0 && i != len(b)-1 {
			p.Next = &ListNode{}
			p = p.Next
		}
	}

	return nh
}

//leetcode submit region end(Prohibit modification and deletion)

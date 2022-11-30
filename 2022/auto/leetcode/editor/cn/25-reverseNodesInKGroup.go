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
func reverseKGroup(head *ListNode, k int) *ListNode {
	a := []int{}
	for p := head; p != nil; p = p.Next {
		a = append(a, p.Val)
	}
	for i := 0; i < len(a); i += k {
		if len(a)-i < k {
			k = len(a) - i
		}
		for j := 0; j < k/2; j++ {
			a[i+j], a[i+k-1-j] = a[i+k-1-j], a[i+j]
		}
	}
	//fmt.Println(a)
	nh := &ListNode{}
	p := nh
	for i, v := range a {
		p.Val = v
		if i != len(a)-1 {
			p.Next = &ListNode{}
			p = p.Next
		}
	}
	return nh
}

//leetcode submit region end(Prohibit modification and deletion)

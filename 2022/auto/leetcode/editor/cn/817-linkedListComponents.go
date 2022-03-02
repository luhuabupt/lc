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
func numComponents(head *ListNode, nums []int) int {
	m := map[int]bool{}
	for _, v := range nums {
		m[v] = true
	}

	ans, tmp := 0, 0
	for p := head; p != nil; p = p.Next {
		if m[p.Val] {
			tmp++
		} else {
			if tmp > 0 {
				ans++
			}
			tmp = 0
		}
	}
	if tmp > 0 {
		ans++
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubPath(head *ListNode, root *TreeNode) bool {
	ans := false

	var chk func(p *TreeNode, q *ListNode) bool
	chk = func(p *TreeNode, q *ListNode) bool {
		if q == nil {
			return true
		}
		if p == nil || p.Val != q.Val {
			return false
		}
		return chk(p.Left, q.Next) || chk(p.Right, q.Next)
	}

	var dfs func(p *TreeNode)
	dfs = func(p *TreeNode) {
		if p == nil {
			return
		}
		if chk(p, head) {
			ans = true
			return
		}
		dfs(p.Left)
		dfs(p.Right)
	}

	dfs(root)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

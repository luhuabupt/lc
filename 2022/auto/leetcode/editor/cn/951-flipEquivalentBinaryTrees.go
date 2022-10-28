package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flipEquiv(r1 *TreeNode, r2 *TreeNode) bool {
	var same func(p, q *TreeNode) bool
	same = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		return same(p.Left, q.Left) && same(p.Right, q.Right) || same(p.Left, q.Right) && same(p.Right, q.Left)
	}

	return same(r1, r2)
}

//leetcode submit region end(Prohibit modification and deletion)

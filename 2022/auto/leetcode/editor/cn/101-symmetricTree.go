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
func isSymmetric(root *TreeNode) bool {
	var dfs func(a, b *TreeNode) bool
	dfs = func(a, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		}
		if a == nil || b == nil {
			return false
		}
		if a.Val != b.Val {
			return false
		}
		if !dfs(a.Left, b.Right) {
			return false
		}
		if !dfs(a.Right, b.Left) {
			return false
		}
		return true
	}

	return dfs(root.Left, root.Right)
}

//leetcode submit region end(Prohibit modification and deletion)

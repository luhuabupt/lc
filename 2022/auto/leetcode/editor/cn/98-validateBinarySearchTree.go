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
func isValidBST(root *TreeNode) bool {
	var dfs func(p *TreeNode, i, a int) bool
	dfs = func(p *TreeNode, i, a int) bool {
		if p == nil {
			return true
		}
		if p.Val < i || p.Val > a {
			return false
		}
		return dfs(p.Left, i, p.Val-1) && dfs(p.Right, p.Val+1, a)
	}

	return dfs(root, -1<<31, 1<<31-1)
}

//leetcode submit region end(Prohibit modification and deletion)

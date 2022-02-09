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
func inorderTraversal(root *TreeNode) (ans []int) {
	var dfs func(p *TreeNode)
	dfs = func(p *TreeNode) {
		if p == nil {
			return
		}
		dfs(p.Left)
		ans = append(ans, p.Val)
		dfs(p.Right)
	}

	dfs(root)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

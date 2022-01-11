package main

import "fmt"

// longest-univalue-path 最长同值路径  2022-01-07 23:42:29
func main() {
	fmt.Println(longestUnivaluePath(nil))
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
func longestUnivaluePath(root *TreeNode) int {
	max := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	ans := 0
	var dfs func(p *TreeNode, v int) int
	dfs = func(p *TreeNode, v int) int {
		if p == nil {
			return 0
		}
		l := dfs(p.Left, p.Val)
		r := dfs(p.Right, p.Val)

		ans = max(ans, l+r)

		if p.Val != v {
			return 0
		}
		return max(l, r) + 1
	}
	dfs(root, 0)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

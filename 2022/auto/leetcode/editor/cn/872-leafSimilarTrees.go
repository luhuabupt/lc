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
func leafSimilar(a *TreeNode, b *TreeNode) bool {
	ar := [2][]int{}
	var dfs func(p *TreeNode, i int)
	dfs = func(p *TreeNode, i int) {
		if p == nil {
			return
		}
		if p.Left == nil && p.Right == nil {
			ar[i] = append(ar[i], p.Val)
		}
		dfs(p.Left, i)
		dfs(p.Right, i)
	}

	dfs(a, 0)
	dfs(b, 1)

	if len(ar[0]) != len(ar[1]) {
		return false
	}
	for i, v := range ar[0] {
		if v != ar[1][i] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

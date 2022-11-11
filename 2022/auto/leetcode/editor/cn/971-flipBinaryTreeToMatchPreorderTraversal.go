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
func flipMatchVoyage(root *TreeNode, a []int) []int {
	ans := []int{}
	flag := false
	ti := 0
	var dfs func(p *TreeNode)
	dfs = func(p *TreeNode) {
		if p.Val != a[ti] {
			flag = true
			return
		}
		ti++
		if p.Left != nil && p.Right != nil {
			if p.Left.Val != a[ti] {
				if p.Right.Val == a[ti] {
					ans = append(ans, p.Val)
					p.Left, p.Right = p.Right, p.Left
				} else {
					flag = true
					return
				}
			}
		}
		if p.Left != nil {
			dfs(p.Left)
		}
		if p.Right != nil {
			dfs(p.Right)
		}
	}
	dfs(root)
	if flag {
		return []int{-1}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

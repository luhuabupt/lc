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
func treeQueries(root *TreeNode, queries []int) []int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := int(1e5) + 2
	son := make([]int, n)
	level := make([][2][2]int, n) // max, val
	for i := 0; i < n; i++ {
		level[i] = [2][2]int{{-1, -1}, {-1, -1}}
	}
	dl := make([]int, n) // level, val
	mxd := 0

	var dfs func(p *TreeNode, d int) int
	dfs = func(p *TreeNode, d int) int {
		if p == nil {
			return 0
		}
		mx := max(dfs(p.Left, d+1), dfs(p.Right, d+1))
		son[p.Val] = mx

		if mx >= level[d][0][0] {
			level[d] = [2][2]int{{mx, p.Val}, level[d][0]}
		} else if mx > level[d][1][0] {
			level[d][1] = [2]int{mx, p.Val}
		}
		mxd = max(mxd, d)
		dl[p.Val] = d

		return mx + 1
	}
	dfs(root, 0)

	ans := []int{}
	for _, v := range queries {
		lv := level[dl[v]]
		x := lv[0][0]
		if lv[0][1] == v {
			x = lv[1][0]
		}
		ans = append(ans, dl[v]+x)
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

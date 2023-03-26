package main

import "fmt"

func main() {
	fmt.Println(" ")
}

//leetcode submit region begin(Prohibit modification and deletion)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	ans := false
	var dfs func(p *TreeNode) int
	dfs = func(p *TreeNode) int {
		if p == nil {
			return 0
		}
		l := dfs(p.Left)
		r := dfs(p.Right)
		if p.Val == x {
			if l > n/2 || r > n/2 {
				ans = true
			}
			if l+r+1 < n/2 {
				ans = true
			}
		}
		return l + r + 1
	}

	dfs(root)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

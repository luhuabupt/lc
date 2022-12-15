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
func distributeCoins(root *TreeNode) int {
	abs := func(v int) int {
		if v > 0 {
			return v
		}
		return -v
	}

	ans := 0
	var cal func(p *TreeNode) int
	cal = func(p *TreeNode) int {
		if p == nil {
			return 0
		}
		l := cal(p.Left)
		r := cal(p.Right)
		ans += abs(l) + abs(r)

		return l + r + p.Val - 1
	}

	cal(root)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

package main

import "fmt"

func main() {
	fmt.Println(tallestBillboard([]int{1, 2, 3, 6}))
	fmt.Println(tallestBillboard([]int{1, 2, 3, 4, 5, 6}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func tallestBillboard(a []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(a)
	dp := make([]map[int]int, n)
	for i, v := range a {
		dp[i] = map[int]int{}
		if i == 0 {
			dp[i] = map[int]int{v: v, 0: 0, -v: v}
			continue
		}
		for j, x := range dp[i-1] {
			dp[i][j] = max(dp[i][j], x)
			dp[i][j-v] = max(dp[i][j-v], x+v)
			dp[i][j+v] = max(dp[i][j+v], x+v)
		}
		//fmt.Println(i, dp[i])
	}
	return dp[n-1][0] / 2
}

//leetcode submit region end(Prohibit modification and deletion)

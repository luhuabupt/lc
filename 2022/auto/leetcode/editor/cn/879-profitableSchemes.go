package main

import "fmt"

func main() {
	fmt.Println(profitableSchemes(5, 3, []int{4, 2}, []int{0, 3}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	M := int(1e9) + 7
	dp := make([][]int, n+1) // 人数, 盈利
	for i, _ := range dp {
		dp[i] = make([]int, minProfit+1)
	}
	dp[0][0] = 1

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i, v := range group {
		for x := n - v; x >= 0; x-- {
			for j, jv := range dp[x] {
				dp[x+v][min(minProfit, j+profit[i])] += jv
				dp[x+v][min(minProfit, j+profit[i])] %= M
			}
		}
	}

	ans := 0
	for _, a := range dp {
		ans += a[minProfit]
		ans %= M
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

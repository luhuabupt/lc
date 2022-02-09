package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func cherryPickup(grid [][]int) int {
	n := len(grid)
	dp := map[int]map[int]int{}
	for i := 0; i < n; i++ {
		dp[i] = map[int]int{}
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i, arr := range grid {
		for j, v := range arr {
			if v == -1 {
				continue
			}
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + v
		}
	}

	return dp[n-1][n-1]
}

//leetcode submit region end(Prohibit modification and deletion)

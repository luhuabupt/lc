package main

func minCostII(c [][]int) int {
	n, k := len(c), len(c[0])

	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, k)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			v := 1<<31 - 1
			for x := 0; x < k; x++ {
				if x == j {
					continue
				}
				cur := c[i][j]
				if i > 0 {
					cur += dp[i-1][x]
				}
				if cur < v {
					v = cur
				}
			}
			dp[i][j] = v
		}
	}

	ans := 1<<31 - 1
	for _, v := range dp[n-1] {
		if v < ans {
			ans = v
		}
	}

	return ans
}

package main

func longestSubsequence(arr []int, difference int) int {
	dp := map[int]int{}
	ans := 1
	for _, v := range arr {
		if dp[v-difference] == 0 {
			dp[v] = 1
		} else if dp[v-difference] >= dp[v] {
			dp[v] = dp[v-difference] + 1
			if dp[v] > ans {
				ans = dp[v]
			}
		}
	}

	return ans
}

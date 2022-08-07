package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func numSquares(n int) int {
	s, j, dp := []int{}, 1, make([]int, n+1)
	for i := 1; i <= n; i++ {
		if j*j == i {
			s = append(s, j*j)
			j++
		}
		for _, v := range s {
			if dp[i] == 0 || dp[i-v]+1 < dp[i] {
				dp[i] = dp[i-v] + 1
			}
		}
	}
	return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)

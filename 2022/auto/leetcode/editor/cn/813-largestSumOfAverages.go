package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func largestSumOfAverages(a []int, k int) float64 {
	max := func(x, y float64) float64 {
		if x > y {
			return x
		}
		return y
	}

	n := len(a)
	dp := make([][]float64, n)
	s := make([]int, n)
	t := 0
	for i := 0; i < n; i++ {
		t += a[i]
		s[i] = t
		dp[i] = make([]float64, k+1)
		dp[i][1] = float64(s[i]) / float64(i+1)
		for j := 2; j <= k; j++ {
			for x := j - 2; x < i; x++ {
				dp[i][j] = max(dp[i][j], dp[x][j-1]+float64(s[i]-s[x])/float64(i-x))
			}
		}
	}

	return dp[n-1][k]
}

//leetcode submit region end(Prohibit modification and deletion)

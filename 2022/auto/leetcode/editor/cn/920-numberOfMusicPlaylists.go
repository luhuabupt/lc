package main

import "fmt"

func main() {
	fmt.Println(numMusicPlaylists(3, 3, 1))
	fmt.Println(numMusicPlaylists(2, 3, 0))
	fmt.Println(numMusicPlaylists(2, 3, 1))
}

//leetcode submit region begin(Prohibit modification and deletion)
func numMusicPlaylists(n int, l int, k int) int {
	dp := make([][]int, l)
	M := int(1e9 + 7)

	for i := 0; i < l; i++ {
		dp[i] = make([]int, n+1)
		if i == 0 {
			dp[0][1] = n
		} else if i <= k {
			dp[i][i+1] = (n - i) * dp[i-1][i]
			dp[i][i+1] %= M
		} else {
			for j := k; j <= n; j++ {
				if j-k > 0 {
					dp[i][j] += dp[i-1][j] * (j - k) // same
				}
				if n-j+1 > 0 && j-1 >= 0 {
					dp[i][j] += dp[i-1][j-1] * (n - j + 1) // not
				}
				dp[i][j] %= M
			}
		}
	}

	return dp[l-1][n]
}

//leetcode submit region end(Prohibit modification and deletion)

package main

import "fmt"

func main() {
	fmt.Println(numTilings(30))
	fmt.Println(numTilings(4))
}

//leetcode submit region begin(Prohibit modification and deletion)
func numTilings(n int) int {
	M := int(1e9 + 7)
	dp := make([]int, 1001)
	dp[0] = 1
	dp[1] = 2
	dp[2] = 4
	dp[3] = 9
	for i := 4; i <= n; i++ {
		dp[i] = 2*dp[i-1] + dp[i-3]
		dp[i] %= M
	}
	return (dp[n] + M - dp[n-1]) % M
}

//leetcode submit region end(Prohibit modification and deletion)

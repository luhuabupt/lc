package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getTriggerTime([][]int{{2, 8, 4}, {2, 5, 0}, {10, 9, 8}}, [][]int{{2, 11, 3}, {15, 10, 7}, {9, 17, 12}, {8, 1, 14}}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func getTriggerTime(inc [][]int, req [][]int) []int {
	n := len(inc)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := make([][3]int, n+1)
	dp[0] = [3]int{0, 0, 0}
	for i := 1; i <= n; i++ {
		dp[i][0] = dp[i-1][0] + inc[i-1][0]
		dp[i][1] = dp[i-1][1] + inc[i-1][1]
		dp[i][2] = dp[i-1][2] + inc[i-1][2]
	}

	ans := []int{}
	for _, a := range req {
		x := 0
		for i, v := range a {
			x = max(x, sort.Search(n+1, func(j int) bool {
				return dp[j][i] >= v
			}))
		}
		if x == n+1 {
			ans = append(ans, -1)
		} else {
			ans = append(ans, x)
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

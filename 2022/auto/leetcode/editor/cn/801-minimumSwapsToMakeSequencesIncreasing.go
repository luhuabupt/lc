package main

import "fmt"

func main() {
	fmt.Println(minSwap([]int{3, 3, 8, 9, 10}, []int{1, 7, 4, 6, 8}))
	fmt.Println(minSwap([]int{0, 4, 4, 5, 9}, []int{0, 1, 6, 8, 10}))
	fmt.Println(minSwap([]int{0, 3, 4, 9, 10}, []int{2, 3, 7, 5, 6}))
	fmt.Println(minSwap([]int{0, 7, 8, 10, 10, 11, 12, 13, 19, 18}, []int{4, 4, 5, 7, 11, 14, 15, 16, 17, 20}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func minSwap(a []int, b []int) int {
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}

	n := len(a)
	dp := make([][2]int, n) // i 0-不换 1-换

	dp[0] = [2]int{0, 1}

	for i := 1; i < n; i++ {
		dp[i] = [2]int{1 << 60, 1 << 60}
		if a[i] > a[i-1] && b[i] > b[i-1] {
			dp[i][0] = min(dp[i][0], dp[i-1][0])
			dp[i][1] = min(dp[i][1], dp[i-1][1]+1)
		}
		if a[i] > b[i-1] && b[i] > a[i-1] {
			dp[i][0] = min(dp[i][0], dp[i-1][1])
			dp[i][1] = min(dp[i][1], dp[i-1][0]+1)
		}
	}

	return min(dp[n-1][0], dp[n-1][1])
}

//leetcode submit region end(Prohibit modification and deletion)

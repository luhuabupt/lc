package classic

func main() {
	zo([]int{4, 3, 8, 5, 7}, []int{3, 2, 6, 4, 12}, 11)
}

// 完全背包问题
// 物品size, 权重weight, 背包容量cap, 每个物品可选任意次, 求最大价值
func all(size, weight []int, cap int) int {
	dp := map[int]int{0: 0}
	for j := 1; j <= cap; j++ {
		dp[j] = dp[j-1]
		for i, v := range size {
			dp[j] = twoMax1(dp[j-v]+weight[i], dp[j])
		}
	}
	return dp[cap]
}

// 01背包问题
// 物品size, 权重weight, 背包容量cap, 每个物品可选1次, 求最大价值
func zo(size, weight []int, cap int) int {
	dp := map[int]int{0: 0}
	for i, v := range size {
		for j := cap; j >= v; j-- { // 倒着取, 防止选重
			dp[j] = twoMax1(dp[j-v]+weight[i], dp[j])
		}
	}
	ans := 0
	for j := 1; j <= cap; j++ {
		if dp[j] > ans {
			ans = dp[j]
		}
	}
	return ans
}

func twoMax1(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

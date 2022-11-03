package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(minDistance([]int{1, 4, 8, 10, 20}, 3))
	fmt.Println(minDistance([]int{2, 3, 5, 12, 18}, 2))
}

//leetcode submit region begin(Prohibit modification and deletion)
func minDistance(a []int, k int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := len(a)
	sort.Ints(a)
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, k+1)
			for x := 0; x <= k; x++ {
				dp[i][j][x] = -1
			}
		}
	}

	cal := func(i, j int) int {
		res := 0
		mid := 0
		if (j-i+1)&1 > 0 {
			mid = a[(j+i)/2]
		} else {
			mid = a[(j+i)/2] + a[(j+i)/2+1]
			mid /= 2
		}
		for x := i; x <= j; x++ {
			if a[x] < mid {
				res += mid - a[x]
			} else {
				res += a[x] - mid
			}
		}
		dp[i][j][1] = res
		return res
	}

	var dfs func(i, j, t int) int
	dfs = func(i, j, t int) int {
		if t >= j-i+1 {
			return 0
		}
		if dp[i][j][t] >= 0 {
			return dp[i][j][t]
		}
		if t == 1 {
			dp[i][j][1] = cal(i, j)
			return dp[i][j][1]
		}

		mi := 1 << 60
		for x := i; x < j; x++ {
			mi = min(mi, dfs(i, x, t-1)+dfs(x+1, j, 1))
		}

		dp[i][j][t] = mi
		return mi
	}

	return dfs(0, n-1, k)
}

//leetcode submit region end(Prohibit modification and deletion)

package main

import (
	"fmt"
	"sort"
)

// video-stitching 视频拼接  2022-12-17 16:29:54
func main() {
	fmt.Println("")
}

//leetcode submit region begin(Prohibit modification and deletion)
func videoStitching(clips [][]int, time int) int {
	sort.Slice(clips, func(i, j int) bool {
		return clips[i][0] < clips[j][0] || clips[i][0] == clips[j][0] && clips[i][1] < clips[j][1]
	})
	if clips[0][0] != 0 || clips[len(clips)-1][1] < time {
		return -1
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dp := [101][101]int{}
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			dp[i][j] = 1 << 60
		}
	}
	dp[0][0] = 0

	for i, v := range clips {
		if i > 0 {
			if v[0] > clips[i-1][1] {
				return -1
			}
		}
		for l := v[0]; l < v[1]; l++ {
			dp[0][v[1]] = min(dp[0][v[1]], dp[0][l]+1)
		}
		for j := v[0]; j < v[1]; j++ {
			dp[0][j] = min(dp[0][j], dp[0][v[1]])
		}
	}

	return dp[0][time]
}

//leetcode submit region end(Prohibit modification and deletion)

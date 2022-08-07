package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(deleteAndEarn([]int{1, 1, 1, 2, 4, 5, 5, 5, 6}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func deleteAndEarn(nums []int) int {
	dp := map[int]int{}
	cnt := map[int]int{}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	arr := []int{}
	for _, v := range nums {
		cnt[v]++
		if cnt[v] == 1 {
			arr = append(arr, v)
		}
	}
	sort.Ints(arr)

	for i, v := range arr {
		pre := 0
		if i > 0 && arr[i-1] < v-1 {
			pre = arr[i-1]
		} else if i > 1 {
			pre = arr[i-2]
		}
		dp[v] = max(dp[pre]+v*cnt[v], dp[v-1])
	}

	//fmt.Println(dp)

	return dp[arr[len(arr)-1]]
}

//leetcode submit region end(Prohibit modification and deletion)

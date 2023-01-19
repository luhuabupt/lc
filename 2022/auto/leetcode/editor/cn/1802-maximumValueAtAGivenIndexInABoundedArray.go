package main

import (
	"fmt"
	"sort"
)

// maximum-value-at-a-given-index-in-a-bounded-array 有界数组中指定下标处的最大值  2023-01-04 09:36:47
func main() {
	fmt.Println("")
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxValue(n int, index int, maxSum int) int {
	cal := func(v int) int {
		t := n - (v - 1)
		if v <= index {
			t += (v - 1) * (1 + v - 1) / 2
		} else {
			t += (index + 1) * (v - 1 + v - 1 - index) / 2
		}
		if index+v <= n {
			t += (v - 1) * (1 + v - 1) / 2
		} else {
			t += (n - index) * (v - 1 + v - (n - index)) / 2
		}
		return t
	}

	return sort.Search(maxSum+1, func(v int) bool {
		return cal(v) > maxSum
	}) - 1
}

//leetcode submit region end(Prohibit modification and deletion)

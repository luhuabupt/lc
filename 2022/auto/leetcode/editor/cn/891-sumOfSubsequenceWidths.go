package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sumSubseqWidths([]int{1, 2, 3}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func sumSubseqWidths(nums []int) int {
	sort.Ints(nums)
	M := int(1e9 + 7)
	n := len(nums)

	sum, x := make([]int, n), 0
	for i, v := range nums {
		x += v
		sum[i] = x
	}

	ans, p := 0, 1
	for i := 1; i < n; i++ {
		ans += (sum[n-1] - sum[n-1-i] - sum[i-1]) * p
		ans %= M
		p *= 2
		p %= M
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

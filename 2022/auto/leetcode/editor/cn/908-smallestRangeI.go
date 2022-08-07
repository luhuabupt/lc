package main

import "sort"

func main() {
}

//leetcode submit region begin(Prohibit modification and deletion)
func smallestRangeI(a []int, k int) int {
	sort.Ints(a)
	if a[len(a)-1] - a[0] <= 2 * k {
		return 0
	}

	return a[len(a)-1]- a[0] - 2 * k
}
//leetcode submit region end(Prohibit modification and deletion)

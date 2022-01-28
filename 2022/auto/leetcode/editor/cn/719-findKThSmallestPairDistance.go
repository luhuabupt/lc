package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(smallestDistancePair([]int{1, 3, 1}, 2))
}

//leetcode submit region begin(Prohibit modification and deletion)
func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)

	return sort.Search(nums[n-1]-nums[0]+1, func(v int) bool {
		sum := 0
		for i := 0; i < n; i++ {
			x := sort.SearchInts(nums, nums[i]+v+1) // 小于等于v的
			sum += x - i - 1
		}
		return sum >= k
	})
}

//leetcode submit region end(Prohibit modification and deletion)

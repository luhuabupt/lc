package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(advantageCount([]int{2, 0, 4, 1, 2}, []int{1, 3, 0, 0, 2}))
	fmt.Println(advantageCount([]int{2, 0, 4, 1, 2}, []int{0, 0, 1, 2, 3}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func advantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	sort.Ints(nums1)
	ans := make([]int, n)

	type pair struct {
		i, v int
	}
	n2 := []pair{}
	for i, v := range nums2 {
		n2 = append(n2, pair{i, v})
	}

	sort.Slice(n2, func(i, j int) bool {
		return n2[i].v < n2[j].v
	})

	not := []int{}
	match := []int{}
	for idx, i := 0, 0; idx < n; idx++ {
		for ; i < n && nums1[i] <= n2[idx].v; i++ {
			not = append(not, i)
		}
		if i == n {
			match = append(match, not...)
		} else {
			match = append(match, i)
		}
		i++
	}

	for i := 0; i < n; i++ {
		ans[n2[i].i] = nums1[match[i]]
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

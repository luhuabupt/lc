package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxWidthRamp([]int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}))
	fmt.Println(maxWidthRamp([]int{3, 4, 1, 2}))
	fmt.Println(maxWidthRamp([]int{5, 4, 3}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxWidthRamp(a []int) int {
	st := []int{}
	ans := 0
	for i, v := range a {
		if i > 0 {
			x := sort.Search(len(st), func(j int) bool {
				return v >= a[st[j]]
			})
			if x < len(st) && i-st[x] > ans {
				ans = i - st[x]
			}
			//fmt.Println(i, v, x, st)
		}
		if len(st) == 0 || v < a[st[len(st)-1]] {
			st = append(st, i)
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

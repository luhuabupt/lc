package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(shortestSubarray([]int{2, -1, 2}, 3))                                                                                //3
	fmt.Println(shortestSubarray([]int{84, -37, 32, 40, 95}, 167))                                                                   //3
	fmt.Println(shortestSubarray([]int{-28, 81, -20, 28, -295}, 89))                                                                 // 3
	fmt.Println(shortestSubarray([]int{-34, 37, 51, 3, -12, -50, 51, 100, -47, 99, 34, 14, -13, 89, 31, -14, -44, 23, -38, 6}, 151)) // 2
	fmt.Println(shortestSubarray([]int{37, 51, 3, -12, -50, 51, 100}, 151))                                                          // 2
}

//leetcode submit region begin(Prohibit modification and deletion)
func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	s := make([]int, n)
	x := 0
	ans := -1

	st := [][2]int{{-1, 0}}
	for i, v := range nums {
		x += v
		s[i] = x

		f := sort.Search(len(st), func(i int) bool {
			return st[i][1] > x-k
		})
		if f > 0 && (ans == -1 || i-st[f-1][0] < ans) {
			ans = i - st[f-1][0]
		}

		for len(st) > 0 && st[len(st)-1][1] > x {
			st = st[:len(st)-1]
		}
		st = append(st, [2]int{i, x})
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

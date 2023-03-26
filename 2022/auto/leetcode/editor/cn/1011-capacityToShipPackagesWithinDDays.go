package main

import (
	"fmt"
	"sort"
)

// capacity-to-ship-packages-within-d-days 在 D 天内送达包裹的能力  2022-12-16 17:16:00
func main() {
	fmt.Println("")
}

//leetcode submit region begin(Prohibit modification and deletion)
func shipWithinDays(w []int, d int) int {
	return sort.Search(1e8, func(c int) bool {
		tc := 0
		t := 0
		for _, v := range w {
			if v > c {
				return true
			}
			if tc+v > c {
				t++
				tc = v
			}
		}
		if tc != 0 {
			t++
		}
		return t > d
	})
}

//leetcode submit region end(Prohibit modification and deletion)

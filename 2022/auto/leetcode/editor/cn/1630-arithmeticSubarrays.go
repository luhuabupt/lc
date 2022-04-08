package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(checkArithmeticSubarrays([]int{4, 6, 5, 9, 3, 7}, []int{0, 0, 2}, []int{2, 3, 5}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func checkArithmeticSubarrays(a []int, l []int, r []int) []bool {
	n := len(l)
	ans := make([]bool, n)

	chk := func(v []int) bool {
		sort.Ints(v)
		for i := 1; i < len(v); i++ {
			if v[i]-v[i-1] != v[1]-v[0] {
				return false
			}
		}

		return true
	}

	for i, v := range l {
		x := append([]int{}, a[v:r[i]+1]...)
		ans[i] = chk(x)
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

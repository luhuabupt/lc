package main

import (
	"fmt"
	"sort"
)

// kth-smallest-number-in-multiplication-table 乘法表中第k小的数  2022-01-07 22:08:40
func main() {
	fmt.Println(findKthNumber(3, 3, 5))                //3
	fmt.Println(findKthNumber(2, 3, 6))                //6
	fmt.Println(findKthNumber(10, 10, 50))             //24
	fmt.Println(findKthNumber(10, 3, 19))              //12
	fmt.Println(findKthNumber(30000, 30000, 12000001)) //1647184
}

//leetcode submit region begin(Prohibit modification and deletion)
func findKthNumber(m int, n int, k int) int {
	return sort.Search(m*n, func(v int) bool {
		cnt := 0
		for i := 1; i <= m; i++ {
			cnt += sort.Search(n, func(vv int) bool {
				return (vv+1)*i > v+1
			})
		}
		return cnt >= k
	}) + 1
}

//leetcode submit region end(Prohibit modification and deletion)

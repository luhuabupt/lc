package main

import "sort"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func largestPerimeter(a []int) int {
	sort.Ints(a)
	ans := 0
	for i := len(a) - 1; i >= 2; i-- {
		if a[i] < a[i-1]+a[i-2] {
			ans = a[i-1] + a[i-2] + a[i]
			break
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

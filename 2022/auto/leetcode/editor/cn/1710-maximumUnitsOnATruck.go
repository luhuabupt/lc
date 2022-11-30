package main

import "sort"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func maximumUnits(a [][]int, t int) int {
	sort.Slice(a, func(i, j int) bool {
		return a[i][1] > a[j][1]
	})
	ans := 0
	for _, v := range a {
		if v[0] > t {
			ans += t * v[1]
			t = 0
			break
		} else {
			t -= v[0]
			ans += v[0] * v[1]
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

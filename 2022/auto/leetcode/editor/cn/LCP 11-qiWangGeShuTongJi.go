package main

import "sort"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func expectNumber(s []int) int {
	sort.Ints(s)
	n := len(s)

	a := []int{}
	t := 1
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			t++
		} else {
			a = append(a, t)
			t = 1
		}
	}
	a = append(a, t)

	return len(a)
}

//leetcode submit region end(Prohibit modification and deletion)

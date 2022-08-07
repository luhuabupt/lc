package main

import "fmt"

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))

}

//leetcode submit region begin(Prohibit modification and deletion)
func singleNonDuplicate(a []int) int {
	n := len(a)
	l, r, m := 0, n-1, 0
	for l <= r {
		m = (l + r) / 2
		if (m == 0 || a[m] != a[m-1]) && (m == n-1 || a[m] != a[m+1]) {
			break
		}

		x := m - 1
		if m > 0 && a[m-1] == a[m] {
			x = m - 2
		}

		if x&1 == 0 {
			r = x
		} else {
			l = x + 3
		}
	}

	return a[m]
}

//leetcode submit region end(Prohibit modification and deletion)

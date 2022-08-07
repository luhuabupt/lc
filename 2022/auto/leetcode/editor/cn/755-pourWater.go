package main

import "fmt"

func main() {
	fmt.Println(pourWater([]int{2, 1, 1, 2, 1, 2, 2}, 4, 3))
	fmt.Println(pourWater([]int{1, 2, 3, 4}, 2, 2))
	fmt.Println(pourWater([]int{1, 2, 3, 4, 3, 2, 1, 2, 3, 4, 3, 2, 1}, 10, 2))
	fmt.Println(pourWater([]int{14, 10, 10, 3, 13, 1, 2, 1, 2, 5}, 1, 0))
}

//leetcode submit region begin(Prohibit modification and deletion)
func pourWater(h []int, v int, k int) []int {
	n := len(h)

	var less func(i, x, d int) bool
	less = func(i, x, d int) bool {
		if i == 0 || i == n-1 {
			if h[i] < x {
				h[i]++
				return true
			} else {
				return false
			}
		}
		if h[i] > x {
			return false
		}
		if less(i+d, h[i], d) {
			return true
		}
		if h[i] < x {
			h[i]++
			return true
		}
		return false
	}

	for {
		for v > 0 && k > 0 && less(k-1, h[k], -1) {
			v--
		}
		for v > 0 && k+1 < n && less(k+1, h[k], 1) {
			v--
		}
		if v == 0 {
			break
		}
		h[k]++
		v--
	}

	return h
}

//leetcode submit region end(Prohibit modification and deletion)

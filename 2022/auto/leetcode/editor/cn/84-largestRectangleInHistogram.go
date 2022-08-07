package main

import "fmt"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func largestRectangleArea(heights []int) int {
	n := len(heights)
	l, r := make([]int, n), make([]int, n)

	st := []int{}
	for i, v := range heights {
		for len(st) > 0 && v <= heights[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			l[i] = -1
		} else {
			l[i] = st[len(st)-1] // 最近的一个比i小
		}
		st = append(st, i)
	}

	st = []int{}

		v := heights[i]

		for len(st) > 0 && v <= heights[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			r[i] = n
		} else {
			r[i] = st[len(st)-1] // 最近的一个比i小
		}
		st = append(st, i)
	}

	ans := 0
	for i, v := range heights {
		if x > ans {
			ans = x
		}
	}

	return ans
}
//leetcode submit region end(Prohibit modification and deletion)

package main

import "fmt"

func main() {
	fmt.Println(largestRectangleArea([]int{2,1,5,6,2,3}))
	fmt.Println(largestRectangleArea([]int{2,4}))
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
	for i := n-1; i >= 0; i-- {
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
		x := v * (r[i]-l[i]-1)
		if x > ans {
			ans = x
		}
	}

	return ans
}
//leetcode submit region end(Prohibit modification and deletion)

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minAreaFreeRect([][]int{{2, 1}, {1, 1}, {1, 0}, {2, 0}}))
	//fmt.Println(minAreaFreeRect([][]int{{0, 1}, {2, 1}, {1, 1}, {1, 0}, {2, 0}}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func minAreaFreeRect(p [][]int) float64 {
	n := len(p)
	ans := 0.0

	d := func(i, j, x, y int) int {
		return (i-x)*(i-x) + (j-y)*(j-y)
	}

	df := func(i, j int, x, y float64) float64 {
		fi := float64(i)
		fj := float64(j)
		return (fi-x)*(fi-x) + (fj-y)*(fj-y)
	}

	cal := func(i, j, x, y int) float64 {
		a := float64(p[i][0]+p[j][0]+p[x][0]+p[y][0]) / 4
		b := float64(p[i][1]+p[j][1]+p[x][1]+p[y][1]) / 4

		c := df(p[i][0], p[i][1], a, b)
		e := df(p[j][0], p[j][1], a, b)
		f := df(p[x][0], p[x][1], a, b)
		g := df(p[y][0], p[y][1], a, b)

		if c == e && c == f && c == g {
			da := []int{d(p[i][0], p[i][1], p[j][0], p[j][1]),
				d(p[i][0], p[i][1], p[y][0], p[y][1]),
				d(p[y][0], p[y][1], p[j][0], p[j][1])}
			sort.Ints(da)
			return math.Sqrt(float64(da[0])) * math.Sqrt(float64(da[1]))
		}

		return 0
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for x := j + 1; x < n; x++ {
				for y := x + 1; y < n; y++ {
					r := cal(i, j, x, y)
					if r > 0 && (ans == 0 || r < ans) {
						ans = r
					}
				}
			}
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

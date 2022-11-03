package main

import "math"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func bestCoordinate(t [][]int, r int) []int {
	df := func(i, j, x, y int) int {
		return (i-x)*(i-x) + (j-y)*(j-y)
	}

	m := map[int]map[int]int{}
	for _, v := range t {
		i, j := v[0], v[1]
		for x := 0; x <= 2*r; x++ {
			for y := 0; y <= 2*r; y++ {
				d := df(i, j, i-r+x, j-r+y)
				if d <= r*r {
					if m[i-r+x] == nil {
						m[i-r+x] = map[int]int{}
					}
					m[i-r+x][j-r+y] += int(float64(v[2]) / (1 + math.Sqrt(float64(d))))
				}
			}
		}
	}

	mx := 0
	ans := []int{0, 0}
	for i, ar := range m {
		for j, v := range ar {
			if len(ans) == 0 || v > mx || v == mx && i >= 0 && j >= 0 && (i < ans[0] || i == ans[0] && j < ans[1]) {
				mx = v
				ans = []int{i, j}
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

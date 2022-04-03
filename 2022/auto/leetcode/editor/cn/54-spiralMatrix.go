package main

import "fmt"

// spiral-matrix 螺旋矩阵  2022-03-28 20:57:19
func main() {
	fmt.Println(spiralMatrix())
}

//leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(a [][]int) []int {
	ans := []int{}
	m, n := len(a), len(a[0])
	dr := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	i, j, do, idx := 0, 0, 0, 0

	for {
		ans = append(ans, a[i][j])
		a[i][j] = 1000000
		do++
		if do == m*n {
			break
		}

		i += dr[idx][0]
		j += dr[idx][1]

		if i < 0 || j < 0 || i >= m || j >= n || a[i][j] > 10000 {
			i -= dr[idx][0]
			j -= dr[idx][1]

			idx++
			idx %= 4

			i += dr[idx][0]
			j += dr[idx][1]
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

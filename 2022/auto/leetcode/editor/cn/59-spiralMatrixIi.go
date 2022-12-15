package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(3))
}

//leetcode submit region begin(Prohibit modification and deletion)
func generateMatrix(n int) [][]int {
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
	}

	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for i, j, d, dx := 0, 0, 1, 0; d <= n*n; d++ {
		g[i][j] = d

		x, y := i+dir[dx][0], j+dir[dx][1]
		if x < 0 || x >= n || y < 0 || y >= n || g[x][y] > 0 {
			dx++
			dx %= 4
			x, y = i+dir[dx][0], j+dir[dx][1]
		}
		i, j = x, y
	}

	return g
}

//leetcode submit region end(Prohibit modification and deletion)

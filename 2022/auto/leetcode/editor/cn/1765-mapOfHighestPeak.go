package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	dir := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans[i] = append(ans[i], -1)
		}
	}

	st := [][2]int{}

	for i, arr := range isWater {
		for j, v := range arr {
			if v == 1 {
				st = append(st, [2]int{i, j})
				ans[i][j] = 0
			}
		}
	}

	idx := 0
	for idx < len(st) {
		cur := st[idx]
		for _, d := range dir {
			x, y := cur[0]+d[0], cur[1]+d[1]
			if x >= 0 && x < m && y >= 0 && y < n {
				if ans[x][y] == -1 {
					ans[x][y] = ans[cur[0]][cur[1]] + 1
					st = append(st, [2]int{x, y})
				}
			}
		}
		idx++
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

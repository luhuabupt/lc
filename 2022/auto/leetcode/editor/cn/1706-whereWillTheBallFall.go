package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func findBall(g [][]int) (ans []int) {
	m, n := len(g), len(g[0])

	// -1 右入, 0 上入, 1 左入
	var do func(d, i, j int) int
	do = func(d, i, j int) int {
		if i == m {
			return j
		}
		if j < 0 || j >= n {
			return -1
		}
		if d == 0 {
			return do(g[i][j], i, j+g[i][j])
		}
		if d != g[i][j] {
			return -1
		}
		return do(0, i+1, j)
	}

	for j := 0; j < n; j++ {
		ans = append(ans, do(0, 0, j))
	}

	return
}

//leetcode submit region end(Prohibit modification and deletion)

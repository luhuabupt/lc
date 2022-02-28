package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func matrixScore(g [][]int) int {
	m, n := len(g), len(g[0])

	for i := 0; i < m; i++ {
		if g[i][0] == 0 {
			for j := 0; j < n; j++ {
				g[i][j] ^= 1
			}
		}
	}

	// 1
	ans := m * (1 << (n - 1))
	for j := 1; j < n; j++ {
		cnt := 0
		for i := 0; i < m; i++ {
			cnt += g[i][j]
		}
		if cnt < m-cnt {
			cnt = m - cnt
		}
		ans += cnt * (1 << (n - 1 - j))
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

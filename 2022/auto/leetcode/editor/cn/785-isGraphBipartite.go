package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func isBipartite(g [][]int) bool {
	n := len(g)
	m := make([]int, n)
	for i, _ := range m {
		m[i] = -1
	}

	var dfs func(i int) bool
	dfs = func(i int) bool {
		for _, x := range g[i] {
			if m[i] == m[x] {
				return false
			}
			if m[x] == -1 {
				m[x] = 1 ^ m[i]
				if !dfs(x) {
					return false
				}
			}
		}
		return true
	}

	for i := 0; i < n; i++ {
		if m[i] == -1 {
			m[i] = 0
		}
		if !dfs(i) {
			return false
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func totalNQueens(n int) int {
	ans := 0

	r := make([]bool, n)
	c := make([]bool, n)
	a := make([]bool, n+n)
	b := make([]bool, n+n)

	cl := func(i, j int) int {
		if i >= j {
			return i - j
		}
		return n + j - i
	}
	cr := func(i, j int) int {
		return i + j
	}

	var dfs func(i, j, u int)
	dfs = func(i, j, u int) {
		if u > n {
			return
		}
		if n*n-(i*n+j) < n-u {
			return
		}

		ni, nj := i, j+1
		if nj == n {
			ni, nj = ni+1, 0
		}

		if i == n {
			if u == n {
				ans++
			}
			return
		}

		dfs(ni, nj, u)

		if !r[i] && !c[j] && !a[cl(i, j)] && !b[cr(i, j)] {
			r[i] = true
			c[j] = true
			a[cl(i, j)] = true
			b[cr(i, j)] = true

			dfs(ni, nj, u+1)

			r[i] = false
			c[j] = false
			a[cl(i, j)] = false
			b[cr(i, j)] = false
		}
	}
	dfs(0, 0, 0)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

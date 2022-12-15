package main

import "fmt"

func main() {
	fmt.Println(solveNQueens(1))
	fmt.Println(solveNQueens(2))
	fmt.Println(solveNQueens(4))
}

//leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
	ans := [][]string{}

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

	g := make([][]byte, n)
	for i := 0; i < n; i++ {
		g[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			g[i][j] = '.'
		}
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

		//fmt.Println(i, j, u)
		if i == n {
			if u == n {
				one := make([]string, n)
				for k := 0; k < n; k++ {
					one[k] = string(g[k])
				}
				ans = append(ans, one)
			}
			return
		}

		dfs(ni, nj, u)

		if !r[i] && !c[j] && !a[cl(i, j)] && !b[cr(i, j)] {
			g[i][j] = 'Q'
			r[i] = true
			c[j] = true
			a[cl(i, j)] = true
			b[cr(i, j)] = true

			dfs(ni, nj, u+1)

			g[i][j] = '.'
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

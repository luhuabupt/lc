package main

func largestMagicSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	c, r := grid, grid
	for i, arr := range grid {
		tmp := 0
		for j, v := range arr {
			tmp += v
			c[i][j] = tmp
		}
	}
	for j := 0; j < n; j++ {
		tmp := 0
		for i := 0; i < m; i++ {
			tmp += grid[i][j]
			r[i][j] = tmp
		}
	}

	var check func(i, j, x int) bool
	check = func(i, j, x int) bool {
		a, b := 0, 0
		for v := 0; v <= x; v++ {
			a += grid[i+v][j+v]
			b += grid[i+v][j+x-v]
		}
		if a != b {
			return false
		}

		for v := 0; v <= x; v++ {
			if r[i+v][j+x]-r[i+v][j+x] != a {
				return false
			}
		}
		for v := 0; v <= x; v++ {
			if c[i+x][j+v]-r[i][j+v] != a {
				return false
			}
		}

		return true
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for x := 1; i+x < m && j+x < n; x++ {
				if check(i, j, x) && x > ans {
					ans = x
				}
			}
		}
	}
	return ans + 1
}

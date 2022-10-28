package main

func minOperations(a []int) (ans int) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return a
	}

	p := a[0]
	for _, v := range a {
		p = gcd(v, p)
	}

	get := func(v int) (int, c int, d int) {
		for {
			if v%2 == 0 {
				v /= 2
				c++
			} else if v%3 == 0 {
				v /= 3
				d++
			} else {
				return v, c, d
			}
		}
	}

	cm, dm := 1, 1
	for _, v := range a {
		v /= p
		b, c, d := get(v)
		if b > 1 {
			return -1
		}
		if c > cm {
			cm = c
		}
		if d > dm {
			dm = 1
		}
	}

	for _, v := range a {
		v /= p
		_, c, d := get(v)
		ans += cm - c
		ans += dm - d
	}
	return ans
}

func lakeCount(g []string) int {
	ans := 0
	vis := make([][]bool, len(g))
	n := len(g)
	m := len(g[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if vis[i][j] {
			return
		}
		for _, d := range [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}} {
			x, y := i+d[0], j+d[1]
			if x >= 0 && x < m && j >= 0 && j < n {
				if !vis[x][y] && g[x][y] == 'w' {
					dfs(x, y)
				}
			}
		}
	}

	for i, a := range g {
		vis[i] = make([]bool, len(g[0]))
		for j, v := range a {
			if v == 'w' && !vis[i][j] {
				dfs(i, j)
				ans++
			}
		}
	}
	return ans
}

func numberOfPairs(nums []int) int {
	get := func(v int) int {
		if v == 0 {
			return 0
		}
		a := []int{}
		for x := v; x > 0; x /= 10 {
			a = append(a, x%10)
		}
		r := 0
		for i := len(a) - 1; i >= 0; i-- {
			r *= 10
			r += a[i]
		}
		return r
	}

	c := map[int]int{}
	ans := 0
	M := int(1e9) + 7
	for _, v := range nums {
		nv := get(v)
		ans += c[nv-v]
		ans %= M
		c[nv-v]++
	}
	return ans
}

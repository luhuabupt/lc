package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(longestPath([]int{-1, 0, 1}, "aab"))
}

func digitSum(s string, k int) string {
	a := []byte(s)
	for len(a) > k {
		na := []byte{}
		for i := 0; i < len(a); i += k {
			x := 0
			for j := i; j < i+k && j < len(a); j++ {
				x += int(a[j] - '0')
			}
			na = append(na, strconv.Itoa(x)...)
		}
		a = na
	}
	return string(a)
}

func minimumRounds(t []int) int {
	m := map[int]int{}
	for _, v := range t {
		m[v]++
	}
	ans := 0
	for _, v := range m {
		if v == 1 {
			return -1
		}
		if v%3 == 0 {
			ans += v / 3
		} else if v%3 == 1 {
			ans += 2 + (v-4)/3
		} else {
			ans += 1 + (v-2)/3
		}
	}
	return ans
}

func maxTrailingZeros(g [][]int) int {
	m, n := len(g), len(g[0])
	c := make([][][2]int, m)
	l := make([][][2]int, m)
	r := make([][][2]int, m)

	for i, a := range g {
		c[i] = make([][2]int, n)
		l[i] = make([][2]int, n)
		r[i] = make([][2]int, n)
		for j, v := range a {
			a, b := 0, 0
			for v%2 == 0 {
				v /= 2
				a++
			}
			for v%5 == 0 {
				v /= 5
				b++
			}
			c[i][j] = [2]int{a, b}
			l[i][j] = c[i][j]
			if j > 0 {
				l[i][j][0] += l[i][j-1][0]
				l[i][j][1] += l[i][j-1][1]
			}
		}
	}

	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			r[i][j] = c[i][j]
			if i > 0 {
				r[i][j][0] += r[i-1][j][0]
				r[i][j][1] += r[i-1][j][1]
			}
		}
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			x := [2]int{0, 0}
			if j > 0 {
				x = l[i][j-1]
			}
			ans = max(ans, min(x[0]+r[i][j][0], x[1]+r[i][j][1]))

			ans = max(ans, min(l[i][n-1][0]-l[i][j][0]+r[i][j][0], l[i][n-1][1]-l[i][j][1]+r[i][j][1]))
			ans = max(ans, min(l[i][j][0]+r[m-1][j][0]-r[i][j][0], l[i][j][1]+r[m-1][j][1]-r[i][j][1]))

			y := [2]int{0, 0}
			if j > 0 {
				y = l[i][j-1]
			}
			ans = max(ans, min(l[i][n-1][0]-y[0]+r[m-1][j][0]-r[i][j][0], l[i][n-1][1]-y[1]+r[m-1][j][1]-r[i][j][1]))
		}
	}

	return ans
}

func longestPath(p []int, s string) int {
	n := len(p)
	next := make([][]int, n)

	for i, v := range p {
		if v == -1 {
			continue
		}
		next[v] = append(next[v], i)
	}

	ans := 0

	var dfs func(i int) int
	dfs = func(i int) int {
		a, b := 0, 0
		for _, j := range next[i] {
			x := dfs(j)
			if s[i] == s[j] {
				continue
			}
			if x >= a {
				a, b = x, a
			} else if x > b {
				b = x
			}
		}
		if a+b+1 > ans {
			ans = a + b + 1
		}
		return a + 1
	}
	dfs(0)

	return ans
}

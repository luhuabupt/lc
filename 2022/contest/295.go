package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(totalSteps([]int{10, 1, 2, 3, 4, 5, 6, 1, 2, 3}))     //6
	fmt.Println(totalSteps([]int{7, 14, 4, 14, 13, 2, 6, 13}))        //3
	fmt.Println(totalSteps([]int{5, 14, 15, 2, 11, 5, 13, 15}))       //3
	fmt.Println(totalSteps([]int{5, 3, 4, 4, 7, 3, 6, 11, 8, 5, 11})) //3
}

func rearrangeCharacters(s string, target string) int {
	c := [26]int{}
	for _, v := range s {
		c[v-'a']++
	}
	ans := 0

loop:
	for {
		for _, v := range target {
			if c[v-'a'] == 0 {
				break loop
			}
			c[v-'a']--
		}
		ans++
	}
	return ans
}

func discountPrices(s string, d int) string {
	a := strings.Split(s, " ")
	for i, v := range a {
		if v[0] == '$' && len(v) > 0 {
			x, err := strconv.Atoi(v[1:])
			if err == nil {
				a[i] = fmt.Sprintf("$%.2f", float64(x)/float64(100)*float64(100-d))
			}
		}
	}
	return strings.Join(a, " ")
}

func totalSteps(a []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	ans := 0
	n := len(a)
	st := []int{}
	dp := make([]int, n)
	for i, _ := range dp {
		dp[i] = 1
	}

	o := []int{}
	for i, v := range a {
		if len(o) == 0 || v >= o[len(o)-1] {
			o = append(o, v)
			dp[i] = 0
		}
	}
	fmt.Println(dp)

	for i, v := range a {
		if dp[i] > 0 {
			for len(st) > 0 && a[st[len(st)-1]] <= v {
				end := st[len(st)-1]
				dp[i] = max(dp[i], dp[end]+1)
				st = st[:len(st)-1]
			}
		} else {
			st = []int{i}
		}

		if dp[i] == 0 {
		} else if i > 0 && a[i] > a[i-1] {
			dp[i] = max(dp[i], dp[i-1]+1)
		} else {
			dp[i] = 1
		}
		if dp[i] > ans {
			ans = dp[i]
		}
		st = append(st, i)
		fmt.Println(i, st)
	}
	fmt.Println(dp)
	return ans
}

func minimumObstacles(g [][]int) int {
	m := len(g)
	n := len(g[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}
	dp[0][0] = 0

	next := [][]int{}
	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		for _, d := range dir {
			x, y := i+d[0], j+d[1]
			if x >= 0 && x < m && y >= 0 && y < n {
				do := dp[i][j] + g[x][y]
				if dp[x][y] == -1 || do < dp[x][y] {
					dp[x][y] = do
					if x == m-1 && y == n-1 {
						return
					}
					if g[x][y] == 0 {
						dfs(x, y)
					} else {
						next = append(next, []int{x, y})
					}
				}
			}
		}
	}
	dfs(0, 0)

	st := [][]int{{0, 0}}
	for len(st) > 0 {
		next = [][]int{}
		for _, x := range st {
			dfs(x[0], x[1])
		}
		st = next
	}

	return dp[m-1][n-1]
}

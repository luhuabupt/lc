package main

import (
	"fmt"
	"sort"
)

// ZbAuEH 打地鼠  2022-04-24 22:34:11
func main() {
	fmt.Println(getMaximumNumber([][]int{{1, 1, 0}, {2, 0, 1}, {4, 2, 2}}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func getMaximumNumber(ma [][]int) int {
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	ma = append(ma, []int{0, 0, 1})

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ga := func(x, y int) int {
		return x*3 + y
	}
	get := func(p, time int) []int {
		if time >= 4 {
			return []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
		}
		x, y := p/3, p%3
		vis := [3][3]bool{}

		vis[x][y] = true
		st := [][]int{{x, y}}
		for i := 1; i <= time; i++ {
			ns := [][]int{}
			for _, tmp := range st {
				for _, d := range dir {
					nx, ny := tmp[0]+d[0], tmp[1]+d[1]
					if nx >= 0 && nx < 3 && ny >= 0 && ny < 3 {
						if vis[nx][ny] == false {
							vis[nx][ny] = true
							ns = append(ns, []int{nx, ny})
						}
					}
				}
			}
			st = ns
		}

		ans := []int{}
		for i, a := range vis {
			for j, v := range a {
				if v {
					ans = append(ans, i*3+j)
				}
			}
		}

		return ans
	}

	t := []int{}
	do := map[int][9]bool{}
	for _, v := range ma {
		if _, ok := do[v[0]]; !ok {
			t = append(t, v[0])
			do[v[0]] = [9]bool{}
		}
		x := do[v[0]]
		x[ga(v[1], v[2])] = true
		do[v[0]] = x
	}
	sort.Ints(t)

	dp := make([][9]int, len(t))
	if do[0][4] {
		dp[0][4] = 1
	}
	for i := 1; i < len(t); i++ {
		v := t[i]
		for j := 0; j < 9; j++ {
			pre := get(j, v-t[i-1])
			for _, x := range pre {
				dp[i][j] = max(dp[i][j], dp[i-1][x])
			}
			if do[t[i]][j] == true {
				dp[i][j]++
			}
		}
	}

	ans := 0
	for _, x := range dp[len(dp)-1] {
		ans = max(ans, x)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

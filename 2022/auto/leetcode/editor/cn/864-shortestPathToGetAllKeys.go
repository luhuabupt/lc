package main

import "fmt"

// shortest-path-to-get-all-keys 获取所有钥匙的最短路径  2022-03-02 22:31:22
func main() {
	fmt.Println(shortestPathAllKeys())
}

//leetcode submit region begin(Prohibit modification and deletion)
func shortestPathAllKeys(g []string) int {
	m, n := len(g), len(g[0])
	ans := -1

	key := [][2]int{}
	o := [2]int{}
	for i, arr := range g {
		for j, v := range arr {
			if v >= 'a' && v <= 'f' {
				key = append(key, [2]int{i, j})
			}
			if v == '@' {
				o = [2]int{i, j}
			}
		}
	}

	dir := [][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

	vis := map[int]map[int]map[int][2]int{}
	var get func(s, e [2]int, h int) (int, int)
	get = func(s, e [2]int, h int) (int, int) {
		if vis[s[0]*n+s[1]] == nil {
			vis[s[0]*n+s[1]] = map[int]map[int][2]int{}
		}
		if vis[s[0]*n+s[1]][e[0]*n+e[1]] == nil {
			vis[s[0]*n+s[1]][e[0]*n+e[1]] = map[int][2]int{}
		}
		if vis[s[0]*n+s[1]][e[0]*n+e[1]][h][0] > 0 {
			return vis[s[0]*n+s[1]][e[0]*n+e[1]][h][0], vis[s[0]*n+s[1]][e[0]*n+e[1]][h][1]
		}
		if g[s[0]][s[1]] >= 'a' && g[s[0]][s[1]] <= 'f' {
			h &= 1 << (g[s[0]][s[1]] - 'a')
		}

		mi, rh := -1, 0
		for _, d := range dir {
			x, y := s[0]+d[0], s[1]+d[1]
			if x >= 0 && x < m && y >= 0 && y < n {
				if g[x][y] == '#' {
					continue
				}
				if g[x][y] >= 'A' && g[x][y] <= 'F' {
					if 1<<(g[x][y]-'A')&h == 0 {
						continue
					}
				}

				re, reh := get([2]int{x, y}, e, h)
				if re != -1 && (re < mi || mi == -1) {
					mi, rh = re, rh&reh
				}
			}
		}
		vis[s[0]*n+s[1]][e[0]*n+e[1]][h] = [2]int{mi, rh}

		return mi, rh
	}

	find := func(arr []int) int {
		res := 0
		for i := 0; i < len(arr); i++ {
			x := -1
			if i == 0 {
				x, h = get(o, key[arr[i]], 0)
			} else {
				x, h = get(key[arr[i-1]], key[arr[i]], 0)
			}
			if x == -1 {
				return -1
			}
			res += x
		}

		return res
	}

	do := Factorial(len(key))
	for _, arr := range do {
		x := find(arr)
		if x != -1 && (x < ans || ans == -1) {
			ans = x
		}
	}

	return ans
}

func Factorial(n int) [][]int { // 阶乘
	a := []int{}
	for i := 0; i < n; i++ {
		a = append(a, i)
	}

	ans := [][]int{}

	var dfs func(undo, cur []int)
	dfs = func(undo, cur []int) {
		if len(cur) == n {
			ans = append(ans, cur)
			return
		}
		for i, x := range undo {
			nc := append([]int{}, cur...)
			nc = append(nc, x)

			nd := append([]int{}, undo[:i]...)
			nd = append(nd, undo[i+1:]...)

			dfs(nd, nc)
		}
	}

	dfs(a, []int{})

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

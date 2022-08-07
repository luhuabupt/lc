package main

import "fmt"

func main() {
	fmt.Println(minimumScore([]int{1, 5, 5, 4, 11}, [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}))
	fmt.Println(minimumScore([]int{5, 5, 2, 4, 4, 2}, [][]int{{0, 1}, {1, 2}, {5, 2}, {4, 3}, {1, 3}}))
}

func minimumScore(nums []int, edges [][]int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := len(nums)
	nx := make([][]int, n)
	for _, v := range edges {
		nx[v[0]] = append(nx[v[0]], v[1])
		nx[v[1]] = append(nx[v[1]], v[0])
	}

	dp := make([]int, n)
	vis := make([]bool, n)
	fa := make([]map[int]bool, n) // 所有父节点
	var dfs0 func(i int, c []int)
	dfs0 = func(i int, c []int) {
		vis[i] = true
		rx := nums[i]
		for _, x := range nx[i] {
			if !vis[x] {
				fa[x] = map[int]bool{}
				for _, t := range c {
					fa[x][t] = true
				}
				dfs0(x, append([]int{x}, c...))
				rx ^= dp[x]
			}
		}
		dp[i] = rx
	}
	dfs0(0, []int{0})

	getSon := func(i int) int {
		s := edges[i][0]
		if fa[edges[i][1]][edges[i][0]] {
			s = edges[i][1]
		}
		return s
	}
	cf := func(i, j int) (int, int) {
		if fa[i][j] {
			return j, i
		}
		if fa[j][i] {
			return i, j
		}
		return -1, -1
	}
	cal := func(a, b, c int) int {
		return max(a, max(b, c)) - min(a, min(b, c))
	}

	ans := 1 << 60
	// 枚举n-1条边
	for i := 0; i < n-1; i++ { // 删除第1条
		s0 := getSon(i)
		for j := i + 1; j < n-1; j++ { // 删除第2条
			s1 := getSon(j)
			a, b := cf(s0, s1)
			if a == -1 { // 不在同一个子树上
				r := dp[0] ^ dp[s0] ^ dp[s1]
				ans = min(ans, cal(dp[s0], dp[s1], r))
			} else { // 同一个子树上
				r := dp[0] ^ dp[a]
				ans = min(ans, cal(dp[b], r, dp[a]^dp[b]))
			}
		}
	}

	return ans
}

func countHousePlacements(n int) int {
	M := int(1e9) + 7
	f, u := 0, 1
	for i := 0; i < n; i++ {
		u, f = u+f, u
		u %= M
	}
	u += f
	u %= M
	return u * u % M
}

func maximumsSplicedArray(a []int, b []int) int {
	sa, sb := 0, 0
	for _, v := range a {
		sa += v
	}
	for _, v := range b {
		sb += v
	}

	ab := getMax(a, b)
	ba := getMax(b, a)
	ans := sa
	if sb > ans {
		ans = sb
	}
	if sa+ba > ans {
		ans = sa + ba
	}
	if sb+ab > ans {
		ans = sb + ab
	}
	return ans
}

func getMax(a, b []int) int {
	n := len(a)
	sa, sb := 0, 0
	ans := 0
	d := make([]int, n)

	mi := 0
	for i := 0; i < n; i++ {
		sa += a[i]
		sb += b[i]
		d[i] = sa - sb
		if d[i]-mi > ans {
			ans = d[i] - mi
		}

		if sa-sb < mi {
			mi = sa - sb
		}
	}

	return ans
}

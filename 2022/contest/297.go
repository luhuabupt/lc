package main

import "fmt"

func main() {
	fmt.Println(distributeCookies([]int{8, 15, 10, 20, 8}, 2))
	fmt.Println(distributeCookies([]int{6, 1, 3, 2, 2, 4, 1, 2}, 3))
}

func distributeCookies(c []int, k int) int {
	n := len(c)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	do := make([]int, 1<<n)
	for i := 1; i < 1<<n; i++ {
		x := 0
		for j := 0; j < n; j++ {
			if (1<<j)&i > 0 {
				x += c[j]
			}
		}
		do[i] = x
	}

	ans := 1 << 60
	var dfs func(mask, idx, ma int)
	dfs = func(mask, idx, ma int) {
		if idx == k-1 {
			x := max(do[((1 << n) - 1) ^ mask], ma)
			if x < ans {
				ans = x
			}
			return
		}
		for i, v := range do {
			if i&mask > 0 || v > ans {
				continue
			}
			dfs(mask|i, idx+1, max(ma, v))
		}
	}
	dfs(0, 0, 0)
	return ans
}

func distinctNames(a []string) int64 {
	m := map[string]bool{}
	for _, v := range a {
		m[v] = true
	}

	ans := 0
	f := [26][26]int{} // srg -> to

	for _, v := range a {
		bv := []byte(v)
		for j := 'a'; j <= 'z'; j++ {
			bv[0] = byte(j)
			sv := string(bv)
			if !m[sv] {
				ans += 2 * f[v[0]-'a'][j-'a']
				f[j-'a'][v[0]-'a']++
			}
		}
		//fmt.Println(i, f)
	}

	return int64(ans)
}

func calculateTax(b [][]int, in int) float64 {
	ans := 0.0
	pre := 0
	for i := 0; i < len(b); i++ {
		do := 0
		if in < b[i][0] {
			do = in - pre
		} else {
			do = b[i][0] - pre
		}
		pre = b[i][0]
		ans += float64(do*b[i][1]) / 100.000000
		if in < b[i][0] {
			break
		}
	}
	return ans
}

func minPathCost(g [][]int, mc [][]int) int {
	m, n := len(g), len(g[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[i][j] = g[i][j]
				continue
			}
			mi := 1 << 60
			for x := 0; x < n; x++ {
				if dp[i-1][x]+mc[g[i-1][x]][j] < mi {
					mi = dp[i-1][x] + mc[g[i-1][x]][j]
				}
			}
			dp[i][j] = mi + g[i][j]
		}
	}
	ans := 1 << 60
	for _, v := range dp[m-1] {
		if v < ans {
			ans = v
		}
	}
	return ans
}

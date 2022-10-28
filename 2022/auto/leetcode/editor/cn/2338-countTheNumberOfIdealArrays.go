package main

import "fmt"

func main() {
	// 4
	// 1 12 13 14 24 124
	// 1 C(n-1, l-1)

	// 6
	// 1 12 13 14 15 16 124 126 136
	// 24 26 36

	// 12
	// 1248 126_12 128 1210 12_12

	fmt.Println(idealArrays(4, 10000))
	fmt.Println(idealArrays(2, 5))
	fmt.Println(idealArrays(5, 3))
}

//leetcode submit region begin(Prohibit modification and deletion)
func idealArrays(n int, mv int) int {
	M := int(1e9) + 7
	ans := 0
	c := make([]int, 16)

	var dfs func(i, cur int)
	dfs = func(i, cur int) {
		c[cur]++
		for j := 2; j <= mv; j++ {
			if i*j > mv {
				break
			}
			dfs(i*j, cur+1)
		}
	}

	for i := 1; i <= mv; i++ {
		dfs(i, 1)
	}
	for i, v := range c {
		if v > 0 {
			ans += v * Cnk(n-1, i-1, M)
			ans %= M
		}
	}

	return ans
}

func Cnk(n, k, M int) int {
	qp := func(a, b int) int {
		ans := 1
		a %= M
		for b > 0 {
			if b&1 > 0 {
				ans = ans * a % M
				b--
			}
			b >>= 1
			a = a * a % M
		}
		return ans
	}

	cal := func(n, k int) int {
		if k > n {
			return 0
		}
		ans := 1
		for i := 1; i <= k; i++ {
			a := (n + i - k) % M
			b := i % M
			ans = ans * (a * qp(b, M-2) % M) % M
		}
		return ans
	}

	// 当n, k > M 时
	//var lucas func(n, k int) int
	//lucas = func(n, k int) int {
	//	if k == 0 {
	//		return 1
	//	}
	//	return cal(n%M, k%M) * lucas(n/M, k/M) % M
	//}

	return cal(n, k)
}

//leetcode submit region end(Prohibit modification and deletion)

package main

import "fmt"

func main() {
	fmt.Println(Cnk(1000, 5, int(1e9+7)))
	fmt.Println(8250291250200 % int(1e9+7))
}

// C-n-r 的排列
func Permutation(n, r int) {

}

// https://blog.csdn.net/ACdreamers/article/details/8037918
// https://blog.csdn.net/ArrowLLL/article/details/52629448
// C(n, k) % M
func Cnk(n, k, M int) int {
	var qp func(a, b int) int
	qp = func(a, b int) int {
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

	var cal func(n, k int) int
	cal = func(n, k int) int {
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

	var lucas func(n, k int) int
	lucas = func(n, k int) int {
		if k == 0 {
			return 1
		}
		return cal(n%M, k%M) * lucas(n/M, k/M) % M
	}

	return lucas(n, k)
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

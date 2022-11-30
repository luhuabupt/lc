package main

import "fmt"

func main() {
	fmt.Println(kthSmallestPath([]int{2, 3}, 1))
	fmt.Println(kthSmallestPath([]int{2, 3}, 2)) // HHVHV
}

//leetcode submit region begin(Prohibit modification and deletion)
func kthSmallestPath(dest []int, k int) string {
	m, n := dest[0], dest[1]
	ans := []byte{}

	i, j := 0, 0
	for l := 0; l < m+n; l++ {
		fmt.Println(i, j)
		if i == m {
			ans = append(ans, 'H')
			continue
		}
		if j == n {
			ans = append(ans, 'V')
			continue
		}
		x := cnk(m+n-l-1, n-j-1, int(1e9+7))
		if x < k {
			k -= x
			ans = append(ans, 'V')
			i++
		} else {
			ans = append(ans, 'H')
			j++
		}
	}

	return string(ans)
}

func cnk(n, k, M int) int {
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

//leetcode submit region end(Prohibit modification and deletion)

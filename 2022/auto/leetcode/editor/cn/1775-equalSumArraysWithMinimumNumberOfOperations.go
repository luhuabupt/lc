package main

import (
	"fmt"
	"sort"
)

// equal-sum-arrays-with-minimum-number-of-operations 通过最少操作次数使数组的和相等  2022-12-07 23:14:59
func main() {
	fmt.Println(minOperations([]int{3, 3, 2, 4, 2, 6, 2}, []int{6, 2, 6, 6, 1, 1, 4, 6, 4, 6, 2, 5, 4, 2, 1}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func minOperations(a []int, b []int) int {
	m, n := len(a), len(b)

	if n*6 < m {
		return -1
	}
	if m*6 < n {
		return -1
	}

	mi := m
	if n > mi {
		mi = n
	}

	ma := m * 6
	if ma > n*6 {
		ma = n * 6
	}

	ans := 1 << 60
	l := cal(a, mi, ma)
	r := cal(b, mi, ma)
	for i := mi; i <= ma; i++ {
		if l[i]+r[i] < ans {
			ans = l[i] + r[i]
		}
	}

	return ans
}

func cal(a []int, mi, ma int) []int {
	n := len(a)
	sort.Ints(a)
	oa := append([]int{}, a...)
	s := 0
	for _, v := range a {
		s += v
	}

	mx := ma + 1
	if s+1 > mx {
		mx = s + 1
	}
	dp := make([]int, mx)
	dp[s] = 1
	for d, j := s-1, n-1; d >= mi; d-- {
		if a[j] > 1 {
			dp[d] = dp[d+1]
			a[j]--
		} else {
			j--
			a[j]--
			dp[d] = dp[d+1] + 1
		}
	}

	a = append([]int{}, oa...)
	for d, j := s+1, 0; d <= ma; d++ {
		if a[j] < 6 {
			a[j]++
			dp[d] = dp[d-1]
		} else {
			j++
			a[j]++
			dp[d] = dp[d-1] + 1
		}
	}

	dp[s] = 0
	return dp
}

//leetcode submit region end(Prohibit modification and deletion)

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(" ")
}

func countPartitions(a []int, k int) int {
	M := int(1e9 + 7)
	qp := func(a, n, m int) int {
		ans := 1
		for n > 0 {
			if n&1 > 0 {
				ans *= a
				ans %= m
			}
			a *= a
			a %= m
			n /= 2
		}
		return ans
	}

	sort.Ints(a)
	dp := make([]int, k+1)
	dp[0] = 1
	for _, v := range a {
		for j := k; j-v >= 0; j-- {
			dp[j] += dp[j-v]
			dp[j] %= M
		}
	}

	s := 0
	for _, v := range dp {
		s += v
		s %= M
	}

	ans := qp(2, len(a), M) - 2*s
	return ans % M
}

func maximumTastiness(a []int, k int) int {
	sort.Ints(a)
	n := len(a)
	return sort.Search(1e9+1, func(f int) bool {
		c := 0
		for i := 0; i < n; {
			i = sort.SearchInts(a, a[i]+f)
			c++
			if c >= k {
				return true
			}
		}
		return false
	})
}

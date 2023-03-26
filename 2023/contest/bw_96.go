package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(" ")
}

func maxScore(a []int, b []int, k int) int64 {
	n := len(a)

	//mi := a[n-k]
	s := 0
	for i := n - k + 1; i < n; i++ {
		s += a[i]
	}

	ans := 0

	c := make([][]int, n)
	for i, v := range b {
		c[i] = []int{v, a[i]}
	}
	sort.Slice(c, func(i, j int) bool {
		return c[i][0] < c[j][0]
	})

	for i := 0; i+k-1 < n; i++ {
		ans = max(ans, (s+c[i][1])*c[i][0])
	}

	return int64(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(" ")
}

func canReceiveAllSignals(a [][]int) bool {
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})
	for i := 1; i < len(a); i++ {
		if a[i][0] < a[i][1] {
			return false
		}
	}
	return true
}

func buildTransferStation(a [][]int) int {
	m, n := len(a), len(a[0])

	g := make([][]int, m)
	for i := 0; i < m; i++ {
		g[i] = make([]int, n)
	}

}

func minSwaps(c []int) int {
	t := 0
	s := make([]int, len(c))
	for i, v := range c {
		t += v
		s[i] = t
	}

	ans := t - s[t]
	for i := t; i < len(c); i++ {
		if s[i]-s[i-t] < ans {
			ans = s[i] - s[i-t]
		}
	}

	return ans
}

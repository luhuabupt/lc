package main

import (
	"fmt"
	"math/bits"
)

func deleteString(s string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(s)
	dp := make([]int, n)
	dp[n-1] = 1

	chk := func(i, j int) bool {
		for x := 0; x < j; x++ {
			if s[i+x] != s[i+j+x] {
				return false
			}
		}
		return true
	}
	for i := n - 2; i >= 0; i-- {
		for j := 1; i+j+j <= n; j++ {
			//if s[i:i+j] == s[i+j:i+j+j] {
			if chk(i, j) {
				dp[i] = max(dp[i], dp[i+j]+1)
			}
		}
	}

	return dp[0]
}

func minimizeXor(a int, b int) int {
	c := bits.OnesCount(uint(b))
	ans := 0
	for i := 32; i >= 0; i-- {
		if (1<<i)&a > 0 && c > 0 {
			c--
			ans |= (1 << i)
		}
	}
	for i := 0; i < 32 && c > 0; i++ {
		if (1<<i)&ans > 0 {
			continue
		} else {
			c--
			ans |= (1 << i)
		}
	}
	return ans
}

//[[520626,685427,788912,800638,717251,683428],[23602,608915,697585,957500,154778,209236],[287585,588801,818234,73530,939116,252369]]
//5095181
func maxSum(g [][]int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	m, n := len(g), len(g[0])
	cal := func(i, j int) int {
		if i+2 >= m || j+2 >= n {
			return 0
		}
		r := 0
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				r += g[i+x][i+y]
			}
		}
		r -= g[i+1][j]
		r -= g[i+1][j+1]
		return r
	}
	ans := 0
	for i, ar := range g {
		for j, _ := range ar {
			fmt.Println(i, j, cal(i, j))
			ans = max(ans, cal(i, j))
		}
	}
	return ans
}

package main

import (
	"fmt"
)

func main() {
	//fmt.Println(beautifulPartitions("23542185131", 3, 2))
	//fmt.Println(beautifulPartitions("23542185131", 3, 3))
	fmt.Println(beautifulPartitions("783938233588472343879134266968", 4, 6)) // 4
}

func beautifulPartitions(s string, k int, miLen int) int {
	M := int(1e9) + 7
	tm := map[byte]bool{'2': true, '3': true, '5': true, '7': true}
	n := len(s)
	dp := make([][]int, n)
	sum := make([][]int, n)
	if !tm[s[0]] || tm[s[n-1]] {
		return 0
	}
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k)
		sum[i] = make([]int, k)
	}
	for i := miLen - 1; i < n; i++ {
		if !tm[s[i]] && (i == n-1 || tm[s[i+1]]) {
			//fmt.Println(i, string(s[i]))
			dp[i][0] = 1
			if i >= miLen {
				for j := k - 1; j >= 1; j-- {
					dp[i][j] += sum[i-miLen][j-1]
					dp[i][j] %= M
				}
			}
		}
		for j := 0; j < k; j++ {
			sum[i][j] += sum[i-1][j] + dp[i][j]
			sum[i][j] %= M
		}
	}
	//fmt.Println(dp)
	return dp[n-1][k-1]
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func minimumFuelCost(g [][]int, st int) int64 {
	n := len(g) + 1
	nx := make([][]int, n)
	for _, v := range g {
		nx[v[0]] = append(nx[v[0]], v[1])
		nx[v[1]] = append(nx[v[1]], v[0])
	}

	vis := make([]bool, n)
	ans := 0
	var dfs func(i int) int
	dfs = func(i int) int {
		vis[i] = true
		r := 0
		for _, x := range nx[i] {
			if vis[x] {
				continue
			}
			r += dfs(x)
		}
		ans += r / st
		if r%st != 0 {
			ans++
		}
		return r
	}
	dfs(0)

	return int64(ans)
}

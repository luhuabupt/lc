package main

import (
	"fmt"
	"sort"
)

type Abc struct {
	node []*Abc
}

func main() {
	fmt.Println(findMaxUp([]int{1, 2, 3}))
	fmt.Println(findMaxUp([]int{1, 2, 2}))
	fmt.Println(findMaxUp([]int{1, 2, 2, 2, 3, 3}))
}

// 最长上升子序列
func findMaxUp(arr []int) int {
	dp := []int{}
	for _, v := range arr {
		if p := sort.SearchInts(dp, v+1); p < len(dp) {
			dp[p] = v
		} else {
			dp = append(dp, v)
		}
	}
	return len(dp)
}

package main

import (
	"fmt"
	"sort"
)

// stack represents a stack of program counters.
type stack []uintptr

type withStack struct {
	error
	*stack
}

type withMessage struct {
	cause error
	msg   string
}

type TaskMsg struct {
	ActionType int   `json:"actionType"`
	ActionId   int64 `json:"actionId"`
	Mid        int64 `json:"mid"`
	Ts         int64 `json:"date"`
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

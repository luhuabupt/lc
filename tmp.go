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
	for v := 1; v <= 9; v++ {
		for i := 1; i <= 3; i++ { //m
			fmt.Println(v, i, sort.Search(3, func(vv int) bool {
				return (vv+1)*i >= v
			}))
		}
	}
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

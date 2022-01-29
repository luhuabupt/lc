package main

import "fmt"

func main() {
	fmt.Println(minJumps([]int{51, 64, -15, 58, 98, 31, 48, 72, 78, -63, 92, -5, 64, -64, 51, -48, 64, 48, -76, -86, -5, -64, -86, -47, 92, -41, 58, 72, 31, 78, -15, -76, 72, -5, -97, 98, 78, -97, -41, -47, -86, -97, 78, -97, 58, -41, 72, -41, 72, -25, -76, 51, -86, -65, 78, -63, 72, -15, 48, -15, -63, -65, 31, -41, 95, 51, -47, 51, -41, -76, 58, -81, -41, 88, 58, -81, 88, 88, -47, -48, 72, -25, -86, -41, -86, -64, -15, -63}))
}

func minJumps(arr []int) int {
	n := len(arr)
	dp := make([]int, n)
	m := map[int][]int{}

	for i, v := range arr {
		if m[v] == nil {
			m[v] = []int{}
		}
		m[v] = append(m[v], i)
		dp[i] = -1
	}

	stack := []int{0}
	dp[0] = 0
	idx := 0

	for idx < len(stack) && dp[n-1] == -1 {
		cur := stack[idx]
		if cur > 0 && dp[cur-1] == -1 {
			stack = append(stack, cur-1)
			dp[cur-1] = dp[cur] + 1
		}
		if cur+1 < n && dp[cur+1] == -1 {
			stack = append(stack, cur+1)
			dp[cur+1] = dp[cur] + 1
		}
		if m[arr[cur]] != nil {
			for _, x := range m[arr[cur]] {
				if x != cur && dp[x] == -1 {
					stack = append(stack, x)
					dp[x] = dp[cur] + 1
				}
			}
			delete(m, arr[cur])
		}
		idx++
	}

	return dp[n-1]
}

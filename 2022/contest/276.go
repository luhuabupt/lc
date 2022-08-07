package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxRunTime(2, []int{1, 3, 1, 100}))
}

func divideString(s string, k int, fill byte) []string {
	arr := []string{}
	sa := []byte(s)
	for len(sa)%k != 0 {
		sa = append(sa, fill)
	}
	for i := 0; i < len(sa); i++ {
		if i%k == k-1 {
			arr = append(arr, string(sa[i-k+1:i+1]))
		}
	}
	return arr
}

func minMoves(target int, maxDoubles int) int {
	if target == 1 || maxDoubles == 0 {
		return target - 1
	}
	if target%2 == 1 {
		return minMoves(target-1, maxDoubles) + 1
	}

	return minMoves(target/2, maxDoubles-1) + 1
}

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := map[int]int64{}
	for i := n - 1; i >= 0; i-- {
		dp[i] = dp[i+1]
		if int64(questions[i][0])+dp[i+questions[i][1]+1] > dp[i] {
			dp[i] = int64(questions[i][0]) + dp[i+questions[i][1]+1]
		}
	}
	return dp[0]
}

func maxRunTime(n int, batteries []int) int64 {
	sort.Ints(batteries)
	m := len(batteries)

	sum := 0
	for i := 0; i < m-n; i++ {
		sum += batteries[i]
	}

	min := batteries[m-n]
	for i := m - n + 1; i < m; i++ { // 后n个电脑, 依次分配sum, 取一个最小值
		if sum > (batteries[i]-batteries[i-1])*(i-m+n) {
			sum -= (batteries[i] - batteries[i-1]) * (i - m + n)
			min = batteries[i]
		} else {
			min += sum / (i - m + n)
			sum = 0
			break
		}
	}
	min += sum / n // batteries[m] < min 的情况

	return int64(min)
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(123)
}

func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	l, r := make([]int, n), make([]int, n) // 蜡烛位置
	sum := map[int]int{}                   // 前缀和
	pre := -1

	for i := 0; i < n; i++ {
		if s[i] == '*' {
			l[i] = pre
		} else {
			pre = i
		}
	}
	pre = -1
	for i := n - 1; i >= 0; i-- {
		if s[i] == '*' {
			r[i] = pre
		} else {
			pre = i
		}
	}

	idx := 0
	for s[idx] != '|' {
		idx++
	}
	tmp := 0
	for i := idx + 1; i < n; i++ {
		if s[i] == '*' {
			tmp++
		} else {
			sum[i] = tmp
		}
	}

	ans := []int{}
	for _, arr := range queries {
		li, ri := r[arr[0]], l[arr[1]]
		x := 0
		if li == -1 || ri == -1 || li >= ri {

		} else {
			x = sum[ri] - sum[li]
		}
		ans = append(ans, x)
	}

	return ans
}

func maxTwoEvents(events [][]int) int {
	pm := map[int]bool{}
	points := []int{}
	s, e := map[int][]int{}, map[int][]int{}
	for i, arr := range events {
		pm[arr[0]] = true
		pm[arr[1]] = true
		if s[arr[0]] == nil {
			s[arr[0]] = []int{}
		}
		if e[arr[1]] == nil {
			e[arr[1]] = []int{}
		}
		s[arr[0]] = append(s[arr[0]], i)
		e[arr[1]] = append(e[arr[1]], i)
	}

	for v, _ := range pm {
		points = append(points, v)
	}
	sort.Ints(points)

	ans := 0
	dp := map[int]int{}
	preDpi := -1
	for _, p := range points {
		for _, i := range s[p] {
			if events[i][2]+dp[preDpi] > ans {
				ans = dp[preDpi] + events[i][2]
			}
		}
		for _, i := range e[p] {
			if events[i][2] > dp[preDpi] && events[i][2] > dp[p] {
				dp[p] = events[i][2]
			}
			preDpi = p
		}
	}
	return ans
}

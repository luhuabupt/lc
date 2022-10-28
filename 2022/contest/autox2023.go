package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minCostToTravelOnDays([]int{1, 2, 3, 4}, [][]int{{1, 3}, {2, 5}, {3, 7}}))
	fmt.Println(minCostToTravelOnDays([]int{1, 4, 5}, [][]int{{1, 4}, {5, 6}, {2, 5}}))
	fmt.Println(minCostToTravelOnDays([]int{1, 4, 5}, [][]int{{2, 3}, {1, 2}}))
	fmt.Println(minCostToTravelOnDays([]int{2, 4, 5}, [][]int{{2, 3}, {1, 2}}))
	fmt.Println(minCostToTravelOnDays([]int{2, 4, 6}, [][]int{{2, 3}, {1, 2}}))
}

func minCostToTravelOnDays(days []int, tickets [][]int) int64 {
	//filter()
	n := len(days)
	dp := make([]int, n)
	for i, v := range days {
		mi := 1 << 60
		for _, t := range tickets {
			idx := sort.SearchInts(days, v-t[0]+1)
			x := t[1]
			if idx > 0 {
				x += dp[idx-1]
			}
			if x < mi {
				mi = x
			}
			//fmt.Println(i, v, t, idx, x, mi)
		}
		dp[i] = mi
	}
	return int64(dp[n-1])
}

func honeyQuotes(a [][]int) []float64 {
	ans := []float64{}
	c := map[int]int{}
	s := 0
	t := 0
	for _, b := range a {
		switch b[0] {
		case 1:
			c[b[1]]++
			s += b[1]
			t++
		case 2:
			c[b[1]]--
			if c[b[1]] == 0 {
				delete(c, b[1])
			}
			s -= b[1]
			t--
		case 3:
			if t == 0 {
				ans = append(ans, -1)
			} else {
				ans = append(ans, float64(s)/float64(t))
			}
		case 4:
			if t == 0 {
				ans = append(ans, -1)
			} else {
				avg := float64(s) / float64(t)
				sqt := 0.0
				for v, cv := range c {
					sqt += (float64(v) - avg) * (float64(v) - avg) * float64(cv)
				}
				sqt /= float64(t)
				ans = append(ans, sqt)
			}
		}
	}
	return ans
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(smallestSubarrays([]int{1, 0, 2, 1, 3}))
}

func minimumMoney(t [][]int) int64 {
	obm := 0
	pm := 0
	ans := 0
	for _, v := range t {
		if v[1] > v[0] {
			if v[0] > obm {
				obm = v[0]
			}
		} else {
			if v[1] > pm {
				pm = v[1]
			}
			ans += v[0] - v[1]
		}
	}
	ans += pm

	if pm < obm {
		ans += obm - pm
	}

	return int64(ans)
}

func smallestSubarrays(a []int) []int {
	n := len(a)
	d := [31][]int{}
	for i, v := range a {
		for j := 0; j < 31; j++ {
			if v&(1<<j) > 0 {
				d[j] = append(d[j], i)
			}
		}
	}

	ans := make([]int, n)
	c := [31]int{}
	for i, _ := range a {
		ma := 1
		for j := 0; j < 31; j++ {
			if c[j] >= len(d[j]) {
				continue
			}
			if d[j][c[j]] < i {
				c[j]++
			}
			if c[j] >= len(d[j]) {
				continue
			}
			if d[j][c[j]]-i+1 > ma {
				ma = d[j][c[j]] - i + 1
			}
		}
		ans[i] = ma
	}

	return ans
}

func matchPlayersAndTrainers(a []int, b []int) int {
	sort.Ints(a)
	sort.Ints(b)
	ans := 0

	for i, j := 0, 0; i < len(a); i++ {
		for ; j < len(b); j++ {
			if a[i] <= b[j] {
				ans++
				j++
				break
			}
		}
	}

	return ans
}

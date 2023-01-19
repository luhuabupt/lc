package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(" ")
}

func maxJump(a []int) int {
	n := len(a)
	mx := 0
	i, j := 0, 0
	u := make([]bool, n)
	for {
		if a[i] < a[j] {
			for u[i] {
				i++
			}
			if i != n-1 {
				u[i] = true
			}
		} else {
			for u[i] {
				j++
			}
			if j != n-1 {
				u[j] = true
			}
		}
		if i == n-1 && j == n-1 {
			break
		}
	}

	return mx
}

func maxStarSum(a []int, g [][]int, k int) int {
	n := len(a)
	nx := make([][]int, n)
	for _, v := range g {
		nx[v[0]] = append(nx[v[0]], a[v[1]])
		nx[v[1]] = append(nx[v[1]], a[v[0]])
	}

	ans := -1 << 60
	for i := 0; i < n; i++ {
		sort.Ints(nx[i])
		t := a[i]
		for j := 0; j < k && j < len(nx[i]); j++ {
			t += nx[i][len(nx[i])-1-j]
		}
		if t > ans {
			ans = t
		}
	}
	return ans
}

func maximumValue(s []string) int {
	ans := 0
	for _, v := range s {
		t, err := strconv.Atoi(v)
		if err != nil {
			if len(v) > ans {
				ans = len(v)
			}
		} else {
			if t > ans {
				ans = t
			}
		}
	}
	return ans
}

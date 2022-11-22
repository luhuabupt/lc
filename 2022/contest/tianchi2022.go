package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(" ")
}
func brilliantSurprise(p [][]int, limit int) int {
	n := len(p)
	s := make([]int, n)
	for i, a := range p {
		st := 0
		for j, v := range a {
			if j >= limit {
				break
			}
			st += v
		}
		s[i] = st
	}

	for i, a := range p {
		dp := make([]int, limit+1)
		for j, b := range p {
			if i == j {
				continue
			}

		}
	}
}

func arrangeBookshelf(a []int, limit int) []int {
	c := map[int]int{}
	for _, v := range a {
		c[v]++
	}

	ans := []int{}
	st := []int{}
	d := map[int]int{}
	for _, v := range a {
		c[v]--
		if c[v]+d[v] >= limit {
			for len(st) > 0 && v < st[len(st)-1] {
				st = st[:len(st)-1]
			}
			st = append(st, v)
		} else {
			x := sort.SearchInts(st, v+1)
			for j := 0; j <= x && j < len(st); j++ {
				if d[st[j]] < limit {
					ans = append(ans, st[j])
					d[st[j]]++
				}

			}
			if d[v] <= limit {
				d[v]++
				ans = append(ans, v)
			}
			st = []int{}
		}

		fmt.Println(v, st, ans, c[v])
	}

	return ans
}

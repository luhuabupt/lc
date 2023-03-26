package main

import "sort"

func putMarbles(a []int, k int) int64 {
	n := len(a)
	b := []int{}
	for i := 1; i < n; i++ {
		b = append(b, a[i]+a[i-1])
	}
	sort.Ints(b)
	c, d := 0, 0
	for i := 0; i < k-1; i++ {
		c += b[i]
		d += b[len(b)-1-i]
	}
	return int64(d - c)
}

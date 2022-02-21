package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(goodTriplets([]int{4, 0, 1, 3, 2}, []int{4, 1, 0, 2, 3}))
}

func goodTriplets(a []int, b []int) int64 {
	n := len(a)
	ma, mb := make([]int, n), make([]int, n)
	for i, v := range a {
		ma[v] = i
	}
	for i, v := range b {
		mb[v] = i
	}

	c := []int{}
	ans := 0
	l := make([]int, n)
	r := make([]int, n)
	for _, v := range a {
		x := sort.SearchInts(c, mb[v])
		c = append(c, mb[v])
		//fmt.Println(v, x, c)
		l[v] = x
	}

	c = []int{}
	for i := n - 1; i >= 0; i-- {
		v := a[i]
		x := sort.SearchInts(c, mb[v])
		r[v] = len(c) - x
		c = append(c, mb[v])
		fmt.Println(v, mb[v], x, c, r[v])
	}
	for i := 0; i < n; i++ {
		fmt.Println(i, l[i], r[i])
		ans += l[i] * r[i]
	}

	return int64(ans)
}

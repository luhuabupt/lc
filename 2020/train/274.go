package main

import (
	"fmt"
	"sort"
)

func numberOfBeams(bank []string) int {
	m, n := len(bank), len(bank[0])
	pre, cur := 0, 0
	ans := 0
	for i := 0; i < m; i++ {
		cur = 0
		for j := 0; j < n; j++ {
			if bank[i][j] == '1' {
				cur++
			}
		}
		if cur > 0 {
			ans += cur * pre
			pre = cur
		}
	}
	return ans
}

func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Ints(asteroids)
	for _, v := range asteroids {
		if mass < v {
			return false
		}
		mass += v
	}
	return true
}

func main() {
	//fmt.Println(maximumInvitations([]int{1, 0, 0, 2, 1, 4, 7, 8, 9, 6, 7, 10, 8}))
	//fmt.Println(maximumInvitations([]int{1,0,3,2,5,6,7,4,9,8,11,10,11,12,10}))
	//fmt.Println(maximumInvitations([]int{6, 10, 10, 0, 6, 0, 4, 4, 1, 3, 3}))
	fmt.Println(maximumInvitations([]int{1, 0, 3, 2, 5, 6, 7, 4, 9, 8, 11, 10, 11, 12, 10})) //11
	//fmt.Println(maximumInvitations([]int{3,0,1,4,1}))
	//fmt.Println(maximumInvitations([]int{1,2,0}))
	//req := []int{23, 14, 17, 5, 19, 2, 0, 18, 15, 12, 2, 8, 21, 3, 3, 1, 6, 5, 11, 17, 3, 7, 14, 13}
	req := []int{5, 13, 3, 14, 1, 1, 13, 13, 14, 10, 5, 13, 3, 14, 2}

	for i, v := range req {
		fmt.Println(i, v)
	}
	fmt.Println(maximumInvitations(req))
}

func maximumInvitations(favorite []int) int {
	ans := 0
	n := len(favorite)
	vis := make([]bool, n)
	pre := make([][]int, n)

	var tw func(i int) int
	tw = func(i int) int {
		x := 0
		for _, v := range pre[i] {
			if !vis[v] {
				x = max(tw(v), x)
			}
		}
		vis[i] = true
		return x + 1
	}

	cl := func(i int) int {
		s, f := favorite[i], favorite[favorite[i]]
		for s != f {
			s = favorite[s]
			f = favorite[favorite[f]]
			if vis[s] || vis[f] {
				return 0
			}
		}
		cl := 1
		for x := favorite[s]; x != s; x = favorite[x] {
			cl++
			vis[x] = true
		}
		for x := i; x != s; x = favorite[x] {
			vis[x] = true
		}
		return cl
	}

	// sum circle = 2
	for i, v := range favorite {
		pre[v] = append(pre[v], i)
	}
	for i, v := range favorite {
		if i < favorite[i] && i == favorite[favorite[i]] {
			vis[v], vis[favorite[v]] = true, true
			ans += tw(v) + tw(favorite[v])
		}
	}

	// get circle
	for i, _ := range favorite {
		if !vis[i] {
			ans = max(cl(i), ans)
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

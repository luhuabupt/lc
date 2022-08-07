package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(minSumSquareDiff([]int{1, 4, 10, 12}, []int{5, 8, 6, 9}, 3, 1))
	//fmt.Println(minSumSquareDiff([]int{10, 10, 10, 11, 5}, []int{1, 0, 6, 6, 1}, 11, 27)) // 0
	fmt.Println(validSubarraySize([]int{1, 3, 4, 3, 1}, 6))
}

func validSubarraySize(a []int, ts int) int {
	n := len(a)
	b := make([]int, n)
	for i, v := range a {
		if v > ts {
			return 1
		}
		b[i] = ts/v + 1
	}
	fmt.Println(b)

	l, r := make([]int, n), make([]int, n)
	st := []int{}
	for i := 0; i < n; i++ {
		for len(st) > 0 && b[st[len(st)-1]] <= b[i] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			l[i] = -1
		} else {
			l[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	st = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && b[st[len(st)-1]] <= b[i] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			r[i] = n
		} else {
			r[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	fmt.Println(l, r)

	ans := -1
	for i, v := range b {
		if i-l[i]+r[i]-i-1 >= v {
			ans = v
			break
		}
	}
	return ans
}

func minSumSquareDiff(a []int, b []int, k1 int, k2 int) int64 {
	abs := func(a int) int {
		if a > 0 {
			return a
		}
		return -a
	}
	n := len(a)
	k := k1 + k2
	c := make([]int, n)
	for i, v := range a {
		c[i] = abs(b[i] - v)
	}
	sort.Ints(c)
	for i := 0; i < n/2; i++ {
		c[i], c[n-1-i] = c[n-1-i], c[i]
	}
	fmt.Println(c)

	av := sort.Search(1e5+1, func(ma int) bool {
		t := 0
		for _, v := range c {
			if v > ma {
				t += v - ma
			} else {
				break
			}
		}
		return t <= k
	})

	for i, v := range c {
		if v > av {
			k -= v - av
			c[i] = av
		} else {
			break
		}
	}
	for i := 0; i < n && k > 0; i++ {
		if c[i] >= 1 {
			c[i]--
			k--
		}
	}
	fmt.Println(av, c)

	ans := 0
	for _, v := range c {
		ans += v * v
	}
	return int64(ans)
}

func latestTimeCatchTheBus(bus []int, p []int, c int) int {
	sort.Ints(bus)
	sort.Ints(p)
	ans := 0
	h := map[int]bool{}
	for _, v := range p {
		h[v] = true
	}
	n, m := len(bus), len(p)
	// fmt.Println(bus)
	// fmt.Println(p)

	for i, j := 0, -1; i < n; i++ {
		tc := 0
		for j+1 < m && p[j+1] <= bus[i] && tc <= c {
			tc++
			j++
		}
		//fmt.Println(i, j)
		if i == n-1 {
			x := 0
			if tc < c {
				x = bus[i]
			} else {
				x = p[j]
			}
			for ; x >= 0; x-- {
				if !h[x] {
					ans = x
					break
				}
			}
		}
	}

	return ans
}

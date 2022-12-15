package main

import (
	"fmt"
)

func main() {
	fmt.Println(secondGreaterElement([]int{2, 4, 0, 9, 6}))
	fmt.Println(secondGreaterElement([]int{3, 3}))
}

func secondGreaterElement(a []int) []int {
	n := len(a)
	ans := make([]int, n)
	dp := make([][][2]int, n)
	st := []int{}
	for i := n - 1; i >= 0; i-- {
		ans[i] = -1
		v := a[i]
		for len(st) > 0 && v >= a[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			ans[i] = -1
		} else {
			dp[st[len(st)-1]] = append(dp[st[len(st)-1]], [2]int{v, i})
		}
		st = append(st, i)
	}
	//fmt.Println(dp)

	st = []int{}
	for i := n - 1; i >= 0; i-- {
		//fmt.Println(i, a[i], st)
		if len(dp[i]) > 0 && len(st) > 0 {
			for _, x := range dp[i] {
				//y := sort.SearchInts(st, x[0]+1)
				//y := sort.Search(len(st), func(xx int) bool {
				//	return st[xx] > x[0]
				//})
				//fmt.Println(x, y)
				//if y == 0 {
				//	ans[x[1]] = -1
				//} else {
				//	ans[x[1]] = st[y]
				//}
				ans[x[1]] = -1
				l, r := 0, len(st)-1
				for l <= r {
					m := (l + r) / 2
					if st[m] > x[0] {
						l = m + 1
						ans[x[1]] = st[m]
					} else {
						r = m - 1
					}
				}
			}
		}
		v := a[i]
		for len(st) > 0 && v > st[len(st)-1] {
			st = st[:len(st)-1]
		}
		st = append(st, v)
	}
	return ans
}

func destroyTargets(a []int, s int) int {
	m := map[int][]int{}
	ms := map[int]int{}

	for _, v := range a {
		m[v%s] = append(m[v%s], v)
		if _, ok := ms[v%s]; !ok {
			ms[v%s] = v
		}
		if v < ms[v%s] {
			ms[v%s] = v
		}
	}

	ma := 0
	ans := 0
	for k, v := range m {
		if len(v) > ma {
			ma = len(v)
			ans = ms[k]
		}
	}
	return ans
}

func twoEditWords(a []string, b []string) []string {
	ans := []string{}

	chk := func(a, b string) bool {
		d := 0
		for i, v := range a {
			if b[i] != uint8(v) {
				d++
			}
		}
		return d <= 2
	}

loop:
	for _, v := range a {
		for _, x := range b {
			if chk(v, x) {
				ans = append(ans, v)
				continue loop
			}
		}
	}
	return ans
}

func oddString(a []string) string {
	n := len(a)
	if n == 2 {
		return a[0]
	}
	t := make([][]int, n)
	for j, v := range a {
		for i := 1; i < len(v); i++ {
			t[j] = append(t[j], int(v[i]-v[i-1]))
		}
	}

	for i := 1; i < n-1; i++ {
		for j, v := range t[i] {
			if v != t[i-1][j] && v != t[i+1][j] {
				return a[i]
			}
		}
	}
	for j, v := range t[0] {
		if v != t[1][j] {
			return a[0]
		}
	}
	return a[n-1]
}

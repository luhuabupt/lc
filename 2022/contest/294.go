package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(totalStrength([]int{1, 3, 1, 2}))
	fmt.Println(totalStrength([]int{5, 4, 6}))
	fmt.Println(13*5 + 12*4 + 14*6)
}

func totalStrength(ar []int) int {
	M := int(1e9 + 7)
	n := len(ar)

	p := make([][2]int, n)
	st := []int{}
	stv := []int{}
	for i, v := range ar {
		for len(st) > 0 && v < stv[len(st)-1] {
			st = st[:len(st)-1]
			stv = stv[:len(stv)-1]
		}
		st = append(st, i)
		stv = append(stv, v)
		if len(st) == 1 {
			p[i][0] = -1
		} else {
			p[i][0] = st[len(st)-2]
		}
	}

	st = []int{}
	stv = []int{}
	for i := n - 1; i >= 0; i-- {
		v := ar[i]
		for len(st) > 0 && v <= stv[len(st)-1] {
			st = st[:len(st)-1]
			stv = stv[:len(stv)-1]
		}
		st = append(st, i)
		stv = append(stv, v)
		if len(st) == 1 {
			p[i][1] = -1
		} else {
			p[i][1] = st[len(st)-2]
		}
	}

	fmt.Println(p)

	dp := make([]int, n)
	for i, v := range ar {
		x := p[i]
		dp[x[0]+1] += v
		if x[1] != -1 {
			dp[x[1]] -= v
		}
	}

	fmt.Println(dp)

	ans := 0
	t := 0
	for i, v := range dp {
		t += dp[i]
		t %= M

		ans += v * t
		ans %= M
	}

	return ans
}

func minimumLines(sp [][]int) int {
	sort.Slice(sp, func(i, j int) bool {
		return sp[i][0] < sp[j][0]
	})

	ans := 1
	n := len(sp)
	for i := 2; i < n; i++ {
		a := sp[i-2][0]
		b := sp[i-2][1]
		c := sp[i-1][0]
		d := sp[i-1][1]
		e := sp[i][0]
		f := sp[i][1]

		if (e-c)*(d-b) != (c-a)*(f-d) {
			ans++
		}
	}
	return ans
}

func maximumBags(c []int, r []int, add int) int {
	n := len(c)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = c[i] - r[i]
	}
	sort.Ints(a)
	ans := 0
	for _, v := range a {
		if add < v {
			break
		}
		add -= v
		ans++
	}
	return ans
}

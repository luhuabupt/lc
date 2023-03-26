package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(" ")
}

func maxPower(a []int, r int, k int) int64 {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := len(a)
	t := 0
	s := make([]int, n)
	for i, v := range a {
		t += v
		s[i] = t
	}

	ans := sort.Search(3e10, func(find int) bool {
		d := make([]int, n)
		ca := 0
		do := 0
		for i, _ := range a {
			x := ca + s[min(i+r, n-1)]
			if i-r != 0 {
				x -= s[max(0, i-r-1)]
			}
			if x < find {
				d[i] += find - x
				do += find - x
				d[min(n-1, i+r+r+1)] -= find - x
			}
			ca += d[i]
		}
		return do < k
	})
	return int64(ans)
}

func categorizeBox(a int, b int, c int, d int) string {
	ib := func() bool {
		if a >= 1e4 || b >= 1e4 || c >= 1e4 || a*b*c >= 1e9 {
			return true
		}
		return false
	}

	id := func() bool {
		if d >= 100 {
			return true
		}
		return false
	}

	bv := ib()
	dv := id()

	if bv && dv {
		return "Both"
	}
	if !bv && !dv {
		return "Neither"
	}
	if bv {
		return "Bulky"
	}
	return "Heavy"
}

type DataStream struct {
	k, v, c int
	a       []int
}

func Constructor(value int, k int) DataStream {
	return DataStream{k, value, 0, []int{}}
}

func (t *DataStream) Consec(num int) bool {
	t.a = append(t.a, num)
	if len(t.a) < t.k {
		return false
	}
	if num == t.v {
		t.c++
	}
	if len(t.a)-1-t.k >= 0 && t.a[len(t.a)-1-t.k] == t.v {
		t.c--
	}
	return t.c == t.k
}

/**
 * Your DataStream object will be instantiated and called as such:
 * obj := Constructor(value, k);
 * param_1 := obj.Consec(num);
 */

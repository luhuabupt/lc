package main

import (
	"fmt"
	"sort"
)

func main() {
	ori := []int{2, 3, 0, 1}
	a := []int{}
	for i := 0; i < 4; i++ {
		a = append(a, i)
	}
	sort.Slice(a, func(i, j int) bool {
		return ori[i] < ori[j]
	})
	fmt.Println(a) // 0 2 1 3

	x := []string{"102", "473", "251", "814"}
	//sort.Strings(x)
	//	//fmt.Println(x)
	ns := []string{}
	for _, v := range x {
		ns = append(ns, string(append([]byte{}, v[2:]...)))
	}
	d := []int{}
	for i := 0; i < 4; i++ {
		d = append(d, i)
	}
	fmt.Println(d, ns)
	sort.Slice(d, func(i, j int) bool {
		return ns[i] < ns[j]
	})
	fmt.Println(d)

	//fmt.Println(canChange("_L__R__R_", "L______RR"))
	fmt.Println(smallestTrimmedNumbers([]string{"102", "473", "251", "814"}, [][]int{{1, 1}})) // 2 0 1 3
}
func smallestTrimmedNumbers(a []string, q [][]int) []int {
	type tt struct {
		idx int
		ss  string
	}
	n := len(q)
	ans := make([]int, n)

	d := make([][]tt, len(a[0])+1)
	m := len(a)
	lm := len(a[0])
	fmt.Println(a[0][lm-1:], a[2][lm-1:], a[0][lm-1:] < a[2][lm-1:])
	for l := 1; l <= len(a[0]); l++ {
		d[l] = make([]tt, m)
		for i := 0; i < len(a); i++ {
			d[l][i] = tt{i, a[i][lm-l:]}
		}
		fmt.Println(d[l])
		sort.Slice(d[l], func(i, j int) bool {
			return d[l][i].ss < d[l][j].ss || d[l][i].ss == d[l][j].ss && d[l][i].idx < d[l][j].idx
		})
		fmt.Println(l, lm-l, a[0][lm-l:], d[l])
	}
	fmt.Println(d)
	for i, v := range q {
		k, t := v[0], v[1]
		ans[i] = d[t][k-1].idx
	}
	return ans
}

func minOperations(a []int, b []int) int {
	n := len(b)
	t := b[0]
	for i := 1; i < n; i++ {
		t = gcd(t, b[i])
	}
	sort.Ints(a)
	ans := -1
	for i, v := range a {
		if t%v == 0 {
			return i + 1
		}
	}
	return ans
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func idealArrays(n int, ma int) int {
	M := int(1e9) + 7
	cal := func(v int) []int {
		r := []int{}
		for i := 1; i*i <= v; i++ {
			if v%i == 0 {
				r = append(r, i)
			}
		}
		for j := len(r) - 1; j >= 0; j-- {
			if r[j]*r[j] != v {
				r = append(r, v/r[j])
			}
		}
		return r
	}
	cn1i := func(v int) int { // C n-1 v
		if v == n-1 {
			return 1
		}
		if v == 1 {
			return 1
		}
		r, x := 1, 1
		for i := 1; i <= v; i++ {
			x *= i
		}
		for i := 1; i <= v; i++ {
			r *= n - i
		}
		return r / x
	}

	ans := 0
	d := make([][]int, ma+1)
	for i := 1; i <= ma; i++ {
		d[i] = make([]int, 202)
	}
	for i := 1; i <= ma; i++ {
		c := cal(i)
		d[i][1] = 1
		for _, v := range c { // gcd
			if v == i {
				continue
			}
			for l, x := range d[v] {
				if l > 0 && x == 0 {
					break
				}
				d[i][l+1] += x
			}
		}
		fmt.Println(i, c, d[i])
		for i := 1; i <= ma; i++ {
			for j := 1; j < 202; j++ {
				if d[i][j] == 0 {
					break
				}
				fmt.Println(i, j, d[i][j]*cn1i(j-1))
				ans += d[i][j] * cn1i(j)
				ans %= M
			}
		}
	}
	return ans
}

func canChange(s string, t string) bool {
	n := len(s)
	a, b := []int{}, []int{}
	for i := 0; i < n; i++ {
		if s[i] != '_' {
			a = append(a, i)
		}
		if t[i] != '_' {
			b = append(b, i)
		}
	}

	fmt.Println(a, b)
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if s[a[i]] != t[b[i]] {
			return false
		}
		if s[a[i]] == 'L' {
			if a[i] < b[i] {
				return false
			}
		} else {
			if a[i] > b[i] {
				return false
			}
		}
	}

	return true
}

type SmallestInfiniteSet struct {
	h map[int]bool
}

func Constructor() SmallestInfiniteSet {
	h := map[int]bool{}
	for i := 1; i < 1002; i++ {
		h[i] = true
	}
	return SmallestInfiniteSet{h}
}

func (t *SmallestInfiniteSet) PopSmallest() int {
	for i := 0; i < 1002; i++ {
		if t.h[i] {
			t.h[i] = false
			return i
		}
	}
	return 0
}

func (t *SmallestInfiniteSet) AddBack(num int) {
	t.h[num] = true
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */

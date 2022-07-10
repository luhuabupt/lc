package main

import "fmt"

func main() {
	//fmt.Println(canChange("_L__R__R_", "L______RR"))
	fmt.Println(idealArrays(2, 5))
	fmt.Println(idealArrays(5, 3))
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
				fmt.Println(i, j, d[i][j] * cn1i(j-1))
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

package main

import "fmt"

func main() {
	//fmt.Println(canChange("_L__R__R_", "L______RR"))
	fmt.Println(idealArrays(2, 5))
	fmt.Println(idealArrays(5, 3))
	fmt.Println(idealArrays(5, 9)) // 111
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

	cd := make([][]int, n+20)
	for i := 0; i < len(cd); i++ {
		cd[i] = make([]int, 16)
	}
	var cnm func(a, b int) int
	cnm = func(a, b int) int {
		if b > a {
			return 0
		}
		if b == 0 {
			return 1
		}
		if b == 1 {
			return a
		}
		if a == b {
			return 1
		}
		if cd[a][b] > 0 {
			return cd[a][b]
		}
		x := cnm(a, b-1) + cnm(a-1, b)
		cd[a][b] = x % M
		return x
	}

	ans := 0
	d := make([][]int, ma+1)
	for i := 1; i <= ma; i++ {
		d[i] = make([]int, 16)
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
		for j := 1; j < 16; j++ {
			if d[i][j] == 0 {
				break
			}
			ans += d[i][j] * cnm(n-1, j-1)
			ans %= M
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

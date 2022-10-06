package main

import "fmt"

func main() {
	fmt.Println(countSpecialNumbers(12))
	fmt.Println(countSpecialNumbers(135))
}

func countSpecialNumbers(n int) int {
	if n == 0 {
		return 0
	}
	get := func(v int) []int {
		r := []int{}
		for x := v; x > 0; x /= 10 {
			r = append(r, x%10)
		}
		for i := 0; i < len(r)/2; i++ {
			r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
		}
		return r
	}
	a := get(n)

	d := 0
	t := 0
	for i, v := range a {
		d = countSpecialNumbers(t-1) * (10 - i)
		u := [10]bool{}
		for j := 0; j <= i; j++ {
			if u[a[j]] {
				break
			} else {
				u[a[j]] = true
			}
		}
		for j := 0; j <= v; j++ {
			if !u[j] {
				d++
			}
		}

		t *= 10
		t += v
	}
	p := 9
	s := 1
	for i := 1; i < len(a); i++ {
		s *= p
		d += s
		p--
	}
	fmt.Println(n, d)
	return d
}

func smallestNumber(p string) string {
	n := len(p)
	a := []int{0}
	for i, v := range p {
		if v == 'I' {
			a = append(a, i+1)
		} else {
			t := 0
			for j := i - 1; j >= 0; j-- {
				if p[j] == 'I' {
					t = j + 1
					break
				}
			}
			//fmt.Println(i, t)
			tt := append([]int{}, a[t:]...)
			//fmt.Println(a, tt)
			a = append([]int{}, a[:t]...)
			a = append(a, i+1)
			a = append(a, tt...)
			//fmt.Println(a)
		}
	}
	//fmt.Println(a)
	ans := make([]byte, n+1)
	for i, v := range a {
		ans[i] = byte(v + '1')
	}
	return string(ans)
}

func largestLocal(g [][]int) [][]int {
	n := len(g)
	ans := make([][]int, n-2)

	get := func(i, j int) int {
		ma := g[i][j]
		for x := -1; x < 2; x++ {
			for y := -1; y < 2; y++ {
				if g[i+x][y+j] > ma {
					ma = g[i+x][y+j]
				}
			}
		}
		return ma
	}

	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {
			ans[i-1][j-1] = get(i, j)
		}
	}
	return ans
}

func edgeScore(e []int) int {
	n := len(e)
	c := make([]int, n)
	for i, v := range e {
		c[v] += i
	}
	ans := 0
	for i, v := range c {
		if v > c[i] {
			ans = i
		}
	}
	return ans
}

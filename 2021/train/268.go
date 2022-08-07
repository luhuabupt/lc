package main

import (
	"fmt"
	"sort"
)

type RangeFreqQuery struct {
	m [][]int
}

func Constructor(arr []int) RangeFreqQuery {
	m := make([][]int, 10000)
	for i, v := range arr {
		m[v] = append(m[v], i)
	}
	return RangeFreqQuery{
		m,
	}
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	tmp := this.m[value]
	if len(tmp) == 0 {
		return 0
	}
	return sort.SearchInts(tmp, right+1) - sort.SearchInts(tmp, left-1)
}

func main() {
	fmt.Println(kMirror(3, 7))
}

func kMirror(k int, n int) int64 {
	tk, t0 := []int{}, []int{}
	ak, a0 := []int{}, []int{}
	for i := 1; i < k; i++ {
		tk = append(tk, i)
		ak = append(ak, i)
	}
	for i := 1; i < 10; i++ {
		t0 = append(t0, i)
		a0 = append(a0, i)
	}

	lk, l0 := 2, 2
	di, dj := 0, 0
	ans := []int{}
	loop := 0
	for loop < 200 && len(ans) < n {
		loop++
		fmt.Println(di, dj, ak[di], a0[dj])
		if ak[di] == a0[dj] {
			fmt.Println("hit__________")
			ans = append(ans, ak[di])
			di++
			dj++
		} else if ak[di] > a0[dj] {
			dj++
		} else {
			di++
		}
		if di == len(ak) {
			p := pow(k, lk-1)
			fmt.Println("in tk", lk, tk, "p", p)
			nt := []int{}
			for ii, v := range tk {
				pp := p
				x := v
				for v > 0 {
					x += (v % k) * pp
					v /= k
					pp /= k
				}
				fmt.Println("x", x)
				if lk%2 == 0 {
					ak = append(ak, x)
				} else {
					for iii := 0; iii < k; iii++ {
						nt = append(nt, tk[ii]+iii*pp)
						ak = append(ak, x+iii*pp)
					}
				}
			}
			if lk%2 == 1 {
				tk = nt
			}
			lk++
			fmt.Println("akkkkkkk", ak, "tk", tk)
		}
		if dj == len(a0) {
			p := pow(10, l0-1)
			nt := []int{}
			for ii, v := range t0 {
				pp := p
				x := v
				for v > 0 {
					x += (v % 10) * pp
					v /= 10
					pp /= 10
				}
				if l0%2 == 0 {
					a0 = append(a0, x)
				} else {
					nt := []int{}
					for iii := 0; iii < 10; iii++ {
						nt = append(nt, t0[ii]+iii*pp)
						a0 = append(a0, x+iii*pp)
					}
				}
			}
			if l0%2 == 1 {
				t0 = nt
			}
			fmt.Println(l0, "a00000000000", a0, t0)
			l0++
		}
	}

	fmt.Println(ans)
	a := int64(0)
	for i := 0; i < n; i++ {
		a += int64(ans[i])
	}
	return a
}

func pow(a, n int) int {
	ans := 1
	for n > 0 {
		if n&1 > 0 {
			ans *= a
		}
		a *= a
		n /= 2
	}
	return ans
}

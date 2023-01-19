package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(" ")
}

func maxPoints(g [][]int, q []int) []int {
	m, n := len(g), len(g[0])
	end := 0
	do := map[int]bool{}
	for _, ar := range g {
		for _, v := range ar {
			if v >= g[0][0] {
				do[v+1] = true
			}
			if v > end {
				end = v

			}
		}
	}
	end += 2

	ad := []int{}
	for k, _ := range do {
		ad = append(ad, k)
	}
	sort.Ints(ad)

	vis := make([]bool, m*n)

	nx := make([]map[int]bool, end)
	for i, _ := range nx {
		nx[i] = map[int]bool{}
	}
	nx[g[0][0]+1][0] = true

	cnt := 0
	var dfs func(i, j, d int)
	dfs = func(i, j, d int) {
		cnt++
		if cnt > 63000 {
			fmt.Println(cnt)
		}
		//fmt.Println(i, d, d)
		if vis[i*n+j] {
			return
		}
		if g[i][j] >= d {
			nx[g[i][j]+1][i*n+j] = true
			return
		}

		nx[d][i*n+j] = true
		vis[i*n+j] = true

		for _, dv := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			x, y := i+dv[0], j+dv[1]
			if x >= 0 && x < m && y >= 0 && y < n {
				if !vis[x*n+y] {
					if g[x][y] >= d {
						nx[g[x][y]+1][x*n+y] = true
						continue
					}
					dfs(x, y, d)
				}
			}
		}

		fmt.Println(i, j, d)
	}

	a := make([]int, end)
	for _, d := range ad {
		for k, _ := range nx[d] {
			dfs(k/n, k%n, d)
		}
		a[d] = len(nx[d])
		//fmt.Println(d, nx)
	}

	t := 0
	for i, v := range a {
		t += v
		a[i] = t
	}

	ans := make([]int, len(q))
	for i, v := range q {
		if v >= len(a) {
			ans[i] = a[len(a)-1]
		} else {
			ans[i] = a[v]
		}
	}

	return ans
}

type Allocator struct {
	a []int
}

func Constructor(n int) Allocator {
	return Allocator{make([]int, n)}
}

func (t *Allocator) Allocate(size int, idx int) int {
	c := 0
	s := 0
	for i, v := range t.a {
		if v == 0 {
			if c == 0 {
				s = i
			}
			c++
			if c == size {
				for j := s; j < s+size; j++ {
					t.a[j] = idx
				}
				return s
			}
		} else {
			c = 0
		}
	}
	return -1
}

func (t *Allocator) Free(idx int) int {
	c := 0
	for i, v := range t.a {
		if v == idx {
			c++
			t.a[i] = 0
		}
	}
	return c
}

/**
 * Your Allocator object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Allocate(size,mID);
 * param_2 := obj.Free(mID);
 */

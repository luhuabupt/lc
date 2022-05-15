package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(divisorSubstrings(430043, 2))
	//[[8051,8057],[8074,8089],[7994,7995],[7969,7987],[8013,8020],[8123,8139],[7930,7950],[8096,8104],[7917,7925],[8027,8035],[8003,8011]]
	//9854
	// 126
}

func divisorSubstrings(num int, k int) int {
	a := []int{}
	for x := num; x > 0; x /= 10 {
		a = append(a, x%10)
	}
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-1-i] = a[len(a)-1-i], a[i]
	}

	fmt.Println(a)

	ans := 0
	for i := 0; i < len(a)-k+1; i++ {
		x := 0
		for j := 0; j < k; j++ {
			x *= 10
			x += a[i+j]
		}
		fmt.Println(i, x)
		if x == 0 {
			continue
		}
		if num%x == 0 {
			ans++
		}
	}
	return ans
}

func waysToSplitArray(a []int) int {
	n := len(a)
	d := make([]int, n)
	s := 0
	for i, v := range a {
		s += v
		d[i] = s
	}

	ans := 0
	for i := 0; i < n-1; i++ {
		if d[i] > s-d[i] {
			ans++
		}
	}

	return ans
}

func maximumWhiteTiles(t [][]int, c int) int {
	n := len(t)
	sort.Slice(t, func(i, j int) bool {
		return t[i][0] < t[j][0]
	})

	w := make([]int, n)
	for i, v := range t {
		if i == 0 {
			w[i] = v[0] - 1
		} else {
			w[i] = w[i-1] + v[0] - t[i-1][1] - 1
		}
	}
	ans := 0

	for i, v := range t {
		do := c
		if i == n-1 || v[0] + c <= t[i+1][0] {
			do = v[1]-v[0]+1
		} else {
			r := v[0] + c -1
			x := sort.Search(n, func(f int) bool {
				return t[f][1] >= v[0]+c
			})
			if x == n {
				do -= w[n-1] + r-t[n-1][1]
			} else {
				do -= w[x]
				if r < t[x][0] {
					do += t[x][0] - r - 1
				}
			}
			do += w[i]
		}

		if do > ans {
			ans = do
		}
	}

	return ans
}

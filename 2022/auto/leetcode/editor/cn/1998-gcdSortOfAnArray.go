package main

import (
	"fmt"
	"sort"
)

// gcd-sort-of-an-array 数组的最大公因数排序  2022-10-22 10:49:38
func main() {
	fmt.Println("")
}

//leetcode submit region begin(Prohibit modification and deletion)
func gcdSort(nums []int) bool {
	n := len(nums)
	a := make([][2]int, n)
	c := make([][]int, 1e5+1)
	for i, v := range nums {
		a[i] = [2]int{v, i}
		c[v] = append(c[v], i)
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})

	t := newUnionFind(n)
	for l := 2; l < 5e4+1; l++ {
		r := -1
		for j := l; j < 1e5+1; j += l {
			for _, x := range c[j] {
				if r == -1 {
					r = x
					continue
				}
				t.union(r, x)
			}
		}
	}

	for i, v := range a {
		if i != v[1] && t.find(i) != t.find(v[1]) {
			return false
		}
	}
	return true
}

type unionFind struct {
	size int
	fa   []int
	son  [][]int
}

func newUnionFind(n int) *unionFind {
	fa := make([]int, n)
	son := make([][]int, n)
	for i := 0; i < n; i++ {
		fa[i] = i
		son[i] = append(son[i], i)
	}
	return &unionFind{n, fa, son}
}

func (t *unionFind) find(i int) int {
	if t.fa[i] == i {
		return i
	}
	return t.find(t.fa[i])
}

func (t *unionFind) union(i, j int) {
	fi, fj := t.find(i), t.find(j)
	if fi == fj {
		return
	}

	if len(t.son[fi]) < len(t.son[fj]) {
		fi, fj = fj, fi
	}

	for _, v := range t.son[fj] {
		t.fa[v] = fi
	}

	t.son[fi] = append(t.son[fi], t.son[fj]...)
	t.son[fj] = []int{}
	t.size--
}

//leetcode submit region end(Prohibit modification and deletion)
func explorationSupply(s []int, p []int) []int {
	sort.Ints(s)
	n := len(s)
	ans := make([]int, n)
	for i, v := range p {
		x := sort.SearchInts(s, v)
		if x >= n {
			ans[i] = n - 1
		} else if x == 0 {
			ans[i] = 0
		} else {
			if v-s[x-1] <= s[x]-v {
				ans[i] = x - 1
			} else {
				ans[i] = x
			}

		}
	}
	return ans
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println(numSimilarGroups([]string{"tars", "rats", "arts", "star"}))
}

//leetcode submit region begin(Prohibit modification and deletion)

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

func numSimilarGroups(strs []string) int {
	n := len(strs)
	m := map[string]bool{}
	for _, v := range strs {
		m[v] = true
	}

	check := func(a, b string) bool {
		if a == b {
			return true
		}
		d1, d2 := []byte{}, []byte{}
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				d1 = append(d1, a[i])
				d2 = append(d2, b[i])
				if len(d1) > 2 {
					return false
				}
			}
		}
		return d1[1] == d2[0] && d1[0] == d2[1]
	}

	t := newUnionFind(n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if check(strs[i], strs[j]) {
				t.union(i, j)
			}
		}
	}

	return t.size
}

//leetcode submit region end(Prohibit modification and deletion)

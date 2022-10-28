package main

import (
	"sort"
)

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
type uf1697 struct {
	size int
	fa   []int
	son  [][]int
}

func newUf1697(n int) *uf1697 {
	fa := make([]int, n)
	son := make([][]int, n)
	for i := 0; i < n; i++ {
		fa[i] = i
		son[i] = append(son[i], i)
	}
	return &uf1697{n, fa, son}
}

func (t *uf1697) find(i int) int {
	if t.fa[i] == i {
		return i
	}
	return t.find(t.fa[i])
}

func (t *uf1697) union(i, j int) {
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

func distanceLimitedPathsExist(n int, e [][]int, q [][]int) []bool {
	ans := make([]bool, len(q))

	for i, _ := range q {
		q[i] = append(q[i], i)
	}
	sort.Slice(q, func(i, j int) bool {
		return q[i][2] < q[j][2]
	})
	sort.Slice(e, func(i, j int) bool {
		return e[i][2] < e[j][2]
	})

	//fmt.Println(ae)
	//fmt.Println(q)

	t := newUf1697(n)

	j := 0
	for _, v := range q {
		for j < len(e) && e[j][2] < v[2] {
			t.union(e[j][0], e[j][1])
			j++
		}
		ans[v[3]] = t.find(v[0]) == t.find(v[1])
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

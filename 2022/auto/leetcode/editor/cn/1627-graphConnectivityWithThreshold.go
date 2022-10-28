package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
type uf1627 struct {
	size int
	fa   []int
	son  [][]int
}

func newUf1627(n int) *uf1627 {
	fa := make([]int, n)
	son := make([][]int, n)
	for i := 0; i < n; i++ {
		fa[i] = i
		son[i] = append(son[i], i)
	}
	return &uf1627{n, fa, son}
}

func (t *uf1627) find(i int) int {
	if t.fa[i] == i {
		return i
	}
	return t.find(t.fa[i])
}

func (t *uf1627) union(i, j int) {
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

func areConnected(n int, t int, q [][]int) []bool {
	ans := make([]bool, len(q))
	if t == 0 {
		for i, _ := range ans {
			ans[i] = true
		}
		return ans
	}

	o := newUf1627(n + 1)
	for l := t + 1; l <= n/2; l++ {
		for i := l * 2; i <= n; i += l {
			o.union(l, i)
		}
	}

	for i, v := range q {
		if o.find(v[0]) == o.find(v[1]) {
			ans[i] = true
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

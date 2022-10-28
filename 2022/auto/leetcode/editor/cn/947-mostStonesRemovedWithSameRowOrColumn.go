package main

func main() {

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

func removeStones(stones [][]int) int {
	n := len(stones)
	t := newUnionFind(len(stones))
	for i, v := range stones {
		for j := i + 1; j < len(stones); j++ {
			c := stones[j]
			if v[0] == c[0] || v[1] == c[1] {
				t.union(i, j)
			}
		}
	}
	return n - t.size
}

//leetcode submit region end(Prohibit modification and deletion)

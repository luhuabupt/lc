package main

import "sort"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func countSmaller(a []int) []int {
	n := len(a)
	ans := make([]int, n)
	a = offline(a)
	t := tree{make([]int, 1e5)}
	for i := n - 1; i >= 0; i-- {
		t.add(a[i], 1)
		ans[i] = t.query(0, a[i]-1)
	}

	return ans
}

type tree struct {
	arr []int
}

func (t tree) add(i, v int) {
	for i < len(t.arr) {
		t.arr[i] = t.arr[i] + v
		i += lowbit(i)
	}
}

func (t tree) sum(i int) (ans int) {
	for i > 0 {
		ans += t.arr[i]
		i -= lowbit(i)
	}
	return
}

func (t tree) query(l, r int) (ans int) {
	return t.sum(r) - t.sum(l-1)
}

func lowbit(v int) int {
	return v & -v
}

func offline(a []int) []int {
	//a := []int{-1e9, -100, -1, 0, 1, 33, 1e8, 1e9}
	n := len(a)

	b := [][]int{}
	for i, v := range a {
		b = append(b, []int{v, i})
	}
	sort.Slice(b, func(i, j int) bool {
		return b[i][0] < b[j][0] || b[i][0] == b[j][0] && b[i][1] < b[j][1]
	})
	rea := make([]int, n)
	for i, idx := 0, 0; i < n; i++ {
		if i == 0 || b[i][0] > b[i-1][0] {
			idx++
		}
		rea[b[i][1]] = idx
	}

	return rea
}

//leetcode submit region end(Prohibit modification and deletion)

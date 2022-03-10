package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
type RLEIterator struct {
	a   []int
	idx int
}

func Constructor(encoding []int) RLEIterator {
	return RLEIterator{encoding, 0}
}

func (t *RLEIterator) Next(n int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for x := n; t.idx < len(t.a); {
		diff := min(x, t.a[t.idx])
		t.a[t.idx] -= diff
		x -= diff
		if x == 0 {
			return t.a[t.idx+1]
		}
		t.idx += 2
	}

	return -1
}

/**
 * Your RLEIterator object will be instantiated and called as such:
 * obj := Constructor(encoding);
 * param_1 := obj.Next(n);
 */
//leetcode submit region end(Prohibit modification and deletion)

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
type StockSpanner struct {
	d  int
	st [][]int
}

func Constructor() StockSpanner {
	return StockSpanner{0, [][]int{}}
}

func (t *StockSpanner) Next(x int) int {
	start := 0

	for len(t.st) > 0 {
		cur := t.st[len(t.st)-1]
		if cur[0] > x {
			start = cur[1] + 1
			break
		}
		t.st = t.st[:len(t.st)-1]
	}
	t.st = append(t.st, []int{x, t.d})
	t.d++

	return t.d - start
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
//leetcode submit region end(Prohibit modification and deletion)

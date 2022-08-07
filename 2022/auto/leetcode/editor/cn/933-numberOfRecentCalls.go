package main

import "sort"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
type RecentCounter struct {
	a []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.a = append(this.a, t)
	return len(this.a) - sort.SearchInts(this.a, t-3000)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
//leetcode submit region end(Prohibit modification and deletion)

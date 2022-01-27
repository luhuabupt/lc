package main

import (
	"fmt"
	"math/rand"
)

func main() {
	obj := Constructor(4, []int{2, 1})
	for i := 0; i < 20; i++ {
		fmt.Println(obj.Pick())
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
type Solution struct {
	n int
	m map[int]int
}

func Constructor(n int, blacklist []int) Solution {
	m := map[int]int{}
	vis := map[int]bool{}
	for _, v := range blacklist {
		vis[v] = true
	}

	idx := n - 1
	for _, v := range blacklist {
		if v >= n - len(blacklist) {
			continue
		}
		for vis[idx] {
			idx--
		}
		m[v] = idx
		idx--
	}

	return Solution{n - len(blacklist), m}
}

func (t *Solution) Pick() int {
	x := rand.Intn(t.n)
	if t.m[x] > 0 {
		return t.m[x]
	}
	return x
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
//leetcode submit region end(Prohibit modification and deletion)

package main

import (
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
)

func main() {
	//fmt.Println(oddEvenJumps([]int{10, 13, 12, 14, 15}))
	//fmt.Println(oddEvenJumps([]int{2, 3, 1, 1, 4}))
	fmt.Println(oddEvenJumps([]int{1, 2, 3, 2, 1, 4, 4, 5})) // 6
	//fmt.Println(oddEvenJumps([]int{5,1,3,4,2}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func oddEvenJumps(a []int) int {
	n := len(a)
	le := make([]int, n)
	ge := make([]int, n)
	for i := 0; i < n; i++ {
		le[i] = -1
		ge[i] = -1
	}

	t := redblacktree.NewWithIntComparator()
	for i := n - 1; i >= 0; i-- {
		l, has := t.Floor(a[i])
		if has {
			le[i] = l.Value.(int)
		}
		r, has := t.Ceiling(a[i])
		if has {
			ge[i] = r.Value.(int)
		}

		t.Put(a[i], i)
	}

	dp := make([][2]bool, n)
	var dfs func(i, loop int) bool
	dfs = func(i, loop int) bool {
		if i == -1 {
			return false
		}
		if i == n-1 {
			return true
		}
		if dp[i][loop] {
			return true
		}

		var nx int
		if loop == 1 { // g
			nx = ge[i]
		} else {
			nx = le[i]
		}

		dp[i][loop] = dfs(nx, loop^1)
		return dp[i][loop]
	}

	ans := 0
	for i := 0; i < n; i++ {
		x := dfs(i, 1)
		if x {
			ans++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

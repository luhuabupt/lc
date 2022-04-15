package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minJump([]int{2, 5, 1, 1, 1, 1}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func minJump(jump []int) int {
	n := len(jump)
	dp := make([]int, n)

	st := make([][]int, 10002)
	st[0] = []int{}
	pre := make([][]int, n)
	for i, v := range jump {
		if i+v >= n {
			dp[i] = 1
			st[1] = append(st[1], i)
		} else {
			pre[i+v] = append(pre[i+v], i)
		}
	}

	idx := 0
	for {
		idx++
		sort.Ints(st[idx])

		for i := st[idx][0]; i < n; i++ { // left
			if dp[i] == 0 {
				dp[i] = idx + 1
				st[idx+1] = append(st[idx+1], i)
			}
		}
		for _, v := range st[idx] { // right
			for _, x := range pre[v] {
				if x == 0 {
					return idx + 1
				}
				if dp[x] == 0 {
					dp[x] = idx + 1
					st[idx+1] = append(st[idx+1], x)
				}
			}
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

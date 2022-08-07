package main

import "fmt"

func main() {
	fmt.Println(canCross([]int{0, 1, 3, 5, 6, 8, 12, 17}))
	fmt.Println(canCross([]int{0, 1, 2, 3, 4, 8, 9, 11}))
	fmt.Println(canCross([]int{0, 2}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func canCross(stones []int) bool {
	n := len(stones)
	h := make(map[int]bool, n)
	vis := map[int]map[int]bool{}
	for _, v := range stones {
		h[v] = true
		vis[v] = map[int]bool{}
	}

	var dfs func(i, k int) bool
	dfs = func(i, k int) bool {
		if i == stones[n-1] {
			return true
		}
		f := false
		for x := -1; x <= 1; x++ {
			if k+x <= 0 {
				continue
			}
			if h[i+k+x] && !vis[i+k+x][k+x] {
				f = dfs(i+k+x, k+x)
				if f {
					return f
				}
			}
		}
		vis[i][k] = true
		return f
	}

	return dfs(0, 0)
}

//leetcode submit region end(Prohibit modification and deletion)

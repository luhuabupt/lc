package main

import "fmt"

func main() {
	fmt.Println(waysToBuildRooms([]int{-1, 0, 0, 1, 2}))
	fmt.Println(waysToBuildRooms([]int{-1, 0, 1}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func waysToBuildRooms(pre []int) int {
	n := len(pre)
	M := int(1e9) + 7

	nx := make([][]int, n)
	for i := 1; i < n; i++ {
		nx[pre[i]] = append(nx[pre[i]], i)
	}

	var dfs func(i int) (int, int)
	dfs = func(i int) (int, int) {
		for _, x := range nx[i] {
			t, b := dfs(x)
			if r == 0 {
				r = t
				continue
			}
			r = r * (t + 1)
			r %= M
		}
		if i > 0 {
			r++
		}
		return r
	}

	return dfs(0) % M
}

//leetcode submit region end(Prohibit modification and deletion)

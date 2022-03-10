package main

import "fmt"

func main() {
	fmt.Println(reachableNodes([][]int{{0, 1, 10}, {0, 2, 1}, {1, 2, 2}}, 6, 3))
}

//leetcode submit region begin(Prohibit modification and deletion)
func reachableNodes(edges [][]int, mm int, n int) int {
	dp := make([]int, n)
	d := make([][]int, n)
	do := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1e16
		d[i] = make([]int, n)
		do[i] = make([]int, n)
	}
	dp[0] = 0

	next := make([][]int, n)
	for _, v := range edges {
		next[v[0]] = append(next[v[0]], v[1])
		next[v[1]] = append(next[v[1]], v[0])
		d[v[0]][v[1]] = v[2]
		d[v[1]][v[0]] = v[2]
	}

	st := []int{0}
	for {
		for _, v := range st {

		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

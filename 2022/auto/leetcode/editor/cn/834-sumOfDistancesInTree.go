package main

import "fmt"

func main() {
	fmt.Println(sumOfDistancesInTree(6, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}))
	fmt.Println(sumOfDistancesInTree(1, [][]int{}))
	fmt.Println(sumOfDistancesInTree(2, [][]int{{1, 0}}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func sumOfDistancesInTree(n int, edges [][]int) []int {
	w := make([][]int, n)
	for _, v := range edges {
		w[v[0]] = append(w[v[0]], v[1])
		w[v[1]] = append(w[v[1]], v[0])
	}

	b := make([]int, n) // 后面有几个节点
	vis := make([]bool, n)
	var dfs func(i int) int
	dfs = func(i int) int {
		vis[i] = true
		x := 0
		for _, v := range w[i] {
			if vis[v] {
				continue
			}
			x += dfs(v)
			x += b[v] + 1
			b[i] += b[v] + 1
		}

		return x
	}

	ans := make([]int, n)
	fa := dfs(0)
	ans[0] = fa

	var cal func(i, pre int)
	cal = func(i, pre int) {
		if i != 0 {
			ans[i] = pre - b[i] + (n - 2 - b[i])
		}

		for _, v := range w[i] {
			if ans[v] == 0 {
				cal(v, ans[i])
			}
		}
	}
	cal(0, fa)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

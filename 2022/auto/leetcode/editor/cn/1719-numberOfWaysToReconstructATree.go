package main

import "fmt"

func main() {
	//fmt.Println(checkWays([][]int{{1, 5}, {1, 3}, {2, 3}, {2, 4}, {3, 5}, {3, 4}})) // 2
	//fmt.Println(checkWays([][]int{{1,2},{2,3},{2,4},{1,5}})) // 0
	//fmt.Println(checkWays([][]int{{1,2}})) // 2
	//fmt.Println(checkWays([][]int{{2,3},{1,3},{1,4},{1,2},{4,5},{2,4},{1,5},{3,5},{3,4}})) // 2
	fmt.Println(checkWays([][]int{{1, 5}, {2, 3}, {4, 5}, {2, 5}, {3, 5}})) // 2
}

//leetcode submit region begin(Prohibit modification and deletion)
func checkWays(pairs [][]int) (ans int) {
	t := make([][]int, 501)
	h := map[int]int{}
	for _, v := range pairs {
		t[v[0]] = append(t[v[0]], v[1])
		t[v[1]] = append(t[v[1]], v[0])

		h[v[0]]++
		h[v[1]]++
	}
	n := len(h)

	root := []int{}
	for k, v := range h {
		if v == n-1 {
			root = append(root, k)
		}
	}
	if len(root) == 0 {
		return 0
	}

	flag := false
	vis := make([]bool, 501)
	var dfs func(i int) bool
	dfs = func(i int) bool {
		//fmt.Println(i)
		for _, v := range t[i] {
			if vis[v] {
				continue
			}

			fa := -1
			for _, x := range t[v] {
				if !vis[x] && (fa == -1 || h[x] > h[fa]) {
					fa = x
				}
			}

			//fmt.Println(v, fa)

			if h[fa] == h[i] || fa == -1 && h[v] == h[i] {
				flag = true
			}
			if fa == -1 {
				continue
			}

			vis[fa] = true
			if !dfs(fa) {
				return false
			}
		}

		return true
	}

	vis[root[0]] = true
	if !dfs(root[0]) {
		return 0
	}

	if flag {
		return 2
	}
	return 1
}

//leetcode submit region end(Prohibit modification and deletion)

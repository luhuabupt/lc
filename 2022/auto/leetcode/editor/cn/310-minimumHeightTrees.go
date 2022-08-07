package main

import "fmt"

func main() {
	fmt.Println(findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}))
	fmt.Println(findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {2, 3}, {3, 4}, {5, 4}}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findMinHeightTrees(n int, edges [][]int) (ans []int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	next := make([][]int, n)
	in := make([]int, n)
	for _, v := range edges {
		in[v[0]]++
		in[v[1]]++
		next[v[0]] = append(next[v[0]], v[1])
		next[v[1]] = append(next[v[1]], v[0])
	}

	st := make([]int, n)
	for i, v := range in {
		if v == 1 {
			st = append(st, i)
		}
	}

	for len(st) > 0 {
		ns := make([]int, n)
		for _, v := range st {
			in[v]--
			if in[v] == 1 {

			}
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

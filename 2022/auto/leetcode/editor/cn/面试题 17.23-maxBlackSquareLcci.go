package main

import "fmt"

// max-black-square-lcci 最大黑方阵  2022-01-15 10:34:38
func main() {
	x := findSquare([][]int{{1, 1, 1, 1, 0, 1, 0, 1, 1, 1}, {1, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 1, 0, 1, 0, 1}, {1, 1, 1, 1, 0, 0, 0, 0, 0, 0}, {1, 0, 1, 0, 1, 1, 1, 1, 1, 1}, {1, 1, 0, 0, 1, 0, 1, 0, 0, 1}, {0, 0, 0, 1, 1, 1, 0, 1, 0, 1}, {0, 0, 0, 1, 0, 1, 0, 1, 0, 1}, {1, 0, 1, 0, 1, 1, 0, 1, 1, 1}, {1, 1, 1, 0, 1, 0, 0, 1, 1, 1}})
	fmt.Println(x)
}

//leetcode submit region begin(Prohibit modification and deletion)
func findSquare(matrix [][]int) []int {
	n := len(matrix)
	ans := []int{0, 0, 0}

	rs := map[int]map[int]int{-1: {}}
	cs := map[int]map[int]int{-1: {}}
	for i := 0; i < n; i++ {
		rs[i] = map[int]int{}
		cs[i] = map[int]int{}
		for j := 0; j < n; j++ {
			rs[i][j] = rs[i][j-1] + matrix[i][j]
			cs[i][j] = cs[i-1][j] + matrix[i][j]
		}
	}
	//for i := -1; i < n; i++ {
	//	fmt.Println(i, rs[i])
	//}
	//for i := -1; i < n; i++ {
	//	fmt.Println(i, cs[i])
	//}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				continue
			}
			for l := n - j; l >= ans[2]+1; l-- {
				//fmt.Println(i, j, l)
				//fmt.Println(rs[i][j+l-1], rs[i][j-1], rs[i+l-1][j+l-1], rs[i+l-1][j-1])
				//fmt.Println(cs[i+l-1][j], cs[i-1][j], cs[i+l-1][j+l-1], cs[i-1][j+l-1])
				if rs[i][j+l-1]-rs[i][j-1] == 0 &&
					rs[i+l-1][j+l-1]-rs[i+l-1][j-1] == 0 &&
					cs[i+l-1][j]-cs[i-1][j] == 0 &&
					cs[i+l-1][j+l-1]-cs[i-1][j+l-1] == 0 {

					if l > ans[2] {
						ans = []int{i, j, l}
						break
					}
					//fmt.Println(ans)
				}
			}
		}

	}

	if ans[2] == 0 {
		ans = []int{}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

package main

import "fmt"

// determine-whether-matrix-can-be-obtained-by-rotation 判断矩阵经轮转后是否一致  2022-03-28 20:37:59
func main() {
	fmt.Println(determineWhetherMatrixCanBeObtainedByRotation())
}

//leetcode submit region begin(Prohibit modification and deletion)
func findRotation(m [][]int, t [][]int) bool {
	n := len(m)

	chk := func(a [][]int) bool {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if a[i][j] != t[i][j] {
					return false
				}
			}
		}
		return true
	}

	do := func(a [][]int) [][]int {
		ans := make([][]int, n)
		for i := 0; i < n; i++ {
			ans[i] = make([]int, n)
			for j := 0; j < n; j++ {
				ans[i][j] = m[n-j][i]
			}
		}
		return ans
	}

	for i := 0; i < 4; i++ {
		if chk(m) {
			return true
		}
		m = do(m)
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)

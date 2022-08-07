package main

func main() {
}

//leetcode submit region begin(Prohibit modification and deletion)
func solveSudoku(board [][]byte) {
	var dfs func(t [][]byte, i, j int) bool
	dfs = func(t [][]byte, i, j int) bool {
		for x := i; x < 9; x++ {
			for y := j; y < 9; y++ {
				if board[x][y] == '.' {

				}
			}
		}
		return false
	}

	dfs(board, 0, 0)
}
//leetcode submit region end(Prohibit modification and deletion)

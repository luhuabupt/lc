package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func tictactoe(board []string) string {
	n := len(board)

	check := func(b int32) bool {
		flag := true

		for _, s := range board {
			flag = true
			for _, v := range s {
				if v != b {
					flag = false
					break
				}
			}
			if flag {
				return true
			}
		}

		for i := 0; i < n; i++ {
			flag = true
			for j := 0; j < n; j++ {
				v := int32(board[j][i])
				if v != b {
					flag = false
					break
				}
			}
			if flag {
				return true
			}
		}

		flag = true
		for i := 0; i < n; i++ {
			if int32(board[i][i]) != b {
				flag = false
				break
			}
		}
		if flag {
			return true
		}

		flag = true
		for i := 0; i < n; i++ {
			if int32(board[i][n-1-i]) != b {
				flag = false
				break
			}
		}
		if flag {
			return true
		}

		return false
	}

	if check('X') {
		return "X"
	}
	if check('O') {
		return "O"
	}

	for _, s := range board {
		for _, v := range s {
			if v == ' ' {
				return "Pending"
			}
		}
	}

	return "Draw"
}

//leetcode submit region end(Prohibit modification and deletion)

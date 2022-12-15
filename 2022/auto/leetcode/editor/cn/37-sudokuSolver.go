package main

import "fmt"

func main() {
	solveSudoku([][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}})
}

//leetcode submit region begin(Prohibit modification and deletion)
func solveSudoku(br [][]byte) {
	for _, v := range br {
		fmt.Println(string(v))
	}
	fmt.Println(" ")

	r := [9][9]bool{}
	c := [9][9]bool{}
	b := [9][9]bool{}
	for i, ar := range br {
		for j, v := range ar {
			if v == '.' {
				continue
			}
			x := v - '1'
			r[i][x] = true
			c[j][x] = true
			b[(i/3)*3+j/3][x] = true
		}
	}

	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i == 9 {
			return true
		}

		ni, nj := i, j+1
		if nj == 9 {
			ni, nj = ni+1, 0
		}

		if br[i][j] == '.' {
			for v := 0; v < 9; v++ {
				if !r[i][v] && !c[j][v] && !b[(i/3)*3+j/3][v] {
					br[i][j] = byte(v + '1')
					r[i][v] = true
					c[j][v] = true
					b[(i/3)*3+j/3][v] = true

					if dfs(ni, nj) {
						return true
					}

					br[i][j] = '.'
					r[i][v] = false
					c[j][v] = false
					b[(i/3)*3+j/3][v] = false
				}
			}
			return false
		} else {
			return dfs(ni, nj)
		}
	}

	dfs(0, 0)
	for _, v := range br {
		fmt.Println(string(v))
	}
}

//leetcode submit region end(Prohibit modification and deletion)

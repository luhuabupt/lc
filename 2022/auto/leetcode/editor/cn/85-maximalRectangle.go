package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func maximalRectangle(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	r := make([][]int, m)
	for i := 0; i < m; i++ {
		r[i] = make([]int, n)
		x := n
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				r[i][j] = x
			} else {
				x = j
			}
		}
	}

	ans := 0
	for i, arr := range matrix {
		for j, v := range arr {
			if v == '0' {
				continue
			}
			w := n
			for x := i; x < m && matrix[x][j] == '1'; x++ {
				if r[x][j]-j < w {
					w = r[x][j] - j
				}
				s := w * (x - i + 1)
				if s > ans {
					ans = s
				}
			}
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

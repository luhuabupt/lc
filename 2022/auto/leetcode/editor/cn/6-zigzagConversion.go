package main

import "fmt"

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}

//leetcode submit region begin(Prohibit modification and deletion)
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	ans, idx, up := make([][]byte, numRows), 0, 1
	for i, v := range s {
		ans[idx] = append(ans[idx], byte(v))
		if i != 0 && (idx == 0 || idx == numRows-1) {
			up = -up
		}
		idx += up
	}
	for i := 1; i < len(ans); i++ {
		ans[0] = append(ans[0], ans[i]...)
	}

	return string(ans[0])
}

//leetcode submit region end(Prohibit modification and deletion)

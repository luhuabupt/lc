package main

import "strings"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func maxRepeating(s string, w string) int {
	n := len(s)
	m := len(w)
	nw := ""
	ans := 0
	for i := 1; i*m <= n; i++ {
		nw += w
		if !strings.Contains(s, nw) {
			break
		}
		ans = i
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

package main

import (
	"fmt"
	"strconv"
)

// add-strings 字符串相加  2022-10-26 22:27:26
func main() {
	fmt.Println("")
}

//leetcode submit region begin(Prohibit modification and deletion)
func addStrings(a string, b string) string {
	ans := []int{}
	if len(a) < len(b) {
		return addStrings(b, a)
	}
	p := 0
	for i, j := len(a)-1, len(b)-1; j >= 0; j-- {
		p += int(a[i] - '0' + b[i] - '0')
		ans = append(ans, p%10)
		p /= 10
	}
	for i := len(a) - len(b); i >= 0; i-- {
		p += int(a[i] - '0')
		ans = append(ans, p%10)
		p /= 10
	}
	if p > 0 {
		ans = append(ans, p)
	}
	as := []byte{}
	for i := len(ans) - 1; i >= 0; i-- {
		as = append(as, strconv.Itoa(ans[i])[0])
	}
	return string(as)
}

//leetcode submit region end(Prohibit modification and deletion)

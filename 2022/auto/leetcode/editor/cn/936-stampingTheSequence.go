package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(movesToStamp("abc", "ababc"))
	fmt.Println(movesToStamp("abca", "aabcaca"))
	fmt.Println(movesToStamp("abc", "aaabcc"))
	fmt.Println(movesToStamp("abc", "aaabcaabc"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func movesToStamp(ss string, tt string) []int {
	ans := []int{}
	if !strings.Contains(tt, ss) {
		return ans
	}

	s, t := []byte(ss), []byte(tt)
	ls, lt := len(s), len(t)

	final := []int{}
	for i := 0; i <= lt-ls; i++ {
		if ss == tt[i:i+ls] {
			final = append(final, i)
		}
	}

	return append(ans, final...)
}

//leetcode submit region end(Prohibit modification and deletion)

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countValidWords("a  a-b-c   bc  d e"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func countValidWords(sentence string) (ans int) {
	arr := strings.Split(sentence, " ")

loop:
	for _, v := range arr {
		if len(v) == 0 || strings.Count(v, "-") > 1 {
			continue
		}
		for i, x := range v {
			if x >= '0' && x <= '9' {
				continue loop
			}
			if x == '-' {
				if i == 0 || i == len(v)-1 || (v[i-1] < 'a' || v[i-1] > 'z' || v[i+1] < 'a' || v[i+1] > 'z') {
					continue loop
				}
			}
			if x == '!' || x == '.' || x == ',' {
				if i != len(v)-1 {
					continue loop
				}
			}
		}
		ans++
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

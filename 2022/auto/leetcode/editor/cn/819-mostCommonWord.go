package main

import (
	"fmt"
	"strings"
)

// most-common-word 最常见的单词  2022-02-16 23:46:51
func main() {
	fmt.Println(mostCommonWord())
}

//leetcode submit region begin(Prohibit modification and deletion)
func mostCommonWord(paragraph string, banned []string) (ans string) {
	paragraph = strings.ToLower(paragraph + " ")
	b := map[string]bool{}
	for _, v := range banned {
		b[v] = true
	}
	cnt := map[string]int{}

	ch := "!?',;. "
	m := map[int32]bool{}
	for _, v := range ch {
		m[v] = true
	}

	pre := []byte{}
	for _, v := range paragraph {
		if m[v] {
			if len(pre) > 0 {
				cnt[string(pre)]++
				pre = []byte{}
			}
		} else {
			pre = append(pre, byte(v))
		}
	}

	for k, v := range cnt {
		if !b[k] && v > cnt[ans] {
			ans = k
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

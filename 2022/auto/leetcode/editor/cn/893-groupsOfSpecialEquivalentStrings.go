package main

import (
	"fmt"
)

func main() {
	fmt.Println(numSpecialEquivGroups([]string{"abcd", "cdab", "cbad", "xyzz", "zzxy", "zzyx"}))
	fmt.Println(numSpecialEquivGroups([]string{"abc", "acb", "bac", "bca", "cab", "cba"}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func numSpecialEquivGroups(w []string) int {
	cnt := map[string]bool{}
	for _, s := range w {
		c := [52]int{}
		for i, v := range s {
			c[26*(i%2)+int(v-'a')]++
		}
		x := []byte{}
		for _, v := range c {
			x = append(x, byte(v))
		}
		cnt[string(x)] = true
	}
	return len(cnt)
}

//leetcode submit region end(Prohibit modification and deletion)

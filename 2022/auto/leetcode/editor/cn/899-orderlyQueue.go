package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(orderlyQueue("baaca", 3))
}

//leetcode submit region begin(Prohibit modification and deletion)
func orderlyQueue(s string, k int) string {
	if k == 1 {
		x := []string{}
		mi := 'z'
		for _, v := range s {
			if v < mi {
				mi = v
			}
		}
		for i, v := range s {
			if v == mi {
				x = append(x, s[i:]+s[:i])
			}
		}
		sort.Strings(x)
		return x[0]
	}
	sa := []byte(s)

	sort.Slice(sa, func(i, j int) bool {
		return sa[i] < sa[j]
	})
	return string(sa)
}

//leetcode submit region end(Prohibit modification and deletion)

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(" ")
}

//leetcode submit region begin(Prohibit modification and deletion)
func alertNames(a []string, b []string) []string {
	c := [][]string{}
	for i, v := range a {
		c = append(c, []string{v, b[i]})
	}
	sort.Slice(c, func(i, j int) bool {
		return c[i][0] < c[j][0] || c[i][0] == c[j][0] && c[i][1] < c[j][1]
	})

	chk := func(a, b string) bool {
		if a[:2] == b[:2] {
			return true
		}

		ai, _ := strconv.Atoi(a[:2])
		bi, _ := strconv.Atoi(b[:2])

		return ai+1 == bi && a[3:] > b[3:]
	}

	ans := []string{}
	for i := 2; i < len(c); i++ {
		if len(ans) > 0 && ans[len(ans)-1] == c[i][0] {
			continue
		}
		if c[i][0] == c[i-1][0] && c[i][0] == c[i-2][0] {
			if chk(c[i][1], c[i-2][1]) {
				ans = append(ans, c[i][0])
			}
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

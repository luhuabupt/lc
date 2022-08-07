package main

import (
	"strconv"
	"strings"
)

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func subdomainVisits(a []string) []string {
	m := map[string]int{}
	for _, s := range a {
		x := strings.Split(s, " ")
		c, _ := strconv.Atoi(x[0])
		m[x[1]] += c
		for k, v := range x[1] {
			if v == '.' {
				m[x[1][k+1:]] += c
			}
		}
	}
	ans := []string{}
	for k, v := range m {
		ans = append(ans, strconv.Itoa(v)+" "+k)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

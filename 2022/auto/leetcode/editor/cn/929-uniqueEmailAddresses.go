package main

import "strings"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func numUniqueEmails(s []string) int {
	m := map[string]map[string]bool{}
	for _, v := range s {
		va := strings.Split(v, "@")
		na := strings.Split(va[0], "+")
		na[0] = strings.ReplaceAll(na[0], ".", "")

		if m[va[1]] == nil {
			m[va[1]] = map[string]bool{}
		}
		m[va[1]][na[0]] = true
	}

	ans := 0
	for _, x := range m {
		ans += len(x)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

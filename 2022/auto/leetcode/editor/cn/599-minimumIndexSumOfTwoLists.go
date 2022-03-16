package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func findRestaurant(a []string, b []string) (ans []string) {
	m := map[string]int{}
	for i, v := range b {
		m[v] = i
	}

	mi := 1 << 30
	for i, v := range a {
		if _, ok := m[v]; ok {
			if i+m[v] == mi {
				ans = append(ans, v)
			}
			if i+m[v] < mi {
				mi = i + m[v]
				ans = []string{v}
			}
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

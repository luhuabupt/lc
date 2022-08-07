package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func findAndReplacePattern(w []string, p string) []string {
	ans := []string{}

loop:
	for _, s := range w {
		h := [27]int{}
		rh := [27]int{}
		for i, v := range s {
			if h[v-'a'+1] != 0 && h[v-'a'+1] != int(p[i]-'a'+1) {
				continue loop
			}
			if rh[int(p[i]-'a'+1)] != 0 && rh[int(p[i]-'a'+1)] != int(v-'a'+1) {
				continue loop
			}
			h[v-'a'+1] = int(p[i] - 'a' + 1)
			rh[int(p[i]-'a'+1)] = int(v - 'a' + 1)
		}
		ans = append(ans, s)
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

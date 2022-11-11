package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func equalFrequency(w string) bool {
	c := [26]int{}
	for _, v := range w {
		c[v-'a']++
	}

	chk := func(i int) bool {
		b := c
		b[i]--
		t := 0
		for _, v := range b {
			if v == 0 {
				continue
			}
			if t == 0 {
				t = v
			}
			if t != v {
				return false
			}
		}
		return true
	}

	for i := 0; i < 26; i++ {
		if chk(i) {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

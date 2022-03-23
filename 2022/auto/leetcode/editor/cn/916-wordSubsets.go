package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func wordSubsets(w1 []string, w2 []string) []string {
	do := [26]int{}
	for _, v := range w2 {
		c := [26]int{}
		for _, x := range v {
			c[x-'a']++
		}

		for i, x := range c {
			if do[i] < x {
				do[i] = x
			}
		}
	}

	ans := []string{}

loop:
	for _, v := range w1 {
		c := [26]int{}
		for _, x := range v {
			c[x-'a']++
		}

		for i, x := range c {
			if do[i] > x {
				continue loop
			}
		}
		ans = append(ans, v)
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

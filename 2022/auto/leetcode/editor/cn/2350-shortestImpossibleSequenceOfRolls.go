package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func shortestSequence(rolls []int, k int) int {
	c := make([]bool, k)
	cc := 0
	ans := 0
	for _, v := range rolls {
		if c[v-1] == false {
			cc++
		}
		c[v-1] = true

		if cc == k {
			ans++
			c = make([]bool, k)
			cc = 0
		}
	}
	return ans + 1
}

//leetcode submit region end(Prohibit modification and deletion)

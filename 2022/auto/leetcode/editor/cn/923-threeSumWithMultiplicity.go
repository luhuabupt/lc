package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func threeSumMulti(a []int, t int) int {
	M := int(1e9 + 7)
	d := make([]int, 101)
	s := make([]int, 101)
	for _, v := range a {
		s[v]++
	}

	ans := 0
	d[a[0]]++
	for i := 1; i < len(a)-1; i++ {
		v := a[i]
		for l := 0; l <= t-v && t-v >= 0 && l <= 100; l++ {
			if t-v-l > 100 {
				continue
			}

			r := s[t-v-l] - d[t-v-l]
			if t-v-l == a[i] {
				r--
			}

			ans += d[l] * r
			ans %= M
		}
		d[v]++
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

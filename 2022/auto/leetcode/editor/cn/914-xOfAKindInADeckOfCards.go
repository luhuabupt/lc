package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func hasGroupsSizeX(a []int) bool {
	m := map[int]int{}
	for _, v := range a {
		m[v]++
	}

	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	c := -1
	for _, x := range m {
		if c == -1 {
			c = x
		} else {
			c = gcd(c, x)
		}
		if c == 1 {
			return false
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)

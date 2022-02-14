package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func rob(a []int) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	p, pp := 0, 0
	for _, v := range a {
		t := p
		p = max(p, pp+v)
		pp = t
	}

	return max(pp, p)
}

//leetcode submit region end(Prohibit modification and deletion)

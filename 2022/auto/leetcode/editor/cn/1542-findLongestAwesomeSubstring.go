package main

import "fmt"

func main() {
	fmt.Println(longestAwesome("213123"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func longestAwesome(s string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	d := [1024]int{}
	for i, _ := range d {
		d[i] = 1e10
	}
	d[0] = -1

	t := 0
	ans := 0
	for i, v := range s {
		t ^= 1 << (v - '0')
		ans = max(ans, i-d[t])

		for j := 0; j < 10; j++ {
			x := t ^ 1<<j
			ans = max(ans, i-d[x])
		}

		if d[t] == 1e10 {
			d[t] = i
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

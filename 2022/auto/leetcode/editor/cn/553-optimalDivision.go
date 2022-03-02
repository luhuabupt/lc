package main

import "fmt"

// optimal-division 最优除法  2022-02-27 14:25:00
func main() {
	fmt.Println(optimalDivision())
}

//leetcode submit region begin(Prohibit modification and deletion)
func optimalDivision(nums []int) string {
	n := len(nums)

	mx := 0.0
	nl := make([]bool, n)
	nr := make([]bool, n)

	var dfs func(i, un int, l, r []bool)
	dfs = func(i, un int, l, r []bool) {
		if i == n {
			for j := 0; j < n; j++ {
				if l[j] 
			}
			return
		}
		dfs(i+1, un, l, r)

		for x := un; un >= 0; x-- {
			nl = append([]bool{}, l...)
			nr = append([]bool{}, r...)
			nr[i] = true
			dfs(i+1, un-1, nl, nr)
		}

		nl = append([]bool{}, l...)
		nr = append([]bool{}, r...)
		nl[i] = true
		dfs(i+1, un+1, nl, nr)
	}
}

//leetcode submit region end(Prohibit modification and deletion)

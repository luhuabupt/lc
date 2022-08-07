package main

import "fmt"

func main() {
	fmt.Println(maxSubarraySumCircular([]int{5, -4, 5}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxSubarraySumCircular(a []int) int {
	ans := -1 << 63

	sum := 0
	mi := 0
	ma := -1 << 40
	miSum := 0
	for _, v := range a {
		sum += v
		if sum-mi > ans {
			ans = sum - mi
		}
		if sum-ma < miSum {
			miSum = sum - ma
		}

		if sum < mi {
			mi = sum
		}
		if sum > ma {
			ma = sum
		}
	}

	if sum-miSum > ans {
		ans = sum - miSum
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

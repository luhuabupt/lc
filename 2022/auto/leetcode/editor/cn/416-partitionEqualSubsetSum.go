package main

import "fmt"

func main() {
	fmt.Println(canPartition([]int{1, 1, 1, 1}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func canPartition(nums []int) bool {
	s, mx := 0, 0
	for _, v := range nums {
		s += v
		if v > mx {
			mx = v
		}
	}
	if s&1 == 1 || mx > s/2 {
		return false
	}

	dp := make([]bool, s/2+1)
	dp[0] = true
	for _, v := range nums {
		if dp[s/2-v] {
			return true
		}
		for i := len(dp) - 1; i >= 0; i-- {
			if dp[i] && v+i < s/2 {
				dp[v+i] = true
			}
		}
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)

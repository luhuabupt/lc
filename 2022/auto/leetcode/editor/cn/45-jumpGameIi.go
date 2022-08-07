package main

import "fmt"

func main() {
	fmt.Println(jump([]int{1, 2, 3}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 0

	st, idx := []int{0}, 0
	for {
		cur := st[idx]
		idx++

		if cur == n-1 {
			return dp[cur]
		}
		for i := 1; cur+i < n && i <= nums[cur]; i++ {
			if dp[cur+i] == 0 {
				st = append(st, cur+i)
				dp[cur+i] = dp[cur] + 1
			}
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

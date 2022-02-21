package main

import "fmt"

func main() {
	fmt.Println(numSubarrayBoundedMax([]int{2, 1, 4, 3}, 2, 3))
}

//leetcode submit region begin(Prohibit modification and deletion)
func numSubarrayBoundedMax(nums []int, left int, right int) int {
	nums = append(nums, right+1)
	cnt := func(t int) (ans int) {
		tmp := 0
		for _, x := range nums {
			if x > t {
				ans += tmp * (tmp + 1) / 2
				tmp = 0
			} else {
				tmp++
			}
		}
		return
	}

	return cnt(right) - cnt(left-1)
}

//leetcode submit region end(Prohibit modification and deletion)

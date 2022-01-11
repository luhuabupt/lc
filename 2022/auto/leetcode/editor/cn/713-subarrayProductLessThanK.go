package main

import "fmt"

// subarray-product-less-than-k 乘积小于K的子数组  2022-01-11 00:02:26
func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{1, 2, 3}, 2))
}

//leetcode submit region begin(Prohibit modification and deletion)
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	ans, sum := 0, 1
	for l, r := 0, 0; r < len(nums); r++ {
		sum *= nums[r]
		for sum >= k {
			sum /= nums[l]
			l++
		}
		ans += r - l + 1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

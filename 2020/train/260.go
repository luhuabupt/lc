package main

func maximumDifference(nums []int) int {
	n := len(nums)
	ans := -1
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[j]-nums[i] > ans {
				ans = nums[j] - nums[i]
			}
		}
	}
	return ans
}

package main

func maximumDifference(nums []int) int {
	n := len(nums)
	ans := -1
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[j]-nums[i] > 0 && nums[j]-nums[i] > ans {
				ans = nums[j] - nums[i]
			}
		}
	}
	return ans
}

func gridGame(grid [][]int) int64 {
	n := len(grid[0])
	one, two := make([]int64, n+1), make([]int64, n)

	one[n] = 0
	for i := n - 1; i >= 0; i-- {
		one[i] = one[i+1] + int64(grid[0][i])
	}

	two[0] = int64(grid[1][0])
	for i := 1; i < n; i++ {
		two[i] = two[i-1] + int64(grid[1][i])
	}

	ans := int64(-1)
	for i := 0; i < n; i++ {
		tmp := one[i+1]
		if i > 0 && two[i-1] > tmp {
			tmp = two[i-1]
		}
		if ans == -1 || tmp < ans {
			ans = tmp
		}
	}

	return ans
}

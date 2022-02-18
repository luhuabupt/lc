package main

import "fmt"

func main() {
	fmt.Println(bestRotation([]int{2, 3, 1, 4, 0}))
	fmt.Println(bestRotation([]int{1, 3, 0, 2, 4}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func bestRotation(nums []int) int {
	n := len(nums)

	score := 0
	for i, v := range nums {
		if v <= i {
			score++
		}
	}

	add := make([]int, n+1)
	for i, v := range nums {
		if v == n {
			continue
		}

		// 出区间: [v,n-1]
		if i >= v {
			add[i-v+1]--
		}
		// 右边入队
		add[i+1]++
		// 移出区间
		if i+1+n-v < n {
			add[i+1+n-v]--
		}
	}

	mx, ans := score, 0
	for i := 1; i < n; i++ {
		score += add[i]
		if score > mx {
			mx, ans = score, i
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

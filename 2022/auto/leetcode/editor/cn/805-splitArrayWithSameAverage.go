package main

import "fmt"

func main() {
	fmt.Println(splitArraySameAverage([]int{1, 3}))
	fmt.Println(splitArraySameAverage([]int{1, 4, 5, 8, 2, 3, 6, 7}))
	fmt.Println(splitArraySameAverage([]int{18, 0, 16, 2})) //1
}

//leetcode submit region begin(Prohibit modification and deletion)
func splitArraySameAverage(nums []int) bool {
	n := len(nums)
	a, b := nums[:(n+1)/2], nums[(n+1)/2:]

	sum := 0
	for _, v := range nums {
		sum += v
	}

	h := map[int]bool{}
	for i := 0; i < (1<<len(b) - 1); i++ {
		x, cnt := 0, 0
		for j := 0; j < len(b); j++ {
			if i>>j&1 == 1 {
				x += b[j]
				cnt++
			}
		}
		xx := n*x - sum*cnt
		h[xx] = true
	}

	for i := 1; i < 1<<len(a); i++ {
		x, cnt := 0, 0
		for j := 0; j < len(a); j++ {
			if i>>j&1 == 1 {
				x += a[j]
				cnt++
			}
		}
		xx := n*x - sum*cnt
		if h[-xx] {
			return true
		}
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)

package main

import "fmt"

func main() {
	fmt.Println(numberOfGoodSubsets([]int{2, 3}))
	fmt.Println(numberOfGoodSubsets([]int{1, 2, 3, 4}))
	fmt.Println(numberOfGoodSubsets([]int{2, 3, 15}))
	fmt.Println(numberOfGoodSubsets([]int{1, 2, 3, 5, 15}))
	fmt.Println(numberOfGoodSubsets([]int{18, 28, 2, 17, 29, 30, 15, 9, 12})) // 19
}

//leetcode submit region begin(Prohibit modification and deletion)
func numberOfGoodSubsets(nums []int) int {
	na := [31]int{0, 0, 1, 2, 0, 4, 3, 8, 0, 0,
		5, 16, 0, 32, 9, 6, 0, 64, 0, 128,
		0, 10, 17, 256, 0, 0, 33, 0, 0, 512, 7}

	one, M, ans := 1, int(1e9+7), 0
	cnt := [31]int{}
	for _, v := range nums {
		if v == 1 {
			one = (one << 1) % M
			continue
		}
		if na[v] > 0 {
			cnt[v]++
		}
	}

	f := make([]int, 1024)
	f[0] = 1
	for k, v := range na {
		if v == 0 {
			continue
		}
		for i := 1023; i > 0; i-- {
			if i|v == i {
				f[i] += f[i^v] * cnt[k]
				f[i] %= M
			}
		}
	}

	for i := 1; i <= 1023; i++ {
		ans += f[i]
		ans %= M
	}

	ans *= one
	ans %= M

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

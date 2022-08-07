package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func maxRotateFunction(nums []int) int {
	n := len(nums)
	s := 0
	t := 0
	for i := 0; i < n; i++ {
		s += nums[i]
		t += nums[i] * i
	}

	ans := t
	for i := 1; i < n; i++ {
		t += s
		t -= nums[n-i] * n
		if t > ans {
			ans = t
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

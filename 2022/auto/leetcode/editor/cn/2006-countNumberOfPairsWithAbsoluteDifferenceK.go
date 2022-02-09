package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func countKDifference(nums []int, k int) (ans int) {
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]-v == k || v-nums[j] == k {
				ans++
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func consecutiveNumbersSum(n int) int {
	ans := 0
	for i := 1; (1+i)*i/2 <= n; i++ {
		if i%2 == 1 {
			if n%i == 0 {
				ans++
			}
		} else {
			if 2*n%i == 0 && 2*n/i%2 == 1 {
				ans++
			}
		}

	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

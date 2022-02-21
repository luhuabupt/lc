package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func numDecodings(s string) int {
	ans := 1
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			return 0
		} else if s[i] == '1' {

		} else if s[i] == '2' {

		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

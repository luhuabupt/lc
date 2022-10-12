package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func breakPalindrome(sa string) string {
	s := []byte(sa)
	n := len(s)
	if n == 1 {
		return ""
	}

	for i := 0; i < n/2; i++ {
		if s[i] != 'a' {
			s[i] = 'a'
			return string(s)
		}
	}

	s[n-1] = 'b'
	return string(s)
}

//leetcode submit region end(Prohibit modification and deletion)

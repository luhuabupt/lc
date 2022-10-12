package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func isFlipedString(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	if a == b {
		return true
	}
	n := len(a)
	for i := 0; i < n; i++ {
		if a[:i+1] == b[n-1-i:] && a[i+1:] == b[:n-1-i] {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

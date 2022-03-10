package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func decodeAtIndex(s string, k int) string {
	pre := 0
	for i, v := range s {
		if v >= 'a' && v <= 'z' {
			pre++
			if pre == k {
				return string(v)
			}
		} else {
			x := int(v - '0')
			if x*pre >= k {
				if k%pre == 0 {
					return decodeAtIndex(s[:i], pre)
				}
				return decodeAtIndex(s[:i], k%pre)
			}
			pre *= x
		}
	}
	return ""
}

//leetcode submit region end(Prohibit modification and deletion)

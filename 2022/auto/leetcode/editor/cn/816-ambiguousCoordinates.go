package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func ambiguousCoordinates(s string) []string {
	n := len(s)
	s = s[1 : n-1]
	n = len(s)

	cal := func(i, j int) []string {
		if s[i] == '0' {
			if i == j {
				return []string{"0"}
			}
			if s[j] == '0' {
				return []string{}
			}
			return []string{s[i:i+1] + "." + s[i+1:j+1]}
		}
		res := []string{s[i : j+1]}

		if s[j] != '0' {
			for x := i; x < j; x++ {
				res = append(res, s[i:x+1]+"."+s[x+1:j+1])
			}
		}

		return res
	}

	ans := []string{}
	for i := 0; i < n-1; i++ {
		l, r := cal(0, i), cal(i+1, n-1)
		if len(l) > 0 && len(r) > 0 {
			for _, v := range l {
				for _, x := range r {
					ans = append(ans, "("+v+", "+x+")")
				}
			}
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

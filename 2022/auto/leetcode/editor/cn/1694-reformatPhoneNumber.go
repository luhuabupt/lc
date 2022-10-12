package main

// reformat-phone-number 重新格式化电话号码  2022-10-01 19:26:25
func main() {
}

//leetcode submit region begin(Prohibit modification and deletion)
func reformatNumber(a string) string {
	ar := []byte{}
	for _, v := range a {
		if v == ' ' || v == '-' {
			continue
		}
		ar = append(ar, byte(v))
	}
	ans := []byte{}
	n := len(ar)
	nn := n
	if n%3 == 1 {
		n -= 4
	}
	if n%3 == 2 {
		n -= 2
	}
	for i := 0; i < n; i++ {
		v := ar[i]
		ans = append(ans, v)
		if i%3 == 0 && (i != 0 && i != n) {
			ans = append(ans, '-')
		}
	}
	for j := n; j < nn; j++ {
		v := ar[j]
		ans = append(ans, v)
		ans = append(ans, '-')
	}

	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)

package main

import "fmt"

func main() {
	//fmt.Println(strongPasswordChecker("aaa111"))
	//fmt.Println(strongPasswordChecker("aaBBB"))
	//fmt.Println(strongPasswordChecker("aaaB1"))
	//fmt.Println(strongPasswordChecker("aaa11"))
	//fmt.Println(strongPasswordChecker("aaabb"))
	fmt.Println(strongPasswordChecker("bbaaaaaaaaaaaaaaacccccc")) // 8
}

//leetcode submit region begin(Prohibit modification and deletion)
func strongPasswordChecker(psd string) int {
	n := len(psd)
	fmt.Println(n)

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	a, b, c := 0, 0, 0
	for _, v := range psd {
		if v >= 'a' && v <= 'z' {
			a = 1
		}
		if v >= 'A' && v <= 'Z' {
			b = 1
		}
		if v >= '0' && v <= '9' {
			c = 1
		}
	}

	diff := 0
	if n <= 4 {
		return 6 - n
	}
	if n == 5 {
		all := true
		for i := 1; i < n; i++ {
			if psd[i] != psd[0] {
				all = false
				break
			}
		}
		if all {
			return 2
		}
		return 1 + max(2-(a+b+c), 0)
	}
	if n <= 20 {
		for i := 2; i < n; i++ {
			if psd[i] == psd[i-1] && psd[i] == psd[i-2] {
				diff++
				i += 2
			}
		}
		if n <= 20 {
			if diff >= 2 {
				return diff
			}
			if diff == 1 {
				if a+b+c >= 2 {
					return diff
				}
				return 2
			}
			if diff == 0 {
				return 3 - (a + b + c)
			}
		}
	}

	del := 0
	for i := 2; i < n; i++ {
		if n-del <= 20 {
			if psd[i] == psd[i-1] && psd[i] == psd[i-2] {
				diff++
				i += 2
			}
		} else {
			if psd[i] == psd[i-1] && psd[i] == psd[i-2] {
				del++
			}
		}
	}
	if a+b+c == 3 {
		return diff + del + max(n-del-20, 0)
	}
	if a+b+c == 2 {
		if diff > 0 {
			return diff + del + max(n-del-20, 0)
		} else {
			return diff + del + max(n-del-20, 0) + 1
		}
	}
	return diff + del + max(n-del-20, 0) + max(3-(a+b+c)-diff, 0)
}

//leetcode submit region end(Prohibit modification and deletion)

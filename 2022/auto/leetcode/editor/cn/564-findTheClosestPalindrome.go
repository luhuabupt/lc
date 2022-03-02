package main

import (
	"strconv"
)

func main() {
	//fmt.Println(nearestPalindromic("123"))   // 121
	//fmt.Println(nearestPalindromic("1"))     // 0
	//fmt.Println(nearestPalindromic("1234"))  //  1221
	//fmt.Println(nearestPalindromic("121"))   // 111
	//fmt.Println(nearestPalindromic("10")) // 9
	//fmt.Println(nearestPalindromic("1221"))  // 1111
	//fmt.Println(nearestPalindromic("99"))    // 101
	//fmt.Println(nearestPalindromic("10001")) // 9999
	//fmt.Println(nearestPalindromic("1283"))  // 1331
	//fmt.Println(nearestPalindromic("1200"))  // 1221
	//fmt.Println(nearestPalindromic("11"))  // 9
	//fmt.Println(nearestPalindromic("999"))  // 1001
	//fmt.Println(nearestPalindromic("1000"))  // 999
	//fmt.Println(nearestPalindromic("100"))  // 99
}

//leetcode submit region begin(Prohibit modification and deletion)
func nearestPalindromic(s string) string {
	n := len(s)
	si, _ := strconv.Atoi(s)

	base, p := 0, 1
	for x := 0; x < (n+1)/2; x++ {
		base *= 10
		base += int(s[x] - '0')
		p *= 10
	}
	//fmt.Println(base, p)

	odd := n%2 == 0
	a := []int{0}

	if s == re(s) {
		if base == 1 {
			a = append(a, 9)
		} else if base == p/10 { // 10
			if odd {
				a = append(a, gt(p-1, !odd))
			} else {
				a = append(a, gt(p/10-1, !odd))
			}
		} else {
			a = append(a, gt(base-1, odd))
		}
		if base == p-1 { // 99
			a = append(a, gt(p/10, !odd))
		} else {
			a = append(a, gt(base+1, odd))
		}
	}

	x := gt(base, odd)
	a = append(a, x)
	if x > si {
		if base == 1 {
			a = append(a, 9)
		} else if base == p/10 { // 10
			if odd {
				a = append(a, gt(p-1, !odd))
			} else {
				a = append(a, gt(p/10-1, !odd))
			}
		} else {
			a = append(a, gt(base-1, odd))
		}
	} else {
		if base == p-1 { // 99
			a = append(a, gt(p, !odd))
		} else {
			a = append(a, gt(base+1, odd))
		}
	}

	//fmt.Println(a)
	ans := 0
	for _, v := range a {
		if v != si && abs(si, v) < abs(ans, si) || abs(si, v) == abs(ans, si) && v < ans {
			ans = v
		}
	}

	return strconv.Itoa(ans)
}

func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func re(s string) string {
	a := []byte(s)
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-1-i] = a[len(a)-1-i], a[i]
	}
	return string(a)
}

func gt(i int, odd bool) int {
	a := []int{}
	for x := i; x > 0; x /= 10 {
		a = append([]int{x % 10}, a...)
	}
	k := len(a) - 1
	if !odd {
		k--
	}
	for x := k; x >= 0; x-- {
		a = append(a, a[x])
	}
	ans := 0
	for _, v := range a {
		ans *= 10
		ans += v
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

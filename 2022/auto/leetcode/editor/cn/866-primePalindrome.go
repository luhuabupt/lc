package main

import "fmt"

func main() {
	fmt.Println(primePalindrome(13))
	fmt.Println(primePalindrome(11))
}

//leetcode submit region begin(Prohibit modification and deletion)
func primePalindrome(n int) int {
	if n <= 7 {
		return []int{0, 1, 2, 3, 5, 5, 7, 7}[n]
	}
	if n < 11 {
		return 11
	}

	a := []int{}
	for x := n; x >= 0; x /= 10 {
		a = append(a, x%10)
	}
	if len(a)%2 == 0 {
		return pow(10, len(a)) + 1
	}

	for i := n / pow(10, len(a)/2); i < pow(10, len(a)/2+1)-1; i++ {
		x := get(i)
		if check(x) {
			return x
		}
	}
	return pow(10, len(a)+1) + 1
}

func get(i int) int {
	a := []int{}
	for x := i; x >= 0; x /= 10 {
		a = append(a, x%10)
	}
	for x := len(a) - 2; x >= 0; x-- {
		a = append(a, a[x])
	}
	ans := 0
	for _, v := range a {
		ans *= 10
		ans += v
	}

	return ans
}

func pow(i, v int) int {
	ans := 1
	for x := 0; x < v; x++ {
		ans *= i
	}
	return ans
}

func check(v int) bool {
	if v%6 != 1 {
		return false
	}
	for i := 2; i*i <= v; i++ {
		if v%i == 0 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

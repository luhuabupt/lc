package main

import "fmt"

func main() {
	fmt.Println(QuickPow(2, 50))
}

func QuickPow(a, n int) int {
	ans := 1
	for n > 0 {
		if n&1 > 0 {
			ans *= a
		}
		a *= a
		n /= 2
	}
	return ans
}

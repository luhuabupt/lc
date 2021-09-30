package lib

func main() {
	QuickPow(2, 5)
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

package lib

// 快速幂
func myPow(x float64, n int) float64 {
	ans := 1.0
	for n > 0 {
		if n%2 == 1 {
			ans *= x
			n--
		}
		x *= x
		n /= 2
	}
	return ans
}

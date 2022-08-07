package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func addToArrayForm(a []int, k int) []int {
	re := func(a []int) []int {
		n := len(a)
		for i := 0; i < n/2; i++ {
			a[i], a[n-1-i] = a[n-1-i], a[i]
		}
		return a
	}

	a = re(a)
	for i := 0; i < len(a); i++ {
		a[i] += k
		k = a[i] / 10
		a[i] %= 10
	}

	for k > 0 {
		a = append(a, k%10)
		k /= 10
	}

	return re(a)
}

//leetcode submit region end(Prohibit modification and deletion)

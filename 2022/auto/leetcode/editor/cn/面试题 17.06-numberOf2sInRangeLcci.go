package main

import "fmt"

func main() {
	fmt.Println(numberOf2sInRange(1234))
}

//leetcode submit region begin(Prohibit modification and deletion)
func numberOf2sInRange(n int) (ans int) {
	for p := 1; p < n; p *= 10 {
		x := n % (p * 10) / p

		if x > 2 {
			ans += (n/(p*10) + 1) * p
		} else if x == 2 {
			ans += (n/(p*10))*p + n%p + 1
		} else {
			ans += (n / (p * 10)) * p
		}

	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

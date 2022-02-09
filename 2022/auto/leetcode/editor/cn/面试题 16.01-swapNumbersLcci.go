package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func swapNumbers(a []int) []int {
	a[1] = a[0] + a[1]
	a[0] = a[1] - a[0]
	a[1] = a[1] - a[0]

	return a
}

//leetcode submit region end(Prohibit modification and deletion)

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func maxTurbulenceSize(a []int) int {
	n := len(a)
	l, r := 0, 0
	ans := 1
	for i := 1; i < n; i++ {
		if a[i] == a[i-1] {
			l, r = i, i
		}
		if a[i] > a[i-1] && i > 1 && a[i-1] > a[i-2] {
			l, r = i-1, i
		}
		if a[i] < a[i-1] && i > 1 && a[i-1] < a[i-2] {
			l, r = i-1, i
		}
		r = i
		if r-l+1 > ans {
			ans = r - l + 1
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

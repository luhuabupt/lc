package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func numsSameConsecDiff(n int, k int) []int {
	a, b := []int{}, []int{}
	for i := 1; i < 10; i++ {
		a = append(a, i)
	}
	for i := 1; i < n; i++ {
		b = []int{}
		for _, x := range a {
			if x%10 >= k {
				b = append(b, x%10-k+x*10)
			}
			if k != 0 && x%10+k < 10 {
				b = append(b, x%10+k+x*10)
			}
		}
		a = b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func powerfulIntegers(x int, y int, bound int) []int {
	a := make([]bool, bound+1)
	for i := 1; i <= bound; i *= x {
		for j := 1; i+j <= bound; j *= y {
			a[i+j] = true
			if y == 1 {
				break
			}
		}
		if x == 1 {
			break
		}
	}
	ans := []int{}
	for i, v := range a {
		if v {
			ans = append(ans, i)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func findSwapValues(array1 []int, array2 []int) []int {
	s1, s2 := 0, 0
	h1, h2 := map[int]bool{}, map[int]bool{}
	ans := []int{}

	for _, v := range array1 {
		s1 += v
		h1[v] = true
	}
	for _, v := range array2 {
		s2 += v
		h2[v] = true
	}

	for _, v := range array1 {
		if (s2-s1+2*v)%2 == 0 && h2[(s2-s1+2*v)/2] {
			ans = []int{v, (s2 - s1 + 2*v) / 2}
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

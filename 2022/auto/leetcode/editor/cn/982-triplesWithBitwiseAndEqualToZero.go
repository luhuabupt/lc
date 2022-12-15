package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func countTriplets(a []int) int {
	n := len(a)
	d := map[int]int{}
	for i, v := range a {
		for j := i; j < n; j++ {
			d[v&a[j]]++
		}
	}

	ans := 0
	for _, v := range a {
		for x, c := range d {
			if v&x == 0 {
				ans += c * 3
			}
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

package main

import "sort"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func carFleet(t int, p []int, s []int) int {
	n := len(p)
	w := make([][2]int, n)

	for i, v := range p {
		w = append(w, [2]int{v, s[i]})
	}

	sort.Slice(w, func(i, j int) bool {
		return w[i][0] > w[j][0]
	})

	ans, pre := n, n-1
	for i := 1; i < n; i++ {
		a, c := w[pre], w[i]
		if (t-a[0])*c[1] > (t-c[0])*a[1] {
			ans--
		} else {
			pre = i
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

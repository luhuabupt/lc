package main

import "sort"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func maxAliveYear(birth []int, death []int) (ans int) {
	cnt := map[int]int{}
	for _, v := range birth {
		cnt[v]++
	}
	for _, v := range death {
		cnt[v+1]--
	}

	arr := []int{}
	for k, _ := range cnt {
		arr = append(arr, k)
	}

	sort.Ints(arr)
	cur, max := 0, 0
	for _, y := range arr {
		cur += cnt[y]
		if cur > max {
			max, ans = cur, y
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

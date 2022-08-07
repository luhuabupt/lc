package main

import (
	"fmt"
)

func main() {
	//fmt.Println(subarrayBitwiseORs([]int{1, 2, 4}))
	//fmt.Println(subarrayBitwiseORs([]int{1, 1, 2}))
	fmt.Println(subarrayBitwiseORs([]int{15, 54, 12, 68, 33, 127, 82, 115, 14})) // 15
}

//leetcode submit region begin(Prohibit modification and deletion)
func subarrayBitwiseORs(arr []int) int {
	m := map[int]bool{}
	pre := map[int]bool{}
	next := map[int]bool{}

	for _, v := range arr {
		m[v] = true
		next = map[int]bool{v: true}
		for x, _ := range pre {
			next[x|v] = true
			m[x|v] = true
		}
		pre = next
	}

	return len(m)
}

//leetcode submit region end(Prohibit modification and deletion)

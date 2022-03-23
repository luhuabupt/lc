package main

import (
	"fmt"
)

func main() {

	fmt.Println(storeWater([]int{3, 2}, []int{20, 17}))

	fmt.Println(storeWater([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, []int{10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000, 10000}))
	//
	//fmt.Println(storeWater([]int{9988,5017,5130, 9309,8999,56}, []int{9991,6973,7192,9951,9034,8299}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func storeWater(bucket []int, vat []int) int {
	// 200次
	n := len(bucket)
	ma := 0
	for _, v := range vat {
		if v > ma {
			ma = v
		}
	}

	ans := 0
	for i := 1; i <= ma; i++ {
		cur := i
		for k := 0; k < n; k++ {
			per := (vat[k]-1)/i + 1
			if per > bucket[k] {
				cur += per - bucket[k]
			}
		}
		if ans == 0 || cur < ans {
			ans = cur
		}

		//fmt.Println(i, ma)
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

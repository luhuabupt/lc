package main

import "fmt"

func main() {
	//fmt.Println(minSwapsCouples([]int{5,4,2,6,3,1,0,7}))
	//fmt.Println(minSwapsCouples([]int{0,5,3,4,1,2}))
	fmt.Println(minSwapsCouples([]int{28, 4, 37, 54, 35, 41, 43, 42, 45, 38, 19, 51, 49, 17, 47, 25, 12, 53, 57, 20, 2, 1, 9, 27, 31, 55, 32, 48, 59, 15, 14, 8, 3, 7, 58, 23, 10, 52, 22, 30, 6, 21, 24, 16, 46, 5, 33, 56, 18, 50, 39, 34, 29, 36, 26, 40, 44, 0, 11, 13}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func minSwapsCouples(row []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	pair := func(v int) int {
		if v&1 == 0 {
			return v + 1
		}
		return v - 1
	}

	var dfs func(a []int) int
	dfs = func(a []int) int {
		na := []int{}
		for i := 0; i < len(a); i += 2 {
			if pair(a[i]) != a[i+1] {
				na = append(na, a[i], a[i+1])
			}
		}
		if len(na) < 4 {
			return 0
		}
		if len(na) == 4 {
			return 1
		}

		v0, v1 := na[0], na[1]
		a1, a2 := []int{}, []int{}
		for i := 2; i < len(na); i++ {
			v := na[i]
			if v == pair(v0) {
				a1 = append(a1, v1)
			} else {
				a1 = append(a1, v)
			}
			if v == pair(v1) {
				a2 = append(a2, v0)
			} else {
				a2 = append(a2, v)
			}
		}

		return min(dfs(a1), dfs(a2)) + 1
	}

	return dfs(row)
}

//leetcode submit region end(Prohibit modification and deletion)

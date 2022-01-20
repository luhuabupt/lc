package main

import (
	"fmt"
	"sort"
)

// circus-tower-lcci 马戏团人塔  2022-01-19 19:44:28
func main() {
	//fmt.Println(bestSeqAtIndex([]int{65,70,56,75,60,68}, []int{100,150,90,190,95,110}))
	fmt.Println(bestSeqAtIndex([]int{1, 2, 2, 3}, []int{1, 4, 3, 5}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func bestSeqAtIndex(height []int, weight []int) int {
	arr := [][2]int{}
	for i, v := range height {
		arr = append(arr, [2]int{v, weight[i]})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0] || arr[i][0] == arr[j][0] && arr[i][1] > arr[j][1]
	})

	dp := [][2]int{arr[0]}
	for _, v := range arr {

		x := sort.Search(len(dp), func(i int) bool {
			return v[1] < dp[i][1]
		})
		//fmt.Println(x)
		if x == len(dp) {
			dp = append(dp, v)
		} else {
			dp[x] = v
		}

		//fmt.Println( v, dp)
	}

	return len(dp)
}

//leetcode submit region end(Prohibit modification and deletion)

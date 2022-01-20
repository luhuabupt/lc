package main

import "fmt"

// find-longest-subarray-lcci  字母与数字  2022-01-19 23:56:03
func main() {
	fmt.Println(findLongestSubarray([]string{"A", "1", "B", "C", "D", "2", "3", "4", "E", "5", "F", "G", "6", "7", "H", "I", "J", "K", "L", "M"}))
	fmt.Println(findLongestSubarray([]string{"A", "A"}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findLongestSubarray(array []string) []string {
	diff := map[int]int{0: -1}
	x := 0
	ml, mr := 0, -1
	for i, v := range array {
		if v[0] >= '0' && v[0] <= '9' {
			x++
		} else {
			x--
		}
		if _, ok := diff[x]; ok {
			if i-diff[x] > mr-ml+1 {
				ml, mr = diff[x]+1, i
			}
		} else {
			diff[x] = i
		}
	}
	return array[ml : mr+1]
}

//leetcode submit region end(Prohibit modification and deletion)

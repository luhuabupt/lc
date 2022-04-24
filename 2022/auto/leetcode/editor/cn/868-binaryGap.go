package main

import "fmt"

func main() {
	fmt.Println(binaryGap(22))
	fmt.Println(binaryGap(8))
	fmt.Println(binaryGap(5))
}

//leetcode submit region begin(Prohibit modification and deletion)
func binaryGap(n int) int {
	a := GetBinary(n)

	ans := 0
	pre := 50
	for i, v := range a {
		if v == 1 {
			if i-pre > ans {
				ans = i - pre
			}
			pre = i
		}
	}
	return ans
}

func GetBinary(v int) []int {
	ans := []int{}
	for ; v > 0; v /= 2 {
		ans = append(ans, v&1)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

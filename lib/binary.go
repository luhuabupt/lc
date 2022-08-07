package main

import "fmt"

func main() {
	fmt.Println(GetBinary(8))
	fmt.Println(GetDecimal(1234))
}

// 左低
func GetBinary(v int) []int {
	ans := []int{}
	for ; v > 0; v /= 2 {
		ans = append(ans, v&1)
	}
	return ans
}

func GetDecimal(v int) []int {
	ans := []int{}
	for ; v > 0; v /= 10 {
		ans = append(ans, v%10)
	}
	return ans
}

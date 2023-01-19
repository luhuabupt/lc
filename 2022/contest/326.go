package main

import "fmt"

func main() {
	fmt.Println(" ")
}

func minimumPartition(s string, k int) int {
	ans := 1
	t := 0
	for _, v := range s {
		t *= 10
		t += int(v - '0')
		if t > k {
			t = int(v - '0')
			if t > k {
				return -1
			}
			ans++
		}
	}
	return ans
}

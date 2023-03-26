package main

import "fmt"

func main() {
	a := make([]int, 168)
	for i, _ := range a {
		a[i] = 2
	}
	fmt.Println(cal(a, 200))

	a[166] = 167
	fmt.Println(cal(a, 167))
	fmt.Println(cal(a, 1))
}

func cal(a []int, k int) []int {
	ma := 0
	ans := []int{-1, -1}

	for i := 0; i < 168; i++ {
		t := 0
		for j := i; j-i < 168; j++ {
			t += a[j%168]
			if t > k {
				break
			}
			if j-i+1 > ma {
				ma = j - i + 1
				ans = []int{i, j % 168}
			}
		}
	}

	return ans
}

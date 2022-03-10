package main

import "fmt"

func main() {
	//fmt.Println(sortJumbled([]int{8,9,4,0,2,1,3,5,7,6}, []int{991,338,38}))
	fmt.Println(minMovesToMakePalindrome("aabb"))
	fmt.Println(minMovesToMakePalindrome("letelt"))
	fmt.Println(minMovesToMakePalindrome("aaaba"))
	fmt.Println(minMovesToMakePalindrome("annnnnaaa"))
}

func minMovesToMakePalindrome(s string) int {
	//fmt.Println(s)
	n := len(s)
	if n <= 2 {
		return 0
	}
	if s[0] == s[n-1] {
		return minMovesToMakePalindrome(s[1 : n-1])
	}
	if n%2 == 1 {
		cnt := 0
		for _, v := range s {
			if v == int32(s[0]) {
				cnt++
			}
		}
		if cnt == 1 {
			return n/2 + minMovesToMakePalindrome(s[1:])
		}
		for _, v := range s {
			if v == int32(s[n-1]) {
				cnt++
			}
		}
		if cnt == 1 {
			return n/2 + minMovesToMakePalindrome(s[:n-1])
		}
	}

	a, b := 1, n-2 // s[n-1], s[0]
	for i := 1; i < n; i++ {
		if s[i] == s[n-1] {
			a = i
			break
		}
	}
	for i := n - 2; i >= 0; i-- {
		if s[i] == s[0] {
			b = i
			break
		}
	}
	//fmt.Println(a, b)

	if a < n-1-b {
		return a + minMovesToMakePalindrome(s[:a]+s[a+1:n-1])
	} else { // 干掉s[0]
		return n - 1 - b + minMovesToMakePalindrome(s[1:b]+s[b+1:])
	}
}

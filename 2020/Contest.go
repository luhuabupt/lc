package main

import "fmt"

func makeFancyString(s string) string {
	if len(s) < 3 {
		return s
	}
	a, b, arr, del := s[0], s[1], []byte(s), map[int]bool{}
	for i := 2; i < len(s); i++ {
		if arr[i] == a && a == b {
			del[i] = true
		} else {
			a, b = b, arr[i]
		}
	}
	ans := []byte{}
	for k, v := range arr {
		if !del[k] {
			ans = append(ans, v)
		}
	}
	return string(ans)
}

func main() {
	//p, word := []string{"a","abc","bc","d"}, "abc"
	fmt.Println(1<<3 - 2)
}

func numOfStrings(patterns []string, word string) int {
	ans, flag := 0, false
	for _, s := range patterns {
		flag = false
		for l := 0; l < len(word); l++ {
			if flag {
				break
			}
			for i, j := l, 0; j < len(s); j++ {
				if s[j] != word[i] {
					break
				}
				if j == len(s)-1 {
					ans++
					flag = true
					break
				}
				if i == len(word)-1 {
					break
				}
				i++
			}
		}
	}
	return ans
}

func rearrangeArray_(nums []int) []int {
	for i := 1; i <= len(nums)-2; i++ {
		if nums[i]*2 == nums[i+1]+nums[i-1] {
			nums[i], nums[i-1] = nums[i-1], nums[i]
		} else {
			i++
		}
	}
	return nums
}

func minNonZeroProduct(p int) int {
	m := int(1e9 + 7)
	return ((1<<p - 1) % m) * (qp((1<<p-2)%m, 1<<(p-1)-1)) % m
}

func qp(x, n int) int {
	ans := 1
	m := int(1e9 + 7)
	for n > 0 {
		if n%2 == 1 {
			ans *= x
			ans %= m
			n--
		}
		x *= x
		x %= m
		n /= 2
	}
	return ans
}

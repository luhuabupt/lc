package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestSubsequence("100110111111000000010011101000111011000001000111010001010111100001111110110010100011100100111000011011000000100001011000000100110110001101011010011", 522399436)) // 92
}

func sellingWood(m int, n int, p [][]int) int64 {
	do := make([][][]int, m+1) // h: {w, p}
	for _, v := range p {
		do[v[0]] = append(do[v[0]], []int{v[1], v[2]})
	}
	for i, v := range do {
		sort.Slice(v, func(x, y int) bool {
			return v[x][1]*do[i][y][0] > v[y][1]*do[i][x][0]
		})
	}
	dp := make([]int64, m+1)
	for i := 1; i <= m; i++ {
		tma, tn := int64(0), n
		for x := 0; x < len(do[i]); x++ {
			tma += int64((tn / do[i][x][0]) * do[i][x][1])
			tn -= (tn / do[i][x][0]) * do[i][x][0]
			if tn == 0 {
				break
			}
		}
		dp[i] = tma
		if i == 0 {
			continue
		}
		for j := 1; j <= i/2; j++ {
			if dp[j]+dp[i-j] > dp[i] {
				dp[i] = dp[j] + dp[i-j]
			}
		}
	}
	return dp[m]
}

func longestSubsequence(s string, k int) int {
	n := len(s)
	t := 0
	ans := 0
	flag := false
	for i := n - 1; i >= 0; i-- {
		v := s[i]
		if v == '1' {
			if (n-1-i) >= 63 || 1<<(n-1-i) > k {
				flag = true
			}
			t += 1 << (n - 1 - i)
			if !flag && t <= k {
				ans++
			} else {
				flag = true
				//fmt.Println(i, t - (1 << (n - 1 - i)), ans)
				//fmt.Println(t)
				//fmt.Printf("%b \n", t)
				//break
			}
		} else {
			//fmt.Println(i, ans)
			ans++
		}
	}
	return ans
}

func minimumNumbers(num int, k int) int {
	if num == 0 {
		return 0
	}
	if k == 0 {
		if num%10 == 0 {
			return 1
		}
		return -1
	}
	if k == 1 {
		if num%10 == 0 {
			return 10
		}
		return num % 10
	}
	ans := -1
	for i := k; i <= num; i += k {
		if (num-i)%10 == 0 {
			ans = i / k
			return ans
		}
	}
	return ans
}

func greatestLetter(s string) string {
	a := [26]int{}
	b := [26]int{}
	for i, v := range s {
		if v >= 'a' && v <= 'z' {
			a[v-'a'] = i + 1
		} else {
			b[v-'A'] = i + 1
		}
	}
	ans := ""
	for i := 25; i >= 0; i-- {
		if a[i] > 0 && b[i] > 0 {
			ans = string(i + 'A')
			return ans
		}
	}
	return ans
}

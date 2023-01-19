package main

import "fmt"

func main() {
	fmt.Println(countPalindromes("103301"))
	fmt.Println(countPalindromes("0000000"))
	fmt.Println(countPalindromes("9999900000"))
}

func countPalindromes(s string) int {
	M := int(1e9) + 7
	n := len(s)
	l := make([][100]int, n)
	r := make([][100]int, n)
	cl := [10]int{}
	cr := [10]int{}

	for i := 0; i < n; i++ {
		v := s[i]
		if i > 0 {
			for j := 0; j < 100; j++ {
				l[i][j] = l[i-1][j]
			}
			for j := 0; j < 10; j++ {
				l[i][j*10+int(v-'0')] += cl[j]
			}
		}
		cl[v-'0']++
	}

	for i := n - 1; i >= 0; i-- {
		v := s[i]
		if i < n-1 {
			for j := 0; j < 100; j++ {
				r[i][j] = r[i+1][j]
			}
			for j := 0; j < 10; j++ {
				r[i][j*10+int(v-'0')] += cr[j]
			}
		}
		cr[v-'0']++
	}

	//for i := 0; i < n; i++ {
	//	for j := 0; j < 100; j++ {
	//		if l[i][j] > 0 {
	//			fmt.Println(i, j, l[i][j])
	//		}
	//		if r[i][j] > 0 {
	//			fmt.Println(i, j, r[i][j])
	//		}
	//	}
	//}

	ans := 0
	for i := 2; i <= n-3; i++ {
		for j := 0; j < 100; j++ {
			ans += l[i-1][j] * r[i+1][j]
			ans %= M
		}
	}
	return ans
}

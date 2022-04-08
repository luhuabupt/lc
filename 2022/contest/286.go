package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(123)
	//fmt.Println(minDeletion([]int{1, 1, 2, 2, 3, 3}))
	//fmt.Println(minDeletion([]int{1, 1, 2, 3, 5}))
	//fmt.Println(minDeletion([]int{1, 1, 2, 1, 2, 1, 2}))
	//fmt.Println(minDeletion([]int{1, 1, 1}))
	//fmt.Println(minDeletion([]int{1, 1, 2, 1}))
	//
	//fmt.Println(kthPalindrome([]int{1, 2, 3, 4, 5, 90}, 3))
	fmt.Println(maxValueOfCoins([][]int{{100, 4, 60}, {50, 50, 50}}, 2))
	fmt.Println(maxValueOfCoins([][]int{{100, 4, 60}, {50, 50, 50}}, 3))
	fmt.Println(maxValueOfCoins([][]int{{100, 4, 60}, {50, 50, 50}}, 4))
	fmt.Println(maxValueOfCoins([][]int{{100, 4, 60}, {50, 50, 50}}, 5))
	//fmt.Println(maxValueOfCoins([][]int{{100, 4, 60}, {50, 50, 50}}, 3))
}

func maxValueOfCoins(p [][]int, k int) int {
	dp := make([]int, k+1)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for _, a := range p {
		l := len(a)
		s := make([]int, k+1)
		for i := 1; i <= min(l, k); i++ {
			s[i] += s[i-1] + a[i-1]
		}
		for i := k; i > 0; i-- {
			for w := 0; w <= i && w <= l; w++ {
				dp[i] = max(dp[i], dp[i-w]+s[w])
			}
		}
	}

	return dp[k]
}

func kthPalindrome(q []int, k int) []int64 {
	qp := func(a, n int) int {
		ans := 1
		for n > 0 {
			if n&1 > 0 {
				ans *= a
			}
			a *= a
			n /= 2
		}
		return ans
	}

	n := len(q)
	t := (k + 1) / 2
	st := qp(10, t-1)
	ma := qp(10, t) - 1

	re := func(v int) int64 {
		sv := strconv.Itoa(v)
		as := []byte(sv)
		//if n % 2 == 0 {
		//	as = append(as, as[len(as)-1])
		//}
		for x := (k / 2) - 1; x >= 0; x-- {
			as = append(as, as[x])
		}

		ai, _ := strconv.Atoi(string(as))
		return int64(ai)
	}

	ans := make([]int64, n)
	for i, v := range q {
		x := st + v - 1
		if x > ma {
			ans[i] = -1
		} else {
			ans[i] = re(x)
		}
	}

	return ans
}

func minDeletion(a []int) int {
	n := len(a)
	ans := 0
	idx := 1
	pre := a[0]
	for i := 1; i < n; i++ {
		if idx%2 == 1 && a[i] == pre {
			ans++
		} else {
			if idx%2 == 0 {
				pre = a[i]
			}
			idx++
		}
	}
	if (n-ans)%2 == 1 {
		ans++
	}

	return ans
}

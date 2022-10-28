package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(atMostNGivenDigitSet([]string{"1", "3", "5", "7"}, 100))
	fmt.Println(atMostNGivenDigitSet([]string{"1", "4", "9"}, 1000000000))
	fmt.Println(atMostNGivenDigitSet([]string{"1", "4", "9"}, 4))
	fmt.Println(atMostNGivenDigitSet([]string{"1", "2", "3", "4"}, 2345))
}

//leetcode submit region begin(Prohibit modification and deletion)
func atMostNGivenDigitSet(d []string, ni int) int {
	n := strconv.Itoa(ni)
	dp := make([][2]int, len(n)+1)
	m := len(d)

	dp[0][1] = 1
	for i := 1; i <= len(n); i++ {
		for _, v := range d {
			if v[0] < n[i-1] {
				dp[i][0] += dp[i-1][1]
			} else if v[0] == n[i-1] {
				dp[i][1] += dp[i-1][1]
			} else {
				break
			}
		}
		if i > 1 {
			dp[i][0] += m + m*(dp[i-1][0])
		}
	}
	//fmt.Println(dp)
	return dp[len(n)][0] + dp[len(n)][1]
}

func atMostNGivenDigitSet_(ds []string, n int) int {
	d := []int{}
	for _, v := range ds {
		d = append(d, int(v[0]-'0'))
	}
	sort.Ints(d)

	a := []int{}
	for x := n; x > 0; x /= 10 {
		a = append(a, x%10)
	}

	dp := make([]int, len(a))
	qp := 1
	ans := -1
	for i := 0; i < len(a); i++ {
		for _, v := range d {
			if v == a[i] {
				if i > 0 {
					dp[i] += dp[i-1]
				} else {
					dp[i]++
				}
				break
			} else if v > a[i] {
				break
			}
			dp[i] += qp
		}
		ans += qp
		qp *= len(d)
	}

	//fmt.Println(dp)

	return ans + dp[len(a)-1]
}

//leetcode submit region end(Prohibit modification and deletion)

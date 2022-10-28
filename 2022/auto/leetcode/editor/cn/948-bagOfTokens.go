package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(bagOfTokensScore([]int{1, 2, 3, 4}, 2))
	fmt.Println(bagOfTokensScore([]int{10, 20}, 15))
	fmt.Println(bagOfTokensScore([]int{}, 15))
	fmt.Println(bagOfTokensScore([]int{68, 85, 34, 25, 60}, 44))
	fmt.Println(bagOfTokensScore([]int{1, 2, 3, 4, 5}, 2))
}

//leetcode submit region begin(Prohibit modification and deletion)
func bagOfTokensScore(a []int, power int) int {
	sort.Ints(a)
	n := len(a)
	if n == 0 || a[0] > power {
		return 0
	}

	t := 0
	s := make([]int, n)
	for i, v := range a {
		t += v
		s[i] = t
	}

	return sort.Search(n, func(k int) bool {
		do := (n - k) / 2
		sum := power + s[n-1] - s[n-1-do]
		use := s[do+k]
		return use > sum
	})
}

//leetcode submit region end(Prohibit modification and deletion)

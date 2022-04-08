package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(nextGreaterElement(241))
	fmt.Println(nextGreaterElement(21))
	fmt.Println(nextGreaterElement(124))
	fmt.Println(nextGreaterElement(124))
	fmt.Println(nextGreaterElement(1224))
	fmt.Println(nextGreaterElement(476))
	fmt.Println(nextGreaterElement(2147483476))
}

//leetcode submit region begin(Prohibit modification and deletion)
func nextGreaterElement(n int) int {
	ar := []int{}
	for x := n; x > 0; x /= 10 {
		ar = append(ar, x%10)
	}
	for i := 1; i < len(ar); i++ {
		if ar[i] < ar[i-1] {
			for y := 0; y < i; y++ {
				if ar[y] > ar[i] {
					ar[i], ar[y] = ar[y], ar[i]
					break
				}
			}
			x := ar[:i]
			xe := ar[i:]
			sort.Ints(x)
			for j := 0; j < len(x)/2; j++ {
				x[j], x[len(x)-1-j] = x[len(x)-1-j], x[j]
			}
			ar = append(x, xe...)
			break
		}
	}
	ans := 0
	for i := len(ar) - 1; i >= 0; i-- {
		ans *= 10
		ans += ar[i]
	}
	if ans == n || ans > (1<<31)-1 {
		return -1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

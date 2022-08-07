//有些数的素因子只有 3，5，7，请设计一个算法找出第 k 个数。注意，不是必须有这些素因子，而是必须不包含其他的素因子。例如，前几个数按顺序应该是 1，3，
//5，7，9，15，21。
//
// 示例 1:
//
// 输入: k = 5
//
//输出: 9
//
// Related Topics 哈希表 数学 动态规划 堆（优先队列）
// 👍 77 👎 0

package main

import "fmt"

func main() {
	fmt.Println(getKthMagicNumber(4))
	fmt.Println(getKthMagicNumber(5))
	fmt.Println(getKthMagicNumber(6))
	fmt.Println(getKthMagicNumber(7))
	fmt.Println(getKthMagicNumber(8))
	fmt.Println(getKthMagicNumber(9))
}

//leetcode submit region begin(Prohibit modification and deletion)
func getKthMagicNumber(k int) int {
	min := func(a, b, c int) (int, int) {
		if a < b && a < c {
			return a, 0
		}
		if b < c {
			return b, 1
		}
		return c, 2
	}

	idx := [3]int{0, 0, 0}
	ans := []int{1}
	for len(ans) < k {
		x, i := min(ans[idx[0]]*3, ans[idx[1]]*5, ans[idx[2]]*7)
		idx[i]++

		if x != ans[len(ans)-1] {
			ans = append(ans, x)
		}
	}

	return ans[k-1]
}

//leetcode submit region end(Prohibit modification and deletion)

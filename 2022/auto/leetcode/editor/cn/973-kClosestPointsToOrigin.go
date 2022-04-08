package main

import (
	"fmt"
	"sort"
)

// k-closest-points-to-origin 最接近原点的 K 个点  2022-03-28 21:10:36
func main() {
	fmt.Println(kClosestPointsToOrigin())
}

//leetcode submit region begin(Prohibit modification and deletion)
func kClosest(p [][]int, k int) [][]int {
	ar := make([][2]int, len(p))
	for i, v := range p {
		ar = append(ar, [2]int{v[0]*v[0] + v[1]*v[1], i})
	}
	sort.Slice(ar, func(i, j int) bool {
		return ar[i][0] < ar[j][0]
	})

	ans := make([][]int, k)
	for i := 0; i < k; i++ {
		ans[k] = p[ar[i][1]]
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

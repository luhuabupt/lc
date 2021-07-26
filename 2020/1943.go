package _020

import "sort"

// 差分 前缀和
func splitPainting(segments [][]int) [][]int64 {
	diff, list, ans, sum := map[int]int64{}, []int{}, [][]int64{}, int64(0)
	for _, arr := range segments {
		diff[arr[0]] += int64(arr[2])
		diff[arr[1]] -= int64(arr[2])
	}
	for k, _ := range diff {
		list = append(list, k)
	}
	sort.Ints(list)
	for k, v := range list {
		if k > 0 && sum > 0 {
			ans = append(ans, []int64{int64(list[k-1]), int64(v), sum})
		}
		sum += diff[v]
	}
	return ans
}
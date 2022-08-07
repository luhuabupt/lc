package main

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

func splitPainting_(segments [][]int) [][]int64 {
	ans := [][]int64{}
	sort.Slice(segments, func(i, j int) bool {
		return segments[i][0] < segments[j][0]
	})
	pm, points := map[int]bool{}, []int{}
	for _, arr := range segments {
		pm[arr[0]], pm[arr[1]] = true, true
	}
	for k, _ := range pm {
		points = append(points, k)
	}
	sort.Ints(points)
	sum, preSum, stack, idx, stackIdx := int64(0), int64(0), []int{}, 0, 0
	for k, point := range points {
		for i := stackIdx; i < len(stack); i++ {
			seg := stack[stackIdx]
			if segments[seg][1] <= point {
				sum -= int64(segments[seg][2])
				stackIdx++
			}
		}
		for idx < len(segments) && segments[idx][0] <= point {
			sum += int64(segments[idx][2])
			stack = append(stack, idx)
			idx++
		}
		if k > 0 && preSum > 0 {
			ans = append(ans, []int64{int64(points[k-1]), int64(point), preSum})
		}
		preSum = sum
	}
	return ans
}

package main

import "sort"

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n1, n2 := len(nums1), len(nums2)
	ans := [][]int{}

	may := []int{}
	for i := 0; i < n1 && i < k && k-1-i < n2; i++ {
		may = append(may, nums1[i]+nums2[k-1-i])
	}

	kth := 1 << 32
	for _, v := range may {
		if v < kth {
			kth = v
		}
	}

	for _, v := range nums1 {
		for _, vv := range nums2 {
			if v+vv <= kth {
				ans = append(ans, []int{v, vv})
				if len(ans) > k && v+vv == kth {
					break
				}
			} else {
				break
			}
		}
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0]+ans[i][1] < ans[j][0]+ans[j][1]
	})

	if len(ans) > k {
		ans = ans[:k]
	}
	return ans
}

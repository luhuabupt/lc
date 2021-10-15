package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minAbsDifference([]int{-2772, 6927, 4773, -2687, 7167, -8995, 2940, 8869, 526}, 969621127))
}

func minimumDifference(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	half := sum / 2

	n := len(nums) / 2
	a := make([][]int, n+1)
	b := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		a[i] = []int{}
		b[i] = []int{}
	}

	for i := 0; i < 1<<n; i++ {
		l, v, v2 := 0, 0, 0
		for j := 0; j < n; j++ {
			if (1<<j)&i > 0 {
				v += nums[j]
				v2 += nums[n+j]
				l++
			}
		}
		a[l] = append(a[l], v)
		b[l] = append(b[l], v2)
	}

	for l := 0; l <= n; l++ {
		sort.Ints(a[l])
		sort.Ints(b[l])
	}

	ans := sum - 2*a[n][0]
	if ans < 0 {
		ans = -ans
	}
	//fmt.Println(a, b)

	for l := 0; l <= n; l++ {
		//fmt.Println(a[l], b[n-l])
		for i, j := 0, len(b[n-l])-1; i < len(a[l]); i++ {
			for a[l][i]+b[n-l][j] > half && j > 0 {
				//fmt.Println(i, j)
				ans = minAns(a[l][i]+b[n-l][j], sum, ans)
				j--
			}
			//fmt.Println(i, j)
			ans = minAns(a[l][i]+b[n-l][j], sum, ans)
		}
	}

	return ans
}

func minAns(a, sum, ans int) int {
	d := sum - 2*a
	if d < 0 {
		d = -d
	}
	if d < ans {
		return d
	}
	return ans
}

func minAbsDifference(nums []int, goal int) int {
	// [-2772,6927,4773,-2687,7167,-8995,2940,8869,526]
	//969621127 | 969590451 969589925
	n := len(nums) / 2
	a, b := []int{}, []int{}

	for i := 0; i < 1<<n; i++ {
		v := 0
		for j := 0; j < n; j++ {
			if (1<<j)&i > 0 {
				v += nums[j]
			}
		}
		a = append(a, v)
	}
	for i := 0; i < 1<<(len(nums)-n); i++ {
		v := 0
		for j := 0; j < len(nums)-n; j++ {
			if (1<<j)&i > 0 {
				v += nums[j+n]
			}
		}
		b = append(b, v)
	}

	sort.Ints(a)
	sort.Ints(b)

	//fmt.Println(a)
	//fmt.Println(b)

	ans := ab(a[0] - goal)

	for i, j := 0, len(b)-1; i < len(a); i++ {
		for j > 0 && a[i]+b[j] > goal {
			if ab(a[i]+b[j]-goal) < ans {
				ans = ab(a[i] + b[j] - goal)
			}
			j--
		}
		if ab(a[i]+b[j]-goal) < ans {
			ans = ab(a[i] + b[j] - goal)
		}
	}

	return ans
}

func ab(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 换位置保证nums1先遍历
	if len(nums2) < len(nums1) || len(nums2) == len(nums1) && nums2[len(nums2)-1] < nums1[len(nums1)-1] {
		return findMedianSortedArrays(nums2, nums1)
	}

	i, j, l, cur, prev := 0, 0, 1, 0, 0
	for i < len(nums1) && j < len(nums2) && l <= (len(nums1)+len(nums2)+2)/2 {
		if i < len(nums1) && nums1[i] <= nums2[j] {
			prev, cur = cur, nums1[i]
			i++
		} else {
			prev, cur = cur, nums2[j]
			j++
		}
		l++
	}

	for j < len(nums2) && l <= (len(nums1)+len(nums2)+2)/2 {
		prev, cur = cur, nums2[j]
		j++
		l++
	}

	if (len(nums1)+len(nums2))%2 == 0 {
		return (float64(prev) + float64(cur)) / 2
	} else {
		return float64(cur)
	}
}

func findJudge(n int, trust [][]int) int {
	m := map[int]map[int]bool{}
	t := map[int]bool{}
	for _, arr := range trust {
		if m[arr[1]] == nil {
			m[arr[1]] = map[int]bool{}
		}
		m[arr[1]][arr[0]] = true
		t[arr[0]] = true
	}
	for k, arr := range m {
		if len(arr) == n-1 && !arr[k] && !t[k] {
			return k
		}
	}
	return -1
}

func peakIndexInMountainArray(arr []int) int {
	i, j := 0, len(arr)-1
	for {
		mid := (i + j) / 2
		if mid > 0 && mid < len(arr)-1 && arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		}
		if arr[mid] > arr[i] && arr[mid] > arr[j] {
			i++
			j--
		} else if arr[mid] <= arr[i] && arr[mid] > arr[j] {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
}

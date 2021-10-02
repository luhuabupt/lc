package main

import "sort"

func main() {

}

func countKDifference(nums []int, k int) int {
	ans := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]-nums[i] == k || nums[i]-nums[j] == k {
				ans++
			}
		}
	}
	return ans
}
func findOriginalArray(changed []int) []int {
	m := map[int]int{}
	for _, v := range changed {
		m[v]++
	}
	sort.Ints(changed)
	ans := []int{}
	for i := 0; i < len(changed)/2; i++ {
		if m[changed[i]] == 0 {
			continue
		}
		m[changed[i]]--
		if m[2*changed[i]] > 0 {
			m[2*changed[i]]--
			ans = append(ans, changed[i])
		} else {
			return []int{}
		}
	}
	return ans
}

func maxTaxiEarnings(n int, rides [][]int) int64 {
	m := map[int][]int{}
	for i, v := range rides {
		if m[v[1]] == nil {
			m[v[1]] = []int{}
		}
		m[v[1]] = append(m[v[1]], i)
	}

	dp := map[int]int64{1: 0}
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1]
		if m[i] == nil {
			continue
		}
		max := dp[i-1]
		for _, x := range m[i] {
			ri := rides[x]
			sum := dp[ri[0]] + int64(i-ri[0]+ri[2])
			if sum > max {
				max = sum
			}
		}
		dp[i] = max
	}
	return dp[n]
}

func minOperations(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	win := []int{}
	max, maxI, pre := 0, 1, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == pre {
			continue
		}
		pre = nums[i]

		win = append(win, nums[i])
		for _, v := range win {
			if v < nums[i]+1-n {
				win = win[1:]
			} else {
				break
			}
		}
		if len(win) > max {
			max = len(win)
			maxI = nums[i]
		}
	}
	ans := n
	m := map[int]bool{}
	for _, v := range nums {
		m[v] = true
	}
	for i := maxI - n + 1; i <= maxI; i++ {
		if m[i] {
			ans--
		}
	}
	return ans
}

func sumOfBeauties(nums []int) int {
	n := len(nums)
	l, r := make([]int, n), make([]int, n)
	l[0], r[n-1] = nums[0], nums[n-1]
	for i := 1; i < n; i++ {
		l[i] = l[i-1]
		if nums[i] > l[i-1] {
			l[i] = nums[i]
		}
	}
	for i := n - 2; i >= 0; i-- {
		r[i] = r[i+1]
		if nums[i] < r[i+1] {
			r[i] = nums[i]
		}
	}
	ans := 0
	for k, v := range nums {
		if k == 0 || k == n-1 {
			continue
		}
		if v > l[k-1] && v < r[k+1] {
			ans += 2
		} else if v > nums[k-1] && v < nums[k+1] {
			ans++
		}
	}
	return ans
}

type DetectSquares struct {
	m map[int]map[int]int
}

func Constructor() DetectSquares {
	return DetectSquares{map[int]map[int]int{}}
}

func (this *DetectSquares) Add(point []int) {
	if this.m[point[0]] == nil {
		this.m[point[0]] = map[int]int{}
	}
	this.m[point[0]][point[1]]++
}

func (this *DetectSquares) Count(point []int) int {
	ans := 0
	for y, c := range this.m[point[0]] {
		l := y - point[1]
		if l == 0 {
			continue
		}

		xi := point[0] - l
		if this.m[xi] != nil && this.m[xi][point[1]] > 0 && this.m[xi][y] > 0 {
			ans += c * this.m[xi][point[1]] * this.m[xi][y]
		}
		xi = point[0] + l
		if this.m[xi] != nil && this.m[xi][point[1]] > 0 && this.m[xi][y] > 0 {
			ans += c * this.m[xi][point[1]] * this.m[xi][y]
		}
	}
	return ans
}

func minSteps(n int) int {
	if n == 1 {
		return 0
	}
	for i := n / 2; i > 0; i-- {
		if n%i == 0 {
			return minSteps(i) + (n / i)
		}
	}
	return 1
}

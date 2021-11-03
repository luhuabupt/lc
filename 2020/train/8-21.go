package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	//fmt.Println(minTimeToType("abc"))
	//fmt.Println(minTimeToType("bza"))
	//fmt.Println(minTimeToType("zjpc"))
	fmt.Println(countPaths(7, [][]int{[]int{0, 6, 7}, []int{0, 1, 2}, []int{1, 2, 3}, []int{1, 3, 3}, []int{6, 3, 3}, []int{3, 5, 1}, []int{6, 5, 1}, []int{2, 5, 1}, []int{0, 4, 5}, []int{4, 6, 2}}))
}

func minTimeToType(word string) int {
	ans, pre := 0, 'a'
	for _, v := range word {
		ans += getX(v, pre)
		ans++
		pre = v
	}
	return ans
}
func getX(a, b int32) int {
	if a > b {
		return getX(b, a)
	}
	shun := b - a
	ni := 'z' - b + a - 'a' + 1
	if shun < ni {
		return int(shun)
	}
	return int(ni)
}

func maxMatrixSum(matrix [][]int) int64 {
	min, cnt, zero, sum := 100001, 0, 0, int64(0)
	for _, arr := range matrix {
		for _, v := range arr {
			if v < 0 {
				cnt++
				sum -= int64(v)
				if -v < min {
					min = -v
				}
			} else {
				sum += int64(v)
				if v < min {
					min = v
				}
			}
			if v == 0 {
				zero++
				min = 0
			}
		}
	}
	if cnt%2 == 0 {
		return sum
	}
	if cnt%2 == 1 && zero > 0 {
		return sum
	}
	return sum - int64(min) - int64(min)
}

type pair struct {
	start int
	i     int
	t     int // 总时间
}
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { a, b := h[i], h[j]; return a.t < b.t }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	ans := []pair{v}
	i := len(a) - 2
	for i >= 0 && a[i].t == v.t {
		ans = append(ans, a[i])
		i--
	}
	*h = a[:i+1]
	return ans
}

func countPaths(n int, roads [][]int) int {
	dp := make([][2]int, n)
	dp[0] = [2]int{0, 1} // time, cnt
	mod := int(1e9) + 7

	m := map[int][][]int{}
	for _, arr := range roads { // 出图
		m[arr[0]] = append(m[arr[0]], []int{arr[1], arr[2]}) //  目的地, 时间
	}

	h := &hp{}
	for _, x := range m[0] { // init
		heap.Push(h, pair{0, x[0], x[1]})
	}

	flag := false
	for !flag {
		arr := heap.Pop(h).([]pair)
		for _, x := range arr {
			if x.i == n-1 {
				flag = true
			}
			if dp[x.i][0] == 0 {
				dp[x.i] = [2]int{0, 0}
				dp[x.i][0] = x.t
			}
			if x.t == dp[x.i][0] { // 最短路径
				dp[x.i][1] += dp[x.start][1]
				dp[x.i][1] %= mod
				for _, xx := range m[x.i] {
					heap.Push(h, pair{x.i, xx[0], dp[x.i][0] + xx[1]})
				}
			}
		}
		fmt.Println(arr, h, dp)
	}
	return dp[n-1][1]
}

func findGCD(nums []int) int {
	sort.Ints(nums)
	ans := 1
	for i := 1; i <= nums[0]; i++ {
		if nums[0]%i == 0 && nums[len(nums)-1]%i == 0 {
			ans = i
		}
	}
	return ans
}

func findDifferentBinaryString(nums []string) string {
	m, n := map[string]bool{}, len(nums[0])
	for _, v := range nums {
		m[v] = true
	}
	for i := 0; i < (1 << len(nums)); i++ {
		if !m[toBin(i, n)] {
			return toBin(i, n)
		}
	}
	return ""
}

func toBin(n, l int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	for i := len(result); i < l; i++ {
		result = "0" + result
	}
	return result
}

func minimizeTheDifference(mat [][]int, target int) int {
	m := map[int]bool{}
	for _, v := range mat[0] {
		m[v] = true
	}

	for k, arr := range mat {
		if k == 0 {
			continue
		}
		nm := map[int]bool{}
		min := -1
		for _, v := range arr {
			for pre, _ := range m {
				if pre+v < target {
					nm[pre+v] = true
				} else {
					if min == -1 || pre+v < min {
						min = pre + v
					}
				}
			}
		}

		if min > 0 {
			nm[min] = true
		}
		for x, _ := range m {
			if !nm[x] {
				delete(m, x)
			}
		}
		for x, _ := range nm {
			m[x] = true
		}
	}

	ans := target
	for i, _ := range m {
		if i > target && i-target < ans {
			ans = i - target
		}
		if i < target && target-i < ans {
			ans = target - i
		}
		if i == target {
			ans = 0
		}
	}
	return ans
}

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	ans := nums[len(nums)-1]
	for i := 0; i <= len(nums)-k; i++ {
		if nums[i+k]-nums[i] < ans {
			ans = nums[i+k] - nums[i]
		}
	}
	return ans
}

func kthLargestNumber(nums []string, k int) string {
	sort.Slice(nums, func(i, j int) bool {
		a, b := nums[i], nums[j]
		if len(a) == len(b) {
			for x := 0; x < len(a); x++ {
				if a[x] == b[x] {
					continue
				}
				return a[x] < b[x]
			}
		}
		return len(a) < len(b)
	})
	return nums[len(nums)-k]
}

func minSessions(tasks []int, sessionTime int) (ans int) {
	sort.Ints(tasks)
	for i := 0; i < len(tasks)/2; i++ {
		tasks[i], tasks[len(tasks)-1-i] = tasks[len(tasks)-1-i], tasks[i]
	}
	mt := map[int]bool{0: true}
	for len(mt) > 0 {
		ans, mt = ans+1, map[int]bool{}
		for k, _ := range tasks {
			mt[k] = true
		}
		dp, path := map[int]int{}, make([][]int, len(tasks))
		for k, v := range tasks { // 背包
			tmpI := k
			dp[k], path[k] = v, []int{}
			for i := 0; i < k; i++ {
				if dp[i]+v <= sessionTime && dp[k] < dp[i]+v {
					dp[k], tmpI = dp[i]+v, i
				}
			}
			path[k] = append(path[tmpI], k)
		}
		mv, mk := 0, 0
		for i := 0; i < len(tasks); i++ {
			if dp[i] > mv {
				mk, mv = i, dp[i]
			}
		}
		for _, v := range path[mk] { // 删除已装
			delete(mt, v)
		}
		nt := []int{}
		for t, tv := range tasks {
			if mt[t] {
				nt = append(nt, tv)
			}
		}
		tasks = nt
	}
	return ans
}

func numberOfUniqueGoodSubsequences(binary string) int {
	mod := int(1e9) + 7
	n, ans, one := len(binary), 0, len(binary)
	for i := n - 1; i >= 0; i-- {
		if binary[i] == 1 {
			ans += 1 << (n - 1 - i)
			one = i
		} else {
			if one < n {
				ans += 1 << (n - 1 - one)
			}
			ans++
		}
		ans %= mod
	}
	return ans
}

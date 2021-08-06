package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	a := []int{1, 1, 0, 0, 0, 0, 0, 0}
	x := sort.Search(len(a), func(i int) bool { return a[i] == 0 })
	fmt.Println(x)
	//p := [][]int{[]int{2,4,0,5,5},[]int{0,5,4,2,5},[]int{2,0,2,3,1},[]int{3,0,5,5,2}} // 15
	//p := [][]int{[]int{1,2,3},[]int{1,5,1},[]int{3,1,1}}
	//p := [][]int{[]int{5,2,1,2},[]int{2,1,5,2},[]int{5,5,5,0}}
	//p := [][]int{[]int{1,7,9},[]int{6,8,15},[]int{8,10,7}}
	//p = [][]int{[]int{1,7,9},[]int{8,10,7}}
	//p = [][]int{[]int{4,5,9}, []int{8,12,5}, []int{4,7,19}, []int{14,15,1}, []int{3,10,8}, []int{17,20,18}, []int{7,19,14}, []int{8,16,6}, []int{14,17,7}, []int{11,13,3}}
	//fmt.Println(splitPainting(p))
	//fmt.Println(p)
	//fmt.Println(maxPoints(p))
	//fmt.Println(getLucky("leetcode", 2))
	//fmt.Println(maximumNumber("2140101", []int{6,7,9,7,4,0,3,4,4,7}))
	//fmt.Println(restoreArray([][]int{[]int{2,1},[]int{3,4},[]int{3,2}}))
	//fmt.Println(minOperations([]int{11, 16, 20, 1, 2, 13, 7, 6, 12, 3}, []int{11, 13, 3, 7, 7, 1, 10, 12, 14, 1}))
	//fmt.Println(minOperations_([]int{11, 16, 20, 1, 2, 13, 7, 6, 12, 3}, []int{11, 13, 3, 7, 7, 1, 10, 12, 14, 1}))
}

func countSpecialSubsequences(nums []int) int {
	c0, c1, c2, mod := 0, 0, 0, int(1e9+7)
	for _, v := range nums {
		if v == 0 {
			c0 = (c0*2 + 1) % mod
		} else if v == 1 {
			c1 = (c0 + c1*2) % mod
		} else if v == 2 {
			c2 = (c1 + c2*2) % mod
		}
	}
	return c2
}

func numberOfWeeks(milestones []int) int64 {
	max, sum := int64(0), int64(0)
	for _, v := range milestones {
		if v > int(max) {
			max = int64(v)
		}
		sum += int64(v)
	}
	if max > sum-max {
		return (sum-max)*2 + 1
	} else {
		return sum
	}
}

func minimumPerimeter(neededApples int64) int64 {
	i, sum := int64(2), int64(0)
	for {
		sum += 3 * i * i
		if sum >= neededApples {
			break
		}
		i += 2
	}
	return i * 4
}

func getLucky(s string, k int) int {
	ns, next, ans := "", 0, 0
	for _, v := range s {
		ns += strconv.Itoa(int(v - 96))
	}
	for i := 0; i < k; i++ {
		for _, v := range ns {
			next += int(v - 48)
		}
		ns, next = strconv.Itoa(next), 0
	}
	ans, _ = strconv.Atoi(ns)
	return ans
}

func maximumNumber(num string, change []int) string {
	up, eq, ans := map[int]bool{}, map[int]bool{}, []byte(num)
	for k, v := range change {
		if v > k {
			up[k] = true
		}
		if v == k {
			eq[k] = true
		}
	}
	flag := false
	for k, v := range num {
		val := int(v - 48)
		if up[val] {
			flag = true
			ans[k] = byte(change[val] + 48)
		} else if flag && !eq[val] {
			break
		}
	}
	return string(ans)
}

func minOperations_(target, arr []int) int {
	n := len(target)
	pos := make(map[int]int, n)
	for i, val := range target {
		pos[val] = i
	}
	d := []int{}
	for _, val := range arr {
		if idx, has := pos[val]; has {
			fmt.Println(pos[val])
			if p := sort.SearchInts(d, idx); p < len(d) {
				d[p] = idx
			} else {
				d = append(d, idx)
			}
		}
	}
	return n - len(d)
}

func minOperations(target []int, arr []int) int {
	pos, mapping := map[int]int{}, []int{}
	for k, v := range target {
		pos[v] = k
	}
	for _, v := range arr {
		if idx, ok := pos[v]; ok {
			fmt.Println(idx)
			mapping = append(mapping, idx)
		}
	}
	fmt.Println(mapping)
	return len(target) - findMaxUp(mapping)
}

// 最长上升子序列
func findMaxUp(arr []int) int {
	dp := []int{}
	for _, v := range arr {
		if p := sort.SearchInts(dp, v); p < len(dp) {
			dp[p] = v
		} else {
			dp = append(dp, v)
		}
	}
	return len(dp)
}

func findTargetSumWays_(nums []int, target int) int {
	ans := 0
	var backTrack func(i, v int)
	backTrack = func(i, v int) {
		if i == len(nums) {
			if v == target {
				ans++
			}
			return
		}
		backTrack(i+1, v+nums[i])
		backTrack(i+1, v-nums[i])
	}
	backTrack(0, 0)
	return ans
}

type pair2 struct{ pow, idx int }

func kWeakestRows(mat [][]int, k int) []int {
	pairs, ans := make([]pair2, len(mat)), []int{}
	for i, arr := range mat {
		pow := sort.Search(len(arr), func(j int) bool { return arr[j] == 0 })
		pairs[i] = pair2{pow, i}
	}
	sort.Slice(pairs, func(i, j int) bool {
		a, b := pairs[i], pairs[j]
		return a.pow < b.pow || a.pow == b.pow && a.idx < b.idx
	})
	for i := 0; i < k; i++ {
		ans = append(ans, pairs[i].idx)
	}
	return ans
}

func verticalTraversal(root *TreeNode) [][]int {
	m, ans := map[int]map[int][]int{}, [][]int{}

	var dfs func(p *TreeNode, row int, col int)
	dfs = func(p *TreeNode, row int, col int) {
		if p == nil {
			return
		}
		if m[col] == nil {
			m[col] = map[int][]int{}
		}
		if m[col][row] == nil {
			m[col][row] = []int{}
		}
		m[col][row] = append(m[col][row], p.Val)
		dfs(p.Left, row+1, col-1)
		dfs(p.Right, row+1, col+1)
	}
	dfs(root, 0, 0)

	min, max := 0, 0
	for r, _ := range m {
		if r < min {
			min = r
		}
		if r > max {
			max = r
		}
	}

	for i := min; i <= max; i++ {
		if m[i] == nil {
			continue
		}
		one, bottom := []int{}, 0
		for row, _ := range m[i] {
			if row > bottom {
				bottom = row
			}
		}
		for j := 0; j <= bottom; j++ {
			if m[i][j] == nil {
				continue
			}
			if len(m[i][j]) > 1 {
				sort.Ints(m[i][j])
			}
			one = append(one, m[i][j]...)
		}
		ans = append(ans, one)
	}
	return ans
}

func findTargetSumWays(nums []int, target int) int {
	ans, sum, dp := 0, 0, make([]map[int]int, len(nums))
	for _, v := range nums {
		sum += v
	}
	if (sum-target)%2 == 1 || sum < target {
		return ans
	}
	neg := (sum - target) / 2
	for i, v := range nums {
		if i == 0 {
			dp[0] = map[int]int{0: 1}
			dp[0][nums[0]] += 1
			continue
		}
		dp[i] = map[int]int{}
		for j, _ := range dp[i-1] {
			if j+v <= neg {
				dp[i][j+v] += dp[i-1][j]
			}
			dp[i][j] += dp[i-1][j]
		}
	}
	return dp[len(nums)-1][neg]
}

func canSeePersonsCount(heights []int) []int {
	stack, ans := []int{}, make([]int, len(heights))
	for i := len(heights) - 1; i >= 0; i-- {
		idx := len(stack) - 1
		for idx >= 0 && heights[i] > stack[idx] { // heights.i 入栈
			ans[i]++
			idx--
		}
		if len(stack) > 0 && heights[i] < stack[0] {
			ans[i]++
		}
		stack = stack[:idx+1]
		stack = append(stack, heights[i])
	}
	return ans
}

func findSecondMinimumValue(root *TreeNode) int {
	ans := 1 << 32
	var dfs func(p *TreeNode)
	dfs = func(p *TreeNode) {
		if p.Val != root.Val && p.Val < ans {
			ans = p.Val
		}
		if p.Left != nil {
			dfs(p.Left)
			dfs(p.Right)
		}
	}
	dfs(root)
	if ans == 1<<32 {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func restoreArray(adjacentPairs [][]int) []int {
	adm, ans := map[int][]int{}, []int{}
	for _, arr := range adjacentPairs {
		if adm[arr[0]] == nil {
			adm[arr[0]] = []int{}
		}
		if adm[arr[1]] == nil {
			adm[arr[1]] = []int{}
		}
		adm[arr[0]] = append(adm[arr[0]], arr[1])
		adm[arr[1]] = append(adm[arr[1]], arr[0])
	}
	p, prev := 0, 0
	for k, arr := range adm {
		if len(arr) == 1 {
			ans = append(ans, k, arr[0])
			p, prev = arr[0], k
			break
		}
	}
	fmt.Println(prev, p)
	for len(adm[p]) == 2 {
		for _, v := range adm[p] {
			if v != prev {
				prev, p = p, v
				break
			}
		}
		fmt.Println(prev, p)
		ans = append(ans, p)
	}
	return ans
}

func maxCompatibilitySum(students [][]int, mentors [][]int) int {
	ans := 0
	var dfs func(cur [][]int, s [][]int)
	dfs = func(cur [][]int, s [][]int) {
		if len(cur) == len(mentors) {
			tmp := 0
			for i, arr := range cur {
				for j, v := range arr {
					if v == mentors[i][j] {
						tmp++
					}
				}
			}
			if tmp > ans {
				ans = tmp
			}
		}
		for k, v := range s {
			n, nc := [][]int{}, [][]int{}
			for i, c := range s {
				if i == k {
					continue
				}
				n = append(n, c)
			}
			nc = append(cur, v)
			dfs(nc, n)
		}
	}
	dfs([][]int{}, students)
	return ans
}

func sortOne(a []int, b []int) {

}

func areOccurrencesEqual(s string) bool {
	m := map[int32]int{}
	for _, v := range s {
		m[v]++
	}
	pre := -1
	for _, v := range m {
		if pre == -1 {
			pre = v
		}
		if v != pre {
			return false
		}
	}
	return true
}

func maximumTime(time string) string {
	t := []byte(time)
	if t[0] == '?' && t[1] == '?' {
		t[0], t[1] = '2', '3'
	} else if t[0] == '?' {
		if t[1] <= '3' {
			t[0] = '2'
		} else {
			t[0] = '1'
		}
	} else if t[1] == '?' {
		if t[0] <= '1' {
			t[1] = '9'
		} else {
			t[1] = '3'
		}
	}
	if t[2] == '?' {
		t[2] = '5'
	}
	if t[3] == '?' {
		t[3] = '9'
	}

	return time
}

func groupAnagrams(strs []string) [][]string {
	ans, ansM := [][]string{}, map[string][]string{}
	for _, arr := range strs {
		m, s := map[int32]int{}, ""
		for _, v := range arr {
			m[v]++
		}
		for i := 'a'; i <= 'z'; i++ {
			if m[i] > 0 {
				s += string(i) + strconv.Itoa(m[i])
			}
		}
		ansM[s] = append(ansM[s], arr)
	}
	for _, v := range ansM {
		ans = append(ans, v)
	}
	return ans
}

func maxPoints(points [][]int) int64 {
	m, n, ans := len(points), len(points[0]), 0
	dp, l, r := make([][]int, m), make([]int, n), make([]int, n)
	for i, arr := range points {
		dp[i] = make([]int, n)
		if i == 0 { // 第一行
			dp[i] = arr
			continue
		}
		for k := 0; k < n; k++ { // 左 前缀和
			if k == 0 || l[k-1] < dp[i-1][k]+k {
				l[k] = dp[i-1][k] + k
			} else {
				l[k] = l[k-1]
			}
		}
		for k := n - 1; k >= 0; k-- { // 右 前缀和
			if k == n-1 || r[k+1] < dp[i-1][k]-k {
				r[k] = dp[i-1][k] - k
			} else {
				r[k] = r[k+1]
			}
		}
		for j, v := range arr { // dp
			dp[i][j] = l[j] + v - j
			if r[j]+v-j > dp[i][j] {
				dp[i][j] = r[j] + v + j
			}
		}
		fmt.Println(i, l, r, dp[i])
	}
	for _, v := range dp[m-1] {
		if v > ans {
			ans = v
		}
	}
	return int64(ans)
}

func getTwoMax(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	sum, max, rec := 0, 0, append([]int{}, nums1...)
	sort.Ints(rec)
	for k, v := range nums2 {
		sum += abs(nums1[k], v)

	}
	return sum - max
}

func abs(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

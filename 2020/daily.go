package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(reverseStr("abcd", 4))
}

// 863 二叉树中所有距离为 K 的结点
// https://leetcode-cn.com/problems/all-nodes-distance-k-in-binary-tree/
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	var path []*TreeNode
	var ans []int
	m := map[int]bool{}

	// 递归出path
	var dfs func(p *TreeNode, tmp []*TreeNode)
	dfs = func(p *TreeNode, tmp []*TreeNode) {
		if p == nil {
			return
		}

		var nt []*TreeNode
		nt = append(nt, p)
		nt = append(nt, tmp...)
		if p == target {
			path = nt
			return
		}

		dfs(p.Left, nt)
		dfs(p.Right, nt)
	}

	// 计算每个节点的自己点是否k=0
	var getK func(p *TreeNode, k int)
	getK = func(p *TreeNode, k int) {
		if k < 0 || p == nil || m[p.Val] {
			return
		}
		if k == 0 {
			ans = append(ans, p.Val)
			return
		}
		getK(p.Left, k-1)
		getK(p.Right, k-1)
	}

	dfs(root, []*TreeNode{})
	for i, v := range path {
		getK(v, k-i)
		m[v.Val] = true // 标记已经遍历过的父节点
	}

	return ans
}

// 1104. 二叉树寻路
// https://leetcode-cn.com/problems/path-in-zigzag-labelled-binary-tree/
func pathInZigZagTree(label int) []int {
	level, idx, ans := 0, 0, []int{label}
	for {
		if 1<<(level+1) > label {
			break
		}
		level++
	}
	if level%2 == 0 { // 偶 左
		idx = label - (1 << level)
	} else {
		idx = (1 << (level + 1)) - 1 - label
	}
	for i := level - 1; i >= 0; i-- {
		idx = idx / 2
		if i%2 == 0 { // 偶 左
			ans = append(ans, (1<<i)+idx)
		} else {
			ans = append(ans, (1<<(i+1))-1-idx)
		}
	}

	for i, n := 0, len(ans)-1; i <= n/2; i++ {
		ans[i], ans[n-i] = ans[n-i], ans[i]
	}

	return ans
}

// 743 网络延迟时间
// https://leetcode-cn.com/problems/network-delay-time/
// Dijkstra 迪杰斯特拉算法 有向图最短路径
// not ac
func networkDelayTime(times [][]int, n int, k int) int {
	m, s, u, dp := map[int]map[int]int{}, map[int]bool{k: true}, map[int]bool{}, map[int]int{}
	for i := 1; i <= n; i++ {
		dp[i], u[i] = -1, true
	}
	dp[k] = 0
	delete(u, k)

	for _, arr := range times {
		if m[arr[0]] == nil {
			m[arr[0]] = map[int]int{}
		}
		m[arr[0]][arr[1]] = arr[2]
	}

	for len(u) > 0 {
		for si, _ := range s {
			for i, v := range m[si] { // 更新U
				if dp[i] < 0 || dp[i] > dp[si]+v {
					dp[i] = dp[si] + v
				}
			}
		}
		minU, minUi := int(1e6), -1
		for ui, _ := range u {
			if dp[ui] > 0 && dp[ui] < minU {
				minU, minUi = dp[ui], ui
			}
		}
		if minUi == -1 {
			return -1
		}
		s[minUi] = true
		delete(u, minUi)
	}

	ans := 0
	for _, v := range dp {
		if v == -1 {
			return -1
		}
		if v > ans {
			ans = v
		}
	}
	return ans
}

// 611. 有效三角形的个数
// https://leetcode-cn.com/problems/valid-triangle-number/
// 双指针
func triangleNumber(nums []int) int {
	ans := 0
	sort.Ints(nums)
	nums = append(nums, 10000)
	for k := 0; k < len(nums)-3; k++ {
		for i, j := k+1, k+2; i < len(nums)-2; i++ {
			for {
				if nums[j] >= nums[i]+nums[k] {
					if j > i+1 {
						ans += j - i - 1
					}
					break
				}
				j++
			}
			if i == j {
				j++
			}
		}
	}
	return ans
}

// https://leetcode-cn.com/problems/find-eventual-safe-states/
// 802. 找到最终的安全状态
// 三色标记 | 拓扑排序
func eventualSafeNodes_(graph [][]int) []int {
	color := map[int]int{} // 0-未访问, 1-访问过, 2-安全
	var safe func(x int) bool
	safe = func(x int) bool {
		if len(graph[x]) == 0 || color[x] == 2 {
			return true
		}
		if color[x] == 1 {
			return false
		}
		color[x] = 1
		for _, v := range graph[x] {
			if !safe(v) {
				return false
			}
		}
		color[x] = 2
		return true
	}

	ans := []int{}
	for i := 0; i < len(graph); i++ {
		if safe(i) {
			ans = append(ans, i)
		}
	}
	return ans
}
func eventualSafeNodes(graph [][]int) []int {
	g := make([][]int, len(graph))
	inDeg := make([]int, len(graph))
	list := []int{}
	ans := []int{}

	for k, arr := range graph {
		inDeg[k] = len(arr) // 入度
		if len(arr) == 0 {
			list = append(list, k)
		}
		for _, v := range arr { // 反向图
			g[v] = append(g[v], k)
		}
	}
	for len(list) > 0 {
		for _, v := range g[list[0]] { // 拆边
			inDeg[v]--
			if inDeg[v] == 0 {
				list = append(list, v)
			}
		}
		list = list[1:]
	}
	for i, v := range inDeg {
		if v == 0 {
			ans = append(ans, i)
		}
	}

	return ans
}

func numberOfArithmeticSlices_(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	diff, ans, t := nums[1]-nums[0], 0, 0
	nums = append(nums, 3000)
	for i := 2; i < len(nums); i++ {
		if diff == nums[i]-nums[i-1] {
			t++
		} else {
			t, diff = 0, nums[i]-nums[i-1]
		}
		ans += t
	}
	return ans
}

// https://leetcode-cn.com/problems/shortest-path-visiting-all-nodes/
// 847. 访问所有节点的最短路径
// 图 广度优先
func shortestPathLength(graph [][]int) int {
	n := len(graph)
	type tuple struct{ u, mask, dist int }
	q := []tuple{}
	seen := []map[int]bool{}

	for i := 0; i < n; i++ {
		q = append(q, tuple{i, 1 << i, 0})
		seen = append(seen, map[int]bool{1 << i: true})
	}

	for {
		//t, q := q[0], q[1:]
		t := q[0]
		q = q[1:]
		if t.mask == (1<<n)-1 {
			return t.dist
		}
		for _, v := range graph[t.u] {
			mask := (1 << v) | t.mask
			if !seen[v][mask] {
				seen[v][mask] = true
				q = append(q, tuple{v, mask, t.dist + 1})
			}
		}
	}
}

// https://leetcode-cn.com/problems/minimum-total-space-wasted-with-k-resizing-operations/
func minSpaceWastedKResizing(nums []int, k int) int {
	n := len(nums)
	dp, g, w, sum := make([][]int, n), make([][]int, n), make([][]int, n), 0
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
		w[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 0; i-- {
		g[i][i], sum = nums[i], nums[i]
		for x := i - 1; x >= 0; x-- {
			sum += nums[x]
			g[x][i] = maxTwo(g[x+1][i], nums[x])
			w[x][i] = g[x][i]*(i-x+1) - sum
		}
	}

	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
		dp[i][0] = w[0][i]
		for j := 1; j <= k; j++ {
			dp[i][j] = w[1][i]
			for x := 1; x <= i-1; x++ {
				dp[i][j] = minTwo(dp[x][j-1]+w[x+1][i], dp[i][j])
			}
			dp[i][j] = minTwo(dp[i][j], dp[i][j-1])
		}
	}

	return dp[n-1][k]
}

func maxTwo(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minTwo(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// https://leetcode-cn.com/problems/finding-the-users-active-minutes/
func findingUsersActiveMinutes(logs [][]int, k int) []int {
	m, u, ans := map[int]map[int]bool{}, map[int]int{}, make([]int, k)
	for _, arr := range logs {
		if m[arr[0]] == nil {
			m[arr[0]] = map[int]bool{}
		}
		if !m[arr[0]][arr[1]] {
			m[arr[0]][arr[1]] = true
			u[arr[0]]++
		}
	}

	for _, num := range u {
		ans[num-1]++
	}

	return ans
}

// https://leetcode-cn.com/problems/arithmetic-slices-ii-subsequence/
func numberOfArithmeticSlices(nums []int) int {
	n, ans := len(nums), 0
	dp := make([]map[int]int, n)
	for i, v := range nums {
		dp[i] = map[int]int{}
		for j, x := range nums[:i] {
			ans += dp[j][v-x]
			dp[i][v-x] += dp[j][v-x] + 1
		}
	}
	return ans
}

// https://leetcode-cn.com/problems/maximum-score-from-removing-substrings/
func maximumGain(s string, x int, y int) int {
	a, b, ca, cb, ans := 0, 0, 'a', 'b', 0
	if x < y {
		ca, cb, x, y = 'b', 'a', y, x
	}
	s += "x"
	for _, v := range s {
		if v == ca || v == cb {
			if v == cb {
				if a > 0 {
					ans += x
					a--
				} else {
					b++
				}
			}
			if v == ca {
				a++
			}
		} else {
			ans += minTwo(a, b) * y
			a, b = 0, 0
		}
	}
	return ans
}

// https://leetcode-cn.com/problems/longest-palindromic-subsequence/solution/
func longestPalindromeSubseq(s string) int {
	n, arr := len(s), []byte(s)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1-i)
	}
	for i := 0; i < n; i++ {
		dp[1][i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			if arr[j] == arr[j+i-1] {
				dp[i][j] = dp[i-2][j+1] + 2
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j+1])
			}
		}
	}
	return dp[n][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func countArrangement(n int) int {
	f := map[int]int{0: 1}
	for mask := 1; mask < 1<<n; mask++ {
		k := oc(mask)
		for i := 0; i < n; i++ {
			if (mask & (1 << i)) > 0 { // 第i位是1
				pre := mask ^ 1<<i                // 第i位设为0
				if (i+1)%k == 0 || k%(i+1) == 0 { // 第k个数整除
					f[mask] += f[pre]
				}
			}
		}
	}
	return f[(1<<n)-1]
}

func oc(x int) (ans int) {
	for x > 0 {
		x -= x & -x
		ans++
	}
	return
}

func rearrangeArray(nums []int) []int {
	sort.Ints(nums)
	ans := []int{}
	for i := 0; i < len(nums)/2; i++ {
		ans = append(ans, nums[i], nums[len(nums)-1-i])
	}
	if len(nums)%2 == 1 {
		ans = append(ans, nums[len(nums)/2])
	}
	return ans
}

func wiggleSort(nums []int) {
	sort.Ints(nums)
	ans := []int{}
	for i := 0; i < len(nums)/2; i++ {
		ans = append(ans, nums[(len(nums)-1)/2-i], nums[len(nums)-1-i])
	}
	if len(nums)%2 == 1 {
		ans = append(ans, nums[0])
	}
	for i, v := range ans {
		nums[i] = v
	}
}

// https://leetcode-cn.com/problems/student-attendance-record-ii/
func checkRecord(n int) int {
	mod := int(1e9) + 7
	dp := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, 3) // 0-P 1-L 2-A
		for j := 0; j <= 2; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	dp[0][0][0] = 1
	dp[1][0][0] = 1
	dp[1][1][0] = 1
	dp[1][2][1] = 1

	for i := 2; i <= n; i++ {
		// P
		dp[i][0][0] = (dp[i-1][0][0] + dp[i-1][1][0]) % mod
		dp[i][0][1] = (dp[i-1][0][1] + dp[i-1][1][1] + dp[i-1][2][1]) % mod

		// L
		dp[i][1][0] = (dp[i-1][0][0] + dp[i-2][0][0]) % mod
		dp[i][1][1] = (dp[i-1][2][1] + dp[i-1][0][1] + dp[i-2][0][1] + dp[i-2][2][1]) % mod

		// A
		dp[i][2][1] = (dp[i-1][0][0] + dp[i-1][1][0]) % mod
	}

	ans := 0
	for _, arr := range dp[n] {
		for _, v := range arr {
			ans += v
		}
	}
	return ans
}

// hello
func reverseVowels(s string) string {
	arr, m := []byte(s), map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	i, j := 0, len(s)-1
	for {
		for !m[arr[i]] {
			if i == len(s)-1 {
				return string(arr)
			}
			i++
		}
		for !m[arr[j]] {
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return string(arr)
}

func reverseStr(s string, k int) string {
	arr, reverse := []byte(s), false
	for i := 0; i < len(s); i += k {
		reverse = !reverse
		if !reverse {
			continue
		}
		for j := 0; j%k != k/2 && j%k != (len(s)-i)/2 && i+j < len(s); j++ {
			r := i + k - 1 - j
			if i+k >= len(s) {
				r = len(s) - 1 - j
			}
			//fmt.Println(i+j, r)
			arr[i+j], arr[r] = arr[r], arr[i+j]
		}
	}
	return string(arr)
}

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	ans := 0
	for i, j := 0, len(people)-1; i <= j; j-- {
		if people[i]+people[j] <= limit {
			i++
		}
		ans++
	}
	return ans
}

type Solution struct {
	sum []int
}

func Constructor_(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}
	return Solution{w}
}

func (this *Solution) PickIndex() int {
	r := rand.Intn(this.sum[len(this.sum)-1]) + 1
	return sort.SearchInts(this.sum, r)
}

func corpFlightBookings(bookings [][]int, n int) []int {
	dp := make([]int, n+1)
	for _, arr := range bookings {
		dp[arr[0]-1] += arr[2]
		dp[arr[1]] -= arr[2]
	}
	for i := 1; i < n; i++ {
		dp[i] += dp[i-1]
	}
	return dp[:n]
}

func distinctSubseqII(s string) int {
	n := len(s)
	dp, last := make([]int, n+1), map[uint8]int{}
	dp[0] = 1
	for k := 1; k <= n; k++ {
		dp[k] = 2 * dp[k-1]
		if last[s[k-1]] > 0 {
			dp[k] -= dp[last[s[k-1]]-1]
		}
		dp[k] %= int(1e9 + 7)
		last[s[k-1]] = k
	}
	ans := dp[n] - 1
	if ans < 0 {
		ans += int(1e9 + 7)
	}
	return ans
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	a, b := head, head
	for i := 0; i < k; i++ {
		b = b.Next
	}
	for b != nil {
		a, b = a.Next, b.Next
	}
	return a
}

// 1d bin packing 背包不可解
func minSessions_(tasks []int, sessionTime int) (ans int) {
	sort.Ints(tasks)
	ans = get(tasks, sessionTime)
	for i := 0; i < len(tasks)/2; i++ {
		tasks[i], tasks[len(tasks)-1-i] = tasks[len(tasks)-1-i], tasks[i]
	}
	x := get(tasks, sessionTime)
	if x < ans {
		ans = x
	}
	return
}

func get(tasks []int, sessionTime int) (ans int) {
	mt := map[int]bool{0: true}
	for len(mt) > 0 {
		ans++
		mt = map[int]bool{}
		for k, _ := range tasks {
			mt[k] = true
		}

		// 背包
		dp, path := map[int]int{}, make([][]int, len(tasks))
		for k, v := range tasks {
			tmpI := k
			dp[k] = v
			path[k] = []int{}
			for i := 0; i < k; i++ {
				if dp[i]+v <= sessionTime && dp[k] < dp[i]+v {
					dp[k] = dp[i] + v
					tmpI = i
				}
			}
			path[k] = append(path[tmpI], k)
		}

		// 找出最多方案
		mv, mk := 0, 0
		for i := 0; i < len(tasks); i++ {
			if dp[i] > mv {
				mk, mv = i, dp[i]
			}
		}

		// 删除已装
		for _, v := range path[mk] {
			delete(mt, v)
		}

		// 重置tasks
		nt := []int{}
		for t, tv := range tasks {
			if mt[t] {
				nt = append(nt, tv)
			}
		}
		tasks = nt
	}
	return
}

func minSessions(tasks []int, sessionTime int) (ans int) {

	return
}

func numberOfUniqueGoodSubsequences(binary string) int {
	pre, dp, i, zero, mod := [2]int{-1, -1}, make([]int, len(binary)), 0, 0, int(1e9+7)
	for binary[i] != 49 {
		i++
		if i == len(binary) {
			return 1
		}
	}

	if i > 0 {
		zero = 1
	}
	dp[i], i = 1, i+1
	for i < len(binary) {
		v := binary[i] - 48
		if v == 0 {
			zero = 1
		}
		dp[i] = dp[i-1] * 2
		if pre[v] >= 0 {
			dp[i] -= dp[pre[v]-1]
		}
		if dp[i] < 0 {
			dp[i] += mod
		} else {
			dp[i] %= mod
		}
		pre[v] = i
		i++
	}
	return dp[len(binary)-1] + zero
}

func minimumOperations(leaves string) int {
	ans := 0
	if leaves[0] != 'r' {
		ans++
	}
	if leaves[len(leaves)-1] != 'r' {
		ans++
	}

	sr, sy := make([]int, len(leaves)), make([]int, len(leaves))
	sr[0], sy[0] = 1, 0
	for i := 1; i < len(leaves)-1; i++ {
		sr[i], sy[i] = sr[i-1], sy[i-1]
		if leaves[i] == 'r' {
			sr[i]++
		} else {
			sy[i]++
		}
	}
	return ans
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func chalkReplacer(chalk []int, k int) int {
	sum := 0
	for _, v := range chalk {
		sum += v
	}
	k = k % sum
	for i, v := range chalk {
		k -= v
		if k < 0 {
			return i
		}
	}
	return 0
}

func numberOfBoomerangs(points [][]int) int {
	n := len(points)
	ans := 0
	for _, cur := range points {
		m := map[int]int{}
		for j := 0; j < n; j++ {
			d := (points[j][1]-cur[1])*(points[j][1]-cur[1]) + (points[j][0]-cur[0])*(points[j][0]-cur[0])
			m[d]++
		}
		for _, v := range m {
			ans += v * (v - 1)
		}
	}
	return ans
}

func smallestMissingValueSubtree(parents []int, nums []int) []int {
	son, ans := map[int][]int{}, make([]int, len(parents))
	for i, v := range parents {
		if son[v] == nil {
			son[v] = []int{i}
			continue
		}
		son[v] = append(son[v], i)
	}

	var f func(i int) map[int]bool
	f = func(i int) map[int]bool {
		inSet, val := map[int]bool{}, 1
		for _, v := range son[i] {
			w := f(v)
			if len(w) > len(inSet) {
				w, inSet = inSet, w
			}
			if ans[v] > val {
				val = ans[v]
			}
			for x, _ := range w {
				inSet[x] = true
			}
		}
		inSet[nums[i]] = true

		for inSet[val] == true {
			val++
		}
		ans[i] = val

		return inSet
	}

	f(0)
	return ans
}

func findPeakElement(nums []int) int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i] > nums[j] {
			j = (i + j) / 2
		} else {
			i = (i + j) / 2
		}
	}

}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		m := map[byte]bool{}
		for _, v := range board[i] {
			if v == '.' {
				continue
			}
			if m[v] {
				return false
			}
			m[v] = true
		}
	}
	for i := 0; i < 9; i++ {
		m := map[byte]bool{}
		for j := 0; j < 9; j++ {
			if board[j][i] == '.' {
				continue
			}
			if m[board[j][i]] {
				return false
			}
			m[board[j][i]] = true
		}
	}
	i, j := 0, 0
	for k := 0; k < 9; k++ {
		m := map[byte]bool{}
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				cur := board[i+x][j+y]
				if cur == '.' {
					continue
				}
				if m[cur] {
					return false
				}
				m[cur] = true
			}
		}
		j += 3
		if j == 9 {
			i, j = i+3, 0
		}
	}
	return true
}

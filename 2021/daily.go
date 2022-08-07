package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

func reversePrefix(word string, ch byte) string {
	w := []byte(word)
	for i := 0; i < len(w); i++ {
		if w[i] == ch {
			for j := 0; j <= i/2; j++ {
				w[j], w[i-j] = w[i-j], w[j]
			}
			break
		}
	}
	return string(w)
}

func interchangeableRectangles(rectangles [][]int) int64 {
	m := map[float64]int64{}
	for _, v := range rectangles {
		m[(float64(v[0])/float64(v[1]))]++
	}
	ans := int64(0)
	for _, v := range m {
		ans += v * (v - 1) / 2
	}
	return ans
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

func checkValidString(s string) bool {
	l, star, cl := 0, 0, 0
	for _, v := range s {
		if v == '*' {
			star++
			cl++ // 可能的右
		} else if v == '(' {
			l++
		} else if v == ')' {
			if l > 0 {
				l--
			} else if star > 0 {
				star--
			} else {
				return false
			}
		}
		if cl > l { // 在L之前出现的删除
			cl = l
		}
	}
	return l <= cl
}

func getSeq(s []byte) {
	for i := 1; i < 1<<len(s)-1; i++ {
		for j := 1; j < len(s); j++ {

		}
	}
}
func getTwo(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Solution struct {
	sum []int
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

func isPrefixString(s string, words []string) bool {
	idx := 0
	for _, str := range words {
		if idx+len(str) > len(s) {
			return false
		}
		for i := 0; i < len(str); i++ {
			if str[i] != s[idx] {
				return false
			}
			idx++
		}
		if idx == len(s) {
			return true
		}
	}
	return idx == len(s)
}

func get(v int, z *[]int) int {
	ans := 0
	for v != 1 {
		for i, x := range *z {
			if v%x == 0 {
				ans |= 1 << i
				v /= x
			}
		}
	}
	return ans
}

func countSubstrings(s string) int {
	dp, ans := map[int]map[int]bool{}, 0
	for i := 1; i <= len(s); i++ {
		dp[i] = map[int]bool{}
	}
	for k, _ := range s {
		dp[1][k] = true
		ans++
	}
	for i := 0; i <= len(s)-2; i++ {
		if s[i] == s[i+1] {
			dp[2][i] = true
			ans++
		}
	}
	for l := 3; l <= len(s); l++ {
		for i := 0; i <= len(s)-l; i++ {
			if dp[l-2][i+1] && s[i] == s[i+l-1] {
				dp[l][i] = true
				ans++
			}
		}
	}
	return ans
}

func maxProduct(s string) (max int) {
	sa, n, max := []byte(s), len(s), 0
	for i := 1; i < 1<<n-1; i++ {
		cur, supp := []byte{}, []byte{}
		for j := 0; j < n; j++ {
			if i&(1<<j) > 0 {
				cur = append(cur, sa[j])
			} else {
				supp = append(supp, sa[j]) // 补集
			}
		}
		if check(cur) {
			x := len(cur) * maxSeq(string(supp))
			max = twoMax(x, max)
		}
	}
	return
}
func check(s []byte) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
func maxSeq(s string) int {
	max := 1
	for i := 1; i < 1<<len(s)-1; i++ {
		tmp := []byte{}
		for j := 1; j < len(s); j++ {
			if i&(1<<j) > 0 {
				tmp = append(tmp, s[j])
			}
		}
		if check(tmp) {
			max = twoMax(max, len(tmp))
		}
	}
	return max
}
func maxSeq_(s string) int {
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
				dp[i][j] = twoMax(dp[i-1][j], dp[i-1][j+1])
			}
		}
	}
	return dp[n][0]
}
func twoMax(a, b int) int {
	if a > b {
		return a
	}
	return b
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

func splitListToParts(head *ListNode, k int) []*ListNode {
	ans, n := make([]*ListNode, k), 0
	for x := head; x != nil; x = x.Next {
		n++
	}

	for j, p, l := 0, head, n/k+1; j < k && p != nil; j++ {
		ans[j] = p
		if n%k == j {
			l--
		}
		for i := 1; i < l && p != nil; i++ {
			p = p.Next
		}
		p, p.Next = p.Next, nil
	}

	return ans
}

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	sum := (ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1)
	return sum - sec(ax1, ax2, bx1, bx2)*sec(ay1, ay2, by1, by2)
}

// 两线段交集长度
func sec(a1, b1, a2, b2 int) int {
	if a1 > a2 {
		return sec(a2, b2, a1, b1)
	}
	if a2 > b1 {
		return 0
	}
	if b2 < b1 {
		return b2 - a2
	}
	return b1 - a2
}

type SummaryRanges struct {
	m   map[int]bool
	res [][]int
}

func Constructor__() SummaryRanges {
	return SummaryRanges{
		m:   map[int]bool{},
		res: [][]int{},
	}
}

func (this *SummaryRanges) AddNum(val int) { // [0,1E4]
	if this.m[val] == true {
		return
	}
	this.m[val] = true

	if this.m[val+1] == false && this.m[val-1] == false {
		// 二分插入
		pos := binaryRight(this.res, val)
		if pos == len(this.res) {
			this.res = append(this.res, []int{val, val})
		} else {
			x := append([][]int{}, this.res[pos:]...)
			this.res = append(append(this.res[:pos], []int{val, val}), x...)
		}
	} else if this.m[val+1] == true && this.m[val-1] == false {
		// 合到右边
		pos := binaryRight(this.res, val)
		this.res[pos][0] = val
	} else if this.m[val+1] == false && this.m[val-1] == true {
		// 合到左边
		pos := binaryRight(this.res, val)
		this.res[pos-1][1] = val
	} else {
		// 合并左右
		pos := binaryRight(this.res, val)
		this.res[pos-1][1] = this.res[pos][1]
		if pos == len(this.res)-1 {
			this.res = this.res[:len(this.res)-1]
		} else {
			this.res = append(this.res[:pos], this.res[pos+1:]...)
		}
	}
}

func binaryRight(arr [][]int, v int) int {
	l, r, pos := 0, len(arr)-1, 0
	for l <= r {
		mid := (l + r) / 2
		if v < arr[mid][1] {
			pos, r = mid, mid-1
		} else {
			pos, l = mid+1, mid+1
		}
	}
	return pos
}

func (this *SummaryRanges) GetIntervals() [][]int {
	return this.res
}

type Hi []int

func (h Hi) Len() int            { return len(h) }
func (h Hi) Less(i, j int) bool  { return h[i] > h[j] }
func (h Hi) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Hi) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *Hi) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

type Hi2 []int

func (h Hi2) Len() int            { return len(h) }
func (h Hi2) Less(i, j int) bool  { return h[i] < h[j] }
func (h Hi2) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Hi2) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *Hi2) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

type StockPrice struct {
	timePrice map[int]int
	bigH      *Hi
	smallH    *Hi2
	priceMap  map[int]int
	mxTs      int
}

func Constructor() StockPrice {
	return StockPrice{
		timePrice: map[int]int{},
		bigH:      &Hi{},
		smallH:    &Hi2{},
		priceMap:  map[int]int{},
		mxTs:      0,
	}
}

func (this *StockPrice) Update(timestamp int, price int) {
	if this.timePrice[timestamp] > 0 {
		this.priceMap[this.timePrice[timestamp]]--
	}

	this.timePrice[timestamp] = price
	this.priceMap[price]++

	if timestamp > this.mxTs {
		this.mxTs = timestamp
	}
	heap.Push(this.bigH, price)
	heap.Push(this.smallH, price)
}

func (this *StockPrice) Current() int {
	return this.timePrice[this.mxTs]
}

func (this *StockPrice) Maximum() int {
	h := this.bigH
	for {
		x := heap.Pop(h).(int)
		if this.priceMap[x] > 0 {
			heap.Push(this.bigH, x)
			return x
		}
	}
}

func (this *StockPrice) Minimum() int {
	h := this.smallH
	for {
		x := heap.Pop(h).(int)
		if this.priceMap[x] > 0 {
			heap.Push(this.smallH, x)
			return x
		}
	}
}

func main() {
	fmt.Println(secondMinimum(11, [][]int{{9, 1}, {9, 2}, {2, 3}, {10, 3}, {7, 1}, {4, 3}, {6, 2}, {8, 9}, {5, 2}, {6, 11}, {9, 7}, {4, 2}, {4, 10}, {1, 10}, {2, 10}, {1, 4}, {5, 7}, {10, 8}, {5, 9}, {7, 3}, {7, 10}, {6, 5}}, 7, 2))
}

func secondMinimum(n int, edges [][]int, time int, change int) int {
	sub := make([][]int, n+1)
	dp := map[int]int{1: 0}
	for _, e := range edges {
		sub[e[0]] = append(sub[e[0]], e[1])
		sub[e[1]] = append(sub[e[1]], e[0])
	}

	// 找到最短路
	next := map[int]bool{}
	cur := map[int]bool{1: true}
	deep, flag := 0, true
	for flag {
		deep++
		next = map[int]bool{}
		for e, _ := range cur {
			for _, v := range sub[e] {
				if v == n {
					flag = false
				}
				if dp[v] == 0 {
					dp[v] = deep
				}
				if deep == dp[v] || deep == dp[v]+1 {
					next[v] = true
				}
			}
		}
		cur = next
	}

	// 再执行一步
	flag = false
	for e, _ := range cur {
		if flag {
			break
		}
		for _, v := range sub[e] {
			if v == n {
				deep++
				flag = true
				break
			}
		}
	}
	if !flag {
		deep += 2
	}

	ans := 0
	for i := 1; i <= deep-1; i++ {
		ans += time
		if (ans/change)%2 == 1 {
			ans += change - ans%change
		}
	}

	return ans + time
}

package main

import (
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
	color := map[int]int{}
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
			if color[v] == 1 || !safe(v) {
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

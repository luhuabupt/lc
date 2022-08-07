package main

func volunteerDeployment(finalCnt []int, totalNum int64, edges [][]int, plans [][]int) []int {
	n := len(finalCnt) + 1
	a := make([][]int, n)
	a[0] = []int{1, 0} //初始 X
	for i, v := range finalCnt {
		a[i+1] = []int{0, v}
	}

	// 相邻
	nb := map[int][]int{}
	for _, arr := range edges {
		if nb[arr[0]] == nil {
			nb[arr[0]] = []int{}
		}
		if nb[arr[1]] == nil {
			nb[arr[1]] = []int{}
		}
		nb[arr[0]] = append(nb[arr[0]], arr[1])
		nb[arr[1]] = append(nb[arr[1]], arr[0])
	}

	for i := len(plans) - 1; i >= 0; i-- {
		op, idx := plans[i][0], plans[i][1]
		if op == 0 {
			a[idx][0] *= 2
			a[idx][1] *= 2
		} else if op == 1 {
			for _, nbs := range nb[idx] {
				a[nbs][0] -= a[idx][0]
				a[nbs][1] -= a[idx][1]
			}
		} else if op == 2 {
			for _, nbs := range nb[idx] {
				a[nbs][0] += a[idx][0]
				a[nbs][1] += a[idx][1]
			}
		}
	}

	x, y := int64(0), int64(0)
	for _, v := range a {
		x += int64(v[0])
		y += int64(v[1])
	}
	base := int((totalNum - y) / x)

	ans := make([]int, n)
	for i, v := range a {
		ans[i] = base*v[0] + v[1]
	}

	return ans
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func numColor(root *TreeNode) int {
	m := map[int]bool{}
	var dfs func(p *TreeNode)
	dfs = func(p *TreeNode) {
		if p == nil {
			return
		}
		m[p.Val] = true
		dfs(p.Left)
		dfs(p.Right)
	}

	dfs(root)
	return len(m)
}

func securityCheck(capacities []int, k int) int {
	n := len(capacities)
	mod := 1000000007

	pre := map[int]int{}
	pre[k]++
	pre[k-capacities[0]+1]++

	cur := pre
	for i := 1; i < n; i++ {
		cur = map[int]int{}
		for j, v := range pre {
			cur[j] = (cur[j] + v) % mod
			cur[j-capacities[i]+1] = (cur[j-capacities[i]+1] + v) % mod
		}
		pre = cur
	}
	return cur[0]
}

func bicycleYard(position []int, terrain [][]int, obstacle [][]int) [][]int {
	ans := [][]int{}
	am := map[int]map[int]map[int]bool{}
	m := len(terrain)
	n := len(terrain[0])
	x := []int{-1, 0, 0, 1}
	y := []int{0, -1, 1, 0}

	var dfs func(i, j, v int)
	dfs = func(i, j, v int) {
		for k := 0; k < 8; k++ {
			xi, yi := i+x[k], j+y[k]
			if xi >= 0 && xi < m && yi >= 0 && yi < n {
				cv := v + terrain[i][j] - terrain[xi][yi] - obstacle[xi][yi]
				if cv <= 0 {
					continue
				}
				if am[xi] == nil {
					am[xi] = map[int]map[int]bool{}
				}
				if am[xi][yi] == nil {
					am[xi][yi] = map[int]bool{}
				}
				if am[xi][yi][cv] {
					continue
				}
				am[xi][yi][cv] = true
				dfs(xi, yi, cv)
			}
		}
	}

	dfs(position[0], position[1], 1)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == position[0] && j == position[1] {
				continue
			}
			if am[i][j][1] {
				ans = append(ans, []int{i, j})
			}
		}
	}

	return ans
}

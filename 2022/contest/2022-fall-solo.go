package main

import "fmt"

func main() {
	fmt.Println(ballGame(4, []string{"..E.", ".EOW", "..W."}))
	fmt.Println(ballGame(5, []string{".....", "..E..", ".WO..", "....."}))
	fmt.Println(ballGame(3, []string{".....", "....O", "....O", "....."}))
	fmt.Println(ballGame(1, []string{".....", "..O..", "....."})) // [[0 2] [2 2]]
	fmt.Println(ballGame(76, []string{
		"E......O..",
		"E.........",
		"W..E...EW.",
		"EE.OE.WWWO",
		"O.WEOEWWW.",
		".OW..W....",
		"W...EW.WE.",
		".E.W...OW.",
		"OW...W.EEO",
		"......W.W."})) // [[0,2],[0,4],[0,5],[1,9],[2,9],[4,9],[5,0],[5,9],[7,0],[7,9],[9,1],[9,2],[9,3]]
	fmt.Println(ballGame(69, []string{
		"W.W.WE..",
		".WWWEW..",
		"EWW.WE.E",
		"E.W.E.E.",
		".OEOO.EO",
		"WE.WOE.W",
		"WW...E..",
		".WEWO..O",
		"E....E..",
		".OWE...."}))
	// [[0,3],[0,6],[1,7],[4,0],[6,7],[9,4],[9,6]]

	fmt.Println(ballGame(41, []string{
		"E...W..WW",
		".E...O...",
		"...WO...W",
		"..OWW.O..",
		".W.WO.W.E",
		"O..O.W...",
		".OO...W..",
		"..EW.WEE."})) //[[0,2],[0,3],[0,5],[0,6],[1,0],[1,8],[3,0],[3,8],[4,0],[6,0],[7,1],[7,4]]
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//[1,0,1,1,1,0,null,null,1] 2
func closeLampInTree(root *TreeNode) int {
	dp := map[*TreeNode][]int{}

	var min func(a, b int) int
	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var xv func(p *TreeNode, s int) int
	xv = func(p *TreeNode, s int) int {
		if p == nil {
			return 0
		}
		if p.Val == s {
			return 0
		}
		return 1
	}

	var cal func(p *TreeNode, s int) int
	cal = func(p *TreeNode, s int) int {
		res := 0
		if p == nil {
			return 0
		}
		if p.Left != nil {
			res += dp[p.Left][s]
		}
		if p.Right != nil {
			res += dp[p.Right][s]
		}
		return res
	}

	var gen func(p *TreeNode, s int) int
	gen = func(p *TreeNode, s int) int {
		if p == nil {
			return 0
		}
		if p.Left != nil && p.Right != nil {
			if p.Left.Val == p.Right.Val && p.Val == p.Left.Val {
				if s == p.Val {
					return 0
				} else {
					return 1
				}
			} else {
				c := [2]int{}
				c[p.Left.Val]++
				c[p.Right.Val]++
				c[p.Val]++
				if c[s] == 2 {
					return 1
				} else {
					return 2
				}
			}
		}
		if p.Left == nil {
			if p.Right.Val == p.Val {
				if p.Val == s {
					return 0
				} else {
					return 1
				}
			} else {
				return 1
			}
		}
		if p.Right == nil {
			if p.Left.Val == p.Val {
				if p.Val == s {
					return 0
				} else {
					return 1
				}
			} else {
				return 1
			}
		}
		return 0
	}

	var dfs func(p *TreeNode)
	dfs = func(p *TreeNode) {
		if p == nil {
			return
		}
		dfs(p.Left)
		dfs(p.Right)
		dp[p] = make([]int, 2)

		if p.Left == nil && p.Right == nil {
			if p.Val == 0 {
				dp[p] = []int{0, 1}
			} else {
				dp[p] = []int{1, 0}
			}
			return
		}

		dp[p][1] = 1e6
		dp[p][1] = min(dp[p][1], gen(p, 1)+cal(p.Left, 1)+cal(p.Right, 1))
		dp[p][1] = min(dp[p][1], gen(p, 0)+cal(p.Left, 0)+cal(p.Right, 0)+1)
		dp[p][1] = min(dp[p][1], cal(p, 1)+xv(p, 1))
		dp[p][1] = min(dp[p][1], cal(p, 0)+xv(p, 0)+1)
		dp[p][1] = min(dp[p][1], cal(p.Left, 1)+cal(p.Right, 0)) // zuo

		dp[p][0] = 1e6
		dp[p][0] = min(dp[p][0], gen(p, 1)+cal(p.Left, 1)+cal(p.Right, 1)+1)
		dp[p][0] = min(dp[p][0], gen(p, 0)+cal(p.Left, 0)+cal(p.Right, 0))
		dp[p][0] = min(dp[p][0], cal(p, 1)+xv(p, 1)+1)
		dp[p][0] = min(dp[p][0], cal(p, 0)+xv(p, 0))
	}

	dfs(root)

	return dp[root][0]
}

func ballGame(num int, g []string) [][]int {
	m, n := len(g), len(g[0])
	dp := make([][][4]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][4]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = [4]int{-2, -2, -2, -2}
		}
	}

	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var dfs func(i, j, d int) int
	dfs = func(i, j, d int) int {
		//fmt.Println(i, j, d, string(g[i][j]))
		pd := d
		if dp[i][j][d] > 0 {
			return dp[i][j][d]
		}
		if g[i][j] == '.' {

		} else if g[i][j] == 'E' {
			d++
			d %= 4
		} else if g[i][j] == 'W' {
			d += 3
			d %= 4
		} else if g[i][j] == 'O' {
			dp[i][j][pd] = 0
			return 0
		}
		x := i + dir[d][0]
		y := j + dir[d][1]
		if x < 0 || x >= m || y < 0 || y >= n {
			dp[i][j][pd] = -1
			return -1
		}
		nx := dfs(x, y, d)
		if nx >= 0 {
			nx++
		}
		dp[i][j][pd] = nx
		return nx
	}

	for i := 1; i < m-1; i++ {
		if g[i][0] != '.' {
			continue
		}
		dfs(i, 0, 0)
	}
	for i := 1; i < m-1; i++ {
		if g[i][n-1] != '.' {
			continue
		}
		dfs(i, n-1, 2)
	}
	for j := 1; j < n-1; j++ {
		if g[0][j] != '.' {
			continue
		}
		dfs(0, j, 1)
	}
	for j := 1; j < n-1; j++ {
		if g[m-1][j] != '.' {
			continue
		}
		dfs(m-1, j, 3)
	}

	ans := [][]int{}
	ed := func(i, j, k int) {
		if g[i][j] != '.' {
			return
		}

		v := dp[i][j][k]
		if v > 0 && v <= num {
			ans = append(ans, []int{i, j})
		}
	}
	for i := 1; i < m-1; i++ {
		ed(i, 0, 0)
	}
	for i := 1; i < m-1; i++ {
		ed(i, n-1, 2)
	}
	for j := 1; j < n-1; j++ {
		ed(0, j, 1)
	}
	for j := 1; j < n-1; j++ {
		ed(m-1, j, 3)
	}

	return ans
}

func transportationHub(path [][]int) int {
	in := make([]map[int]bool, 1001)
	out := make([]int, 1001)
	h := map[int]bool{}

	for _, v := range path {
		if in[v[1]] == nil {
			in[v[1]] = map[int]bool{}
		}
		in[v[1]][v[0]] = true

		out[v[0]]++
		h[v[0]] = true
		h[v[1]] = true
	}

	fmt.Println(in)
	fmt.Println(out)
	fmt.Println(h)

	for i := 0; i < 1001; i++ {
		if out[i] > 0 {
			continue
		}
		if len(in[i]) == len(h) {
			return i
		}
	}

	return -1
}

func temperatureTrend(a []int, b []int) int {
	n := len(a)
	d := make([]int, n)
	d2 := make([]int, n)
	for i := 1; i < n; i++ {
		if a[i] > a[i-1] {
			d[i] = 1
		} else if a[i] == a[i-1] {
			d[i] = 0
		} else {
			d[i] = -1
		}

		if b[i] > b[i-1] {
			d2[i] = 1
		} else if b[i] == b[i-1] {
			d2[i] = 0
		} else {
			d2[i] = -1
		}
	}

	t := 0
	ans := 0
	for i := 1; i < n; i++ {
		if d[i] == d2[i] {
			t++
		} else {
			t = 0
		}
		if t > ans {
			ans = t
		}
	}
	return ans
}

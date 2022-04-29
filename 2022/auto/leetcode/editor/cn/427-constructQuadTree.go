package main

func main() {

}

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

//leetcode submit region begin(Prohibit modification and deletion)

func construct(g [][]int) *Node {
	n := len(g)
	eq := func(a, b, c, d, e int) bool {
		return a == b && b == c && c == d && d == e
	}
	get := func(x, y int) [][]int {
		ans := make([][]int, n/2)
		for i := 0; i < n/2; i++ {
			ans[i] = make([]int, n/2)
			for j := 0; j < n/2; j++ {
				ans[i][j] = g[x+i][y+j]
			}
		}
		return ans
	}

	p := &Node{}
	p.Val = g[0][0] == 1

	for i := 0; i < n/2; i++ {
		for j := 0; j < n/2; j++ {
			if !eq(g[0][0], g[i][j], g[i][j+n/2], g[i+n/2][j], g[i+n/2][j+n/2]) {
				p.IsLeaf = false
				p.TopLeft = construct(get(0, 0))
				p.TopRight = construct(get(0, n/2))
				p.BottomLeft = construct(get(n/2, 0))
				p.BottomRight = construct(get(n/2, n/2))

				return p
			}
		}
	}

	p.IsLeaf = true
	return p
}

//leetcode submit region end(Prohibit modification and deletion)

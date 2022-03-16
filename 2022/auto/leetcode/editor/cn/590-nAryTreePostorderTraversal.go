package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

//leetcode submit region begin(Prohibit modification and deletion)

func postorder(root *Node) []int {
	ans := []int{}
	var dfs func(p *Node)
	dfs = func(p *Node) {
		if p == nil {
			return
		}
		for _, x := range p.Children {
			dfs(x)
		}
		ans = append(ans, p.Val)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

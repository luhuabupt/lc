package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	st := []*Node{root}
	ans := [][]int{}
	if root == nil {
		return ans
	}
	for len(st) > 0 {
		ns := []*Node{}
		one := []int{}
		for _, v := range st {
			one = append(one, v.Val)
			for _, x := range v.Children {
				ns = append(ns, x)
			}
		}

		ans = append(ans, one)
		st = ns
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

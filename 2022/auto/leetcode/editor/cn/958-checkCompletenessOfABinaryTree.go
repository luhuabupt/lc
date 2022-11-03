package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCompleteTree(root *TreeNode) bool {
	st := []*TreeNode{root}
	end := false
	for len(st) > 0 {
		ns := []*TreeNode{}
		for _, x := range st {
			if x == nil {
				end = true
				continue
			}
			if end {
				return false
			}
			ns = append(ns, x.Left, x.Right)
		}
		st = ns
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

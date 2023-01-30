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
func bstFromPreorder(a []int) *TreeNode {
	if len(a) == 0 {
		return nil
	}

	h := &TreeNode{a[0], nil, nil}
	for i, v := range a {
		if v > a[0] {
			h.Left = bstFromPreorder(a[1:i])
			h.Right = bstFromPreorder(a[i:])
			return h
		}
	}
	h.Left = bstFromPreorder(a[1:])

	return h
}

//leetcode submit region end(Prohibit modification and deletion)

package main

import "fmt"

// construct-binary-tree-from-preorder-and-inorder-traversal 从前序与中序遍历序列构造二叉树  2022-12-26 16:47:02
func main() {
	fmt.Println("")
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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(p []int, m []int) *TreeNode {
	if len(p) == 0 {
		return nil
	}
	h := &TreeNode{p[0], nil, nil}

	for i, v := range m {
		if v == p[0] {
			h.Left = buildTree(p[1:i+1], m[:i])
			h.Right = buildTree(p[i+1:], m[i+1:])
			break
		}
	}

	return h
}

//leetcode submit region end(Prohibit modification and deletion)

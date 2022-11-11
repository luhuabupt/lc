package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	p := initBinaryTree("[1, 2, null, 3, 4]")
	fmt.Println(minCameraCover(p))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func initBinaryTree(a string) *TreeNode {
	arr := strings.Split(a[1:len(a)-1], ", ")
	if len(arr) == 0 {
		return nil
	}

	iv, _ := strconv.Atoi(arr[0])
	head := &TreeNode{iv, nil, nil}
	st := []*TreeNode{head}
	sti := []int{-1}
	for i, v := range arr {
		if v != "null" {
			iv, _ := strconv.Atoi(v)
			p := st[i]
			if sti[i] == 0 {
				st[i].Left = &TreeNode{}
				p = st[i].Left
			} else if sti[i] == 1 {
				st[i].Right = &TreeNode{}
				p = st[i].Right
			}
			p.Val = iv
			st = append(st, p, p)
			sti = append(sti, 0, 1)
		}
	}

	return head
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

func minCameraCover(root *TreeNode) int {
	min := func(a ...int) int {
		sort.Ints(a)
		return a[0]
	}

	// 0-n 1-y 2-n&need
	var dfs func(p *TreeNode) []int
	dfs = func(p *TreeNode) []int {
		r := []int{1e9, 1, 0}
		if p.Left != nil && p.Right != nil {
			x := dfs(p.Left)
			y := dfs(p.Right)
			r[0] = min(x[1]+y[1], x[1]+y[0], y[1]+x[0])
			r[1] = min(x[0], x[1], x[2]) + min(y[0], y[1], y[2]) + 1
			r[2] = x[0] + y[0]
		} else if p.Left != nil || p.Right != nil {
			if p.Left == nil {
				p = p.Right
			} else {
				p = p.Left
			}
			x := dfs(p)
			r[0] = x[1]
			r[1] = min(x[0], x[1], x[2]) + 1
			r[2] = x[0]
		}
		return r
	}

	ans := dfs(root)
	return min(ans[0], ans[1])
}

//leetcode submit region end(Prohibit modification and deletion)

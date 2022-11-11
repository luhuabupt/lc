package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := "1, 2, null, 3, null, 4, null, null, 5"
	r := initBinaryTree(a)
	var dfs func(p *TreeNode)
	dfs = func(p *TreeNode) {
		if p == nil {
			return
		}
		fmt.Println(p.Val)
		dfs(p.Left)
		dfs(p.Right)
	}
	dfs(r)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func initBinaryTree(a string) *TreeNode {
	arr := strings.Split(a, ", ")
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

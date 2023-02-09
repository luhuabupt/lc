package main

import "fmt"

func main() {
	fmt.Println(" ")
}

// To execute Go code, please declare a func main() in a package "main"

// The TestCase is shown below
// Input : 1 2
// Output : 3

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}

func binTree(pre []int, mid []int) *TreeNode {
	if len(pre) == 1 {
		return &TreeNode{nil, nil, pre[0]}
	}

	cur := TreeNode{nil, nil, pre[0]}
	for i, v := range mid {
		if v == pre[0] {
			if i == len(mid)-1 {
				cur.Left = binTree(pre[1:1+i], mid[:i])
			} else {
				cur.Left = binTree(pre[1:1+i], mid[:i])
				cur.Right = binTree(pre[1+i:], mid[i+1:])
			}
			break
		}
	}

	return &cur
}

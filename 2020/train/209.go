package main

import "sort"

// 存在一个数 x ，使得 nums 中恰好有 x 个元素 大于或者等于 x
func specialArray(nums []int) int {
	sort.Ints(nums)
	big := map[int]int{}
	for k, v := range nums {
		if big[v] == 0 {
			big[v] = len(nums) - k
		}
	}
	for i := 1000; i > 0; i-- {
		if big[i] == 0 {
			big[i] = big[i+1]
		}
	}
	for i := 1; i < 1000; i++ {
		if big[i] == i {
			return i
		}
	}
	return -1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	if root.Val&1 == 0 {
		return false
	}

	stack := []*TreeNode{root}
	ns := []*TreeNode{}
	deep := 1
	pre := 0

	for len(stack) != 0 {
		deep ^= 1
		pre = 0
		for _, v := range stack {
			if v.Val%2 != deep {
				return false
			}

			if pre > 0 {
				// 奇数行 偶数 递减  || 偶数行 奇数 递增
				if deep == 0 && v.Val >= pre || deep == 1 && v.Val <= pre {
					return false
				}
			}
			pre = v.Val

			if v.Left != nil {
				ns = append(ns, v.Left)
			}
			if v.Right != nil {
				ns = append(ns, v.Right)
			}
		}
		stack = ns
		ns = []*TreeNode{}
	}
	return true
}

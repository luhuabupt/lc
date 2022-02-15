package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	s := Constructor()
	p := s.deserialize("[5, 4, 7, 3, null, 2, null, -1, null, 9]")
	fmt.Println(s.serialize(p))
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

type Codec struct {
	a []int
}

// 二叉树序列化
func Constructor() Codec {
	return Codec{[]int{}}
}

// Serializes a tree to a single string.
func (t *Codec) serialize(p *TreeNode) string {
	st := []*TreeNode{p}
	ans := []byte{'['}
	for len(st) > 0 {
		cur := st[0]
		st = st[1:]
		if cur == nil {
			ans = append(ans, "null, "...)
		} else {
			st = append(st, cur.Left)
			st = append(st, cur.Right)
			ans = append(ans, strconv.Itoa(cur.Val)...)
			ans = append(ans, ',', ' ')
		}
	}
	ans = ans[:len(ans)-2]
	return string(append(ans, ']'))
}

// Deserializes your encoded data to tree.
func (t *Codec) deserialize(data string) *TreeNode {
	f := func(v string) *TreeNode {
		if v == "null" {
			return nil
		}
		x, _ := strconv.Atoi(v)
		return &TreeNode{x, nil, nil}
	}

	a := strings.Split(data[1:len(data)-1], ", ")
	at := []*TreeNode{}
	if a[0] == "null" {
		at = append(at, nil)
	} else {
		at = append(at, f(a[0]))
	}

	i, j := 0, 1
	for j < len(a) {
		p := at[i]
		if p != nil {
			p.Left = f(a[j])
			p.Right = nil
			if j+1 < len(a) {
				p.Right = f(a[j+1])
			}
			at = append(at, p.Left, p.Right)
			j += 2
		}

		i++
	}

	return at[0]
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
//leetcode submit region end(Prohibit modification and deletion)

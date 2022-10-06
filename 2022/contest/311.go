package main

import "fmt"

func main() {
	fmt.Println(sumPrefixScores([]string{"abc", "ab", "bc", "b"}))
}

func longestContinuousSubstring(s string) int {
	t := 0
	ans := 0
	p := ' '
	for _, v := range s {
		if (v-'a'+1)%26 == p-'a' {
			t++
		} else {
			t = 1
		}
		p = v
		if t > ans {
			ans = t
		}
	}

	return ans
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	st := []*TreeNode{root}
	d := 0
	for len(st) > 0 {
		ns := []*TreeNode{}
		for _, x := range st {
			if x == nil {
				return root
			}
			ns = append(ns, x.Left, x.Right)
		}
		if d%2 == 0 && ns[0] != nil {
			for i := 0; i < len(ns)/2; i++ {
				//fmt.Println(d, i)
				ns[i].Val, ns[len(ns)-1-i].Val = ns[len(ns)-1-i].Val, ns[i].Val
			}
		}
		d++
		st = ns
	}

	return root
}

func sumPrefixScores(words []string) []int {
	n := len(words)
	ans := make([]int, n)

	type dt struct {
		sum int
		sub [26]*dt
	}

	head := &dt{}
	for _, s := range words {
		p := head
		for _, v := range s {
			if p.sub[v-'a'] == nil {
				p.sub[v-'a'] = &dt{}
			}
			p = p.sub[v-'a']
			p.sum++
		}
	}

	//fmt.Println(head)
	//fmt.Println(head.sub[0])
	//fmt.Println(head.sub[1])

	for i, s := range words {
		p := head
		for _, v := range s {
			p = p.sub[v-'a']
			ans[i] += p.sum
		}
	}

	return ans
}

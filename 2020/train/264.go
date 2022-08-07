package main

import "fmt"

func main() {
	fmt.Println(nextBeautifulNumber(3000))
}

func countValidWords(sentence string) int {
	s := []byte(sentence)
	s = append(s, ' ')
	n := len(s)
	ans := 0
	pre := []byte{}
	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			flag := true
			cnt := 0
			if len(pre) == 0 {
				flag = false
			}
			for k, v := range pre {
				if v >= '0' && v <= '9' {
					flag = false
					break
				}
				if v == '!' || v == ',' || v == '.' {
					if k != len(pre)-1 {
						flag = false
						break
					}
				}
				if v == '-' {
					if cnt == 1 {
						flag = false
						break
					}
					if k == 0 || k == len(pre)-1 {
						flag = false
						break
					}
					if pre[k-1] < 'a' || pre[k-1] > 'z' || pre[k+1] < 'a' || pre[k+1] > 'z' {
						flag = false
						break
					}
					cnt = 1
				}
			}
			if flag {
				fmt.Println(string(pre))
				ans++
			}
			pre = []byte{}
		} else {
			pre = append(pre, s[i])
		}
	}
	return ans
}

func nextBeautifulNumber(n int) int {
	for i := n + 1; i <= 1224444; i++ {
		if checki(i) {
			return i
		}
	}
	return 0
}

func checki(x int) bool {
	m := map[int]int{}
	for x > 0 {
		m[x%10]++
		x /= 10
	}
	for k, v := range m {
		if k != v {
			return false
		}
	}
	return true
}

func countHighestScoreNodes(parents []int) int {
	n := len(parents)
	sub := make([][]int, n)
	max := 0
	ans := 0
	for k, v := range parents {
		if k == 0 {
			continue
		}
		sub[v] = append(sub[v], k)
	}

	cnt := make([][]int, n)

	var dfs func(i int)
	dfs = func(i int) {
		if len(sub[i]) > 0 {
			for _, v := range sub[i] {
				dfs(v)
				cnt[i] = append(cnt[i], cnt[v][0]+cnt[v][1]+1)
			}
			if len(sub[i]) == 1 {
				cnt[i] = append(cnt[i], 0)
			}
		} else {
			cnt[i] = []int{0, 0}
		}
	}
	dfs(0)

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		x := 0
		x = lf(cnt[i][0]) * lf(cnt[i][1]) * lf((n - 1 - cnt[i][0] - cnt[i][1]))
		dp[i] = x
		if x > max {
			max = x
		}
	}

	for _, v := range dp {
		if v == max {
			ans++
		}
	}
	return ans
}

func lf(a int) int {
	if a == 0 {
		return 1
	}
	return a
}

// 拓扑排序
func minimumTime(n int, relations [][]int, time []int) int {
	st := map[int]bool{}
	pre := make([]int, n)
	next := make([][]int, n)

	max := 0
	dp := make([]int, n)

	for _, arr := range relations {
		pre[arr[1]-1]++
	}
	for _, arr := range relations {
		next[arr[0]-1] = append(next[arr[0]-1], arr[1]-1)
	}

	for i := 0; i < n; i++ {
		if pre[i] == 0 {
			st[i] = true
		}
	}

	for len(st) > 0 {
		nd := map[int]bool{}
		for v, _ := range st {
			dp[v] += time[v]
			for _, t := range next[v] {
				dp[t] = max_t(dp[t], dp[v])
				pre[t]--
				if pre[t] == 0 {
					nd[t] = true
				}
			}
		}
		st = nd
	}

	for _, v := range dp {
		if v > max {
			max = v
		}
	}

	return max
}

func max_t(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumEvenGrandparent(root *TreeNode) int {
	ans := 0
	var dfs func(even bool, p *TreeNode)
	dfs = func(even bool, p *TreeNode) {
		if p.Left != nil {
			if even {
				ans += p.Left.Val
			}
			dfs(p.Val%2 == 0, p.Left)
		}
		if p.Right != nil {
			if even {
				ans += p.Right.Val
			}
			dfs(p.Val%2 == 0, p.Right)
		}
	}

	dfs(false, root)
	return ans
}

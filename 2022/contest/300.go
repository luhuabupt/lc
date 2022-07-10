package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(peopleAwareOfSecret(6, 2, 4))
	fmt.Println(peopleAwareOfSecret(4, 1, 3))
}

func countPaths(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	M := int(1e9) + 7

	ar := [][]int{}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = 1
			ar = append(ar, []int{grid[i][j], i, j})
		}
	}

	sort.Slice(ar, func(i, j int) bool {
		return ar[i][0] >= ar[j][0]
	})

	dfs := func(i, j int) {
		for _, d := range [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
			x, y := i+d[0], j+d[1]
			if x < 0 || x >= m || y < 0 || y >= n {
				continue
			}
			if grid[x][y] <= grid[i][j] {
				continue
			}
			dp[i][j] += dp[x][y]
			dp[i][j] %= M
		}
	}

	ans := 0
	for _, v := range ar {
		dfs(v[1], v[2])
		ans += dp[v[1]][v[2]]
		ans %= M
	}
	return ans
}

func peopleAwareOfSecret(n int, delay int, forget int) int {
	M := int(1e9) + 7
	a := make([]int, n*2+1) // share add
	s := make([]int, n*2+1) // sum

	s[0]++
	s[forget]--
	a[delay]++
	a[forget]--

	da := 0 // can share user
	for i := 1; i < n; i++ {
		da += a[i]

		// new
		a[i+delay] += da
		a[i+delay] %= M

		a[i+forget] += M - da
		a[i+forget] %= M

		s[i] += s[i-1] + da
		s[i] %= M

		s[i+forget] += M - da
		s[i+forget] %= M
	}

	return s[n-1] % M
}

func peopleAwareOfSecret_(n int, delay int, forget int) int {
	M := int(1e9) + 7
	a := make([][]int, n)
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n+1)
		f[i] = make([]int, n+1)
	}

	a[0][delay] = 1
	f[0][forget] = 1
	s := 1
	for i := 1; i < n; i++ {
		for j := 0; j < delay; j++ {
			a[i][j] = a[i-1][j+1]
		}

		// share
		a[i][delay] += a[i][0]
		a[i][delay] %= M
		f[i][forget] += a[i][0]
		f[i][forget] %= M

		for j := 0; j < forget; j++ {
			f[i][j] = f[i-1][j+1]
		}

		s += M + a[i][0] - f[i][0]
		s %= M

		fmt.Println(i, a[i], f[i], s)
	}

	return s
}

func decodeMessage(s string, m string) string {
	r := map[int32]byte{}
	idx := 0
	for _, v := range s {
		if v == ' ' {
			continue
		}
		if _, ok := r[v]; !ok {
			r[v] = byte('a' + idx)
			//fmt.Println(string(v), string(r[v]))
			idx++
		}
	}
	ans := []byte{}
	for _, x := range m {
		if x == ' ' {
			ans = append(ans, ' ')
			continue
		}
		ans = append(ans, r[x])
	}
	return string(ans)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			ans[i][j] = -1
		}
	}

	i, j := 0, 0
	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	id := 0
	for p := head; p != nil; {
		ans[i][j] = p.Val
		d := dir[id]
		if i+d[0] < 0 || i+d[0] >= m || j+d[1] < 0 || j+d[1] >= n {
			id++
			id %= 4
			continue
		}
		if ans[i+d[0]][j+d[1]] >= 0 {
			id++
			id %= 4
			continue
		}
		i += d[0]
		j += d[1]
		p = p.Next
	}
	return ans
}

package main

import (
	"fmt"
	"sort"
)

func longestSubsequence(arr []int, difference int) int {
	dp := map[int]int{}
	ans := 1
	for _, v := range arr {
		if dp[v-difference] == 0 {
			dp[v] = 1
		} else if dp[v-difference] >= dp[v] {
			dp[v] = dp[v-difference] + 1
			if dp[v] > ans {
				ans = dp[v]
			}
		}
	}

	return ans
}

func maximalPathQuality(values []int, edges [][]int, maxTime int) int {
	n := len(values)
	ans := 0

	sub := make([][][2]int, n)
	for _, arr := range edges {
		sub[arr[1]] = append(sub[arr[1]], [2]int{arr[0], arr[2]})
		sub[arr[0]] = append(sub[arr[0]], [2]int{arr[1], arr[2]})
	}

	vis := map[int]int{}
	var dfs func(i, t, cur int)
	dfs = func(i, t, cur int) {
		if vis[i] == 0 {
			cur += values[i]
		}
		vis[i]++

		if cur > ans && i == 0 {
			ans = cur
		}

		for _, son := range sub[i] {
			if son[1] <= t {
				dfs(son[0], t-son[1], cur)
			}
		}

		vis[i]--
	}
	dfs(0, maxTime, 0)

	return ans
}

func minimizedMaximum(n int, quantities []int) int {
	return sort.Search(100000, func(i int) bool {
		if i == 0 {
			return false
		}

		t := 0
		for _, v := range quantities {
			t += v / i
			if v%i > 0 {
				t++
			}
		}

		return t <= n
	})
}

func possiblyEquals(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)

	dp := make([][][]int, n1) // i,j s2比s1多几个字符
	for i := 0; i < n1; i++ {
		dp[i] = make([][]int, n2)
		for j := 0; j < n2; j++ {
			dp[i][j] = []int{}
		}
	}

	for i, j := 0, 0; i < n1 && j < n2; i++ {
		for k, _ := range dp[i][j-1] {
			if k == 0 {
				if num(s2[j]) {
					dp[i][j] = append(dp[i][j], toInt(s2[j:j+1]))

					if j+1 < n2 && num(s2[j+1]) {
						dp[i][j+1] = append(dp[i][j], toInt(s2[j:j+2]))
					}

					if j+2 < n2 && num(s2[j+2]) {
						dp[i][j+2] = append(dp[i][j], toInt(s2[j:j+3]))
					}
				} else {
					break
				}
			} else if k > 0 {
				if num(s2[j]) {
					dp[i][j] = append(dp[i][j], toInt(s2[j:j+1])+k)

					if j+1 < n2 && num(s2[j+1]) {
						dp[i][j+1] = append(dp[i][j], toInt(s2[j:j+2])+k)
					}

					if j+2 < n2 && num(s2[j+2]) {
						dp[i][j+2] = append(dp[i][j], toInt(s2[j:j+3])+k)
					}
				} else {
					break
				}
			}
		}
	}

	return false
}

func num(a byte) bool {
	return a >= '1' && a < '9'
}

func toInt(v string) int {
	ans := 0
	p := 1
	for i := len(v) - 1; i >= 0; i++ {
		ans += int(v[i]-'0') * p
		p *= 10
	}
	return ans
}

type p struct {
	b    []byte
	h    map[byte]int
	used int
}

func findMinStep(board string, hand string) int {
	hc := map[byte]int{}
	for _, v := range hand {
		hc[byte(v)]++
	}
	m := len(hand)

	vis := map[string]bool{}
	queue := []p{{[]byte(board), hc, 1}}
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		for k, v := range head.h {
			nh := map[byte]int{}
			for x, y := range head.h {
				nh[x] = y
			}
			if v == 1 { // 计数剪枝
				delete(nh, k)
			} else {
				nh[k]--
			}
			for i := 0; i <= len(head.b); i++ {
				for i < len(head.b) && k == head.b[i] { // 剪枝: 相同的球无序
					i++
				}
				todo := append([]byte{}, head.b[:i]...)
				todo = append(todo, k)
				todo = append(todo, head.b[i:]...)
				left := rm(todo)

				if len(left) == 0 {
					return head.used
				} else if head.used < m {
					key := fmt.Sprintf("%v_%v", left, nh)
					if vis[key] { // 记忆剪枝
						continue
					} else {
						vis[key] = true
					}

					queue = append(queue, p{left, nh, head.used + 1})
				}
			}
		}
	}

	return -1
}

func rm(s []byte) []byte {
	for i := 2; i < len(s); i++ {
		if i >= 2 && s[i] == s[i-1] && s[i-1] == s[i-2] {
			ori := i - 2
			for ; i+1 < len(s) && s[i] == s[i+1]; i++ {
			}
			s = append(append([]byte{}, s[:ori]...), s[i+1:]...)
			i = ori - 1
		}
	}
	return s
}

func getMoneyAmount(n int) int {
	dp := make([][2]int, n+2) // 几个数:0-所需钱, 1-选择次数
	dp[1] = [2]int{0, 0}
	dp[2] = [2]int{1, 1}
	//dp[3] = [2]int{2, 1}
	//dp[4] = [2]int{4, 2}
	//dp[5] = [2]int{6, 2}
	//dp[6] = [2]int{9, 2}
	//dp[7] = [2]int{10, 2}
	//dp[8] = [2]int{6, 2}
	//dp[9] = [2]int{6, 2}
	//dp[10] = [2]int{6, 2}
	for i := 3; i <= n; i++ {
		min, minTimes := 100000000, 0
		for j := 2; j <= i; j++ {
			cost, times := j, 1+max(dp[j-1][1], dp[i-j][1])
			cost += max(dp[j-1][0], dp[i-j][0]+dp[i-j][1]*j)
			if cost < min {
				min, minTimes = cost, times
			}
		}
		dp[i] = [2]int{min, minTimes}
	}
	fmt.Println(dp[n])
	return dp[n][0]
}

func isRectangleCover(rectangles [][]int) bool {
	cnt := map[int]map[int]int{}
	s := 0
	l, r, t, b := 1<<30, -1<<31, -1<<31, 1<<30
	for _, arr := range rectangles {
		l, r, t, b = min(l, arr[0]), max(r, arr[2]), max(t, arr[3]), min(b, arr[1])
		s += (arr[3] - arr[1]) * (arr[2] - arr[0])

		if cnt[arr[0]] == nil {
			cnt[arr[0]] = map[int]int{}
		}
		if cnt[arr[2]] == nil {
			cnt[arr[2]] = map[int]int{}
		}
		cnt[arr[0]][arr[1]]++
		cnt[arr[0]][arr[3]]++
		cnt[arr[2]][arr[1]]++
		cnt[arr[2]][arr[3]]++
	}

	if cnt[l][t] != 1 || cnt[r][t] != 1 || cnt[l][b] != 1 || cnt[r][b] != 1 {
		return false
	}
	delete(cnt[l], t)
	delete(cnt[l], b)
	delete(cnt[r], t)
	delete(cnt[r], b)

	if (r-l)*(t-b) != s {
		return false
	}

	for _, arr := range cnt {
		for _, v := range arr {
			if v != 2 && v != 4 {
				return false
			}
		}
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	sort.Ints(tasks)
	sort.Ints(workers)
	m, n := len(tasks), len(workers)

	return sort.Search(min(m, n)+1, func(k int) bool {
		if k > min(m, n) { // 为了兼容二分查找(Go用的不好想的野路子 Orz)
			return false
		}
		todo, p := []int{}, pills
		idx := 0 // 任务

		for i := n - k; i < n; i++ {
			// 吃药可做任务加入列表
			for ; idx < k && workers[i]+strength >= tasks[idx]; idx++ {
				todo = append(todo, tasks[idx])
			}
			if len(todo) == 0 { // 没有任何任务
				return true
			}
			if workers[i] >= todo[0] { // 不吃药则做最小task
				todo = todo[1:]
			} else { // 吃药则做可做列表最大task
				if p == 0 {
					return true
				}
				todo = todo[:len(todo)-1]
				p--
			}
		}
		return false
	}) - 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(kMirror(7, 17))
}

func kMirror(k int, n int) int64 {
	tk, t0 := []int{}, []int{}
	ak, a0 := []int{}, []int{}
	for i := 1; i < k; i++ {
		tk = append(tk, i)
		ak = append(ak, i)
	}
	for i := 1; i < 10; i++ {
		t0 = append(t0, i)
		a0 = append(a0, i)
	}

	var mk func(l, k int, t, a []int) ([]int, []int)
	mk = func(l, k int, t, a []int) ([]int, []int) {
		p := pow(k, l-1)
		nt := []int{}
		for _, v := range t {
			x := mir(v, k, p)
			if l%2 == 0 {
				a = append(a, x)
			} else {
				add := pow(k, l/2)
				for j := 0; j < k; j++ {
					nt = append(nt, v+j*add)
					a = append(a, x+j*add)
				}
			}
		}
		if l%2 == 1 {
			t = nt
		}
		return a, t
	}

	lk, l0 := 2, 2
	di, dj := 0, 0
	ans := []int{}
	for len(ans) < n {
		if ak[di] == a0[dj] {
			ans = append(ans, ak[di])
			di++
			dj++
		} else if ak[di] > a0[dj] {
			dj++
		} else {
			di++
		}
		if di == len(ak) {
			ak, tk = mk(lk, k, tk, ak)
			lk++
		}
		if dj == len(a0) {
			a0, t0 = mk(l0, 10, t0, a0)
			l0++
		}
	}

	a := int64(0)
	for i := 0; i < n; i++ {
		a += int64(ans[i])
	}
	return a
}

func pow(a, n int) int {
	ans := 1
	for n > 0 {
		if n&1 > 0 {
			ans *= a
		}
		a *= a
		n /= 2
	}
	return ans
}

func mir(v, k, p int) int {
	x := v
	for v > 0 {
		x += (v % k) * p
		v /= k
		p /= k
	}
	return x
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	a := [][]int{}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			a = append(a, []int{arr[i], arr[j]})
		}
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i][0]*a[j][1] < a[j][0]*a[i][1]
	})

	return a[k-1]
}

package main

import "fmt"

func main() {
	fmt.Println(" ")
}

//[2,3,1,3,2,3,3,3,1,1,3,2,2,2]
//18
//9

func maxOutput(n int, g [][]int, p []int) int64 {
	nx := make([][]int, n)
	for _, v := range g {
		nx[v[0]] = append(nx[v[0]], v[1])
	}

	mx := 0
	mi := -1
	vis := make([]bool, n)
	var d func(i, s int)
	d = func(i, s int) {
		vis[i] = true
		s += p[i]
		if s > mx {
			mx = s
			mi = i
		}
		for _, x := range nx[i] {
			if vis[x] {
				continue
			}
			d(x, s)
		}
	}
	d(0, 0)

	vis = make([]bool, n)
	var do func(i, s int)
	do = func(i, s int) {
		vis[i] = true

		for _, x := range nx[i] {
			if vis[x] {
				continue
			}
			d(x, s)
		}
	}

	return int64(ans)
}

//func maxOutput_(n int, g [][]int, p []int) int64 {
//	min := func(a, b int) int {
//		if a < b {
//			return a
//		}
//		return b
//	}
//	max := func(a, b int) int {
//		if a > b {
//			return a
//		}
//		return b
//	}
//	maxTwo := func(a,b, v int) (int, int) {
//		if v >= a {
//			a, b = v, a
//		} else if v > b {
//			b = v
//		}
//		return a, b
//	}
//
//	nx := make([][]int, n)
//	for _, v := range g {
//		nx[v[0]] = append(nx[v[0]], v[1])
//	}
//
//	vis := make([]bool, n)
//	var dfs func(i, s int) int
//	dfs = func(i, s int) int {
//		// 过根的最长边
//		vis[i] =true
//		s += p[i]
//		a, b := 0, 0
//		for _, x := range nx[i] {
//			if vis[x] {
//				continue
//			}
//			xa := dfs(x, s)
//			a,b = maxTwo(a, b, xa)
//		}
//	}
//
//	dfs(0, 0, 1<<60)
//}

func countGood(a []int, k int) int64 {
	c := []int{}
	n := len(a)
	t := 0
	l := 0

	ans := 0
	for i, v := range a {
		t += c[v]
		c[v]++
		if t < k {
			continue
		}
		ans += n - i
		for l <= i && t >= k {
			t -= c[l] - 1
			c[l]--
			l++
		}
	}

	return int64(ans)
}

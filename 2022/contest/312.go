package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(" ")
}

func numberOfGoodPaths(a []int, e [][]int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(a)
	ev := make([][]int, len(e))
	for i := 0; i < len(e); i++ {
		ev[i] = []int{max(a[e[i][0]], a[e[i][1]]), e[i][0], e[i][1]}
	}
	sort.Slice(ev, func(i, j int) bool {
		return ev[i][0] < ev[j][0]
	})

	fa := make([]int, n)
	son := make([]map[int]int, n) // 联通块下面的值
	for i, v := range a {
		fa[i] = i
		son[i] = map[int]int{v: 1}
	}

	var find func(i int) int
	find = func(i int) int {
		x := fa[i]
		if x == i {
			return i
		}
		return find(x)
	}

	ans := n
	for _, v := range ev {
		i, j := v[1], v[2]
		fi, fj := find(i), find(j)
		if fi == fj {
			continue
		}

		ans += son[fi][v[0]] * son[fj][v[0]]

		// union
		if len(son[fi]) < len(son[fj]) {
			fi, fj = fj, fi
		}
		fa[j], fa[i] = fi, fi
		for k, v := range son[fj] {
			son[fi][k] += v
		}
		son[fj] = map[int]int{}
	}

	return ans
}

func goodIndices(a []int, k int) []int {
	n := len(a)
	l := make([]int, n)
	r := make([]int, n)
	for i := 1; i < n; i++ {
		if a[i] <= a[i-1] {
			l[i] = l[i-1]
		} else {
			l[i] = i
		}
	}
	for i := n - 2; i >= 0; i-- {
		if a[i] <= a[i+1] {
			r[i] = r[i+1]
		} else {
			r[i] = i
		}
	}

	fmt.Println(a)
	fmt.Println(l)
	fmt.Println(r)

	ans := []int{}
	for i := k; i+k < n; i++ {
		if l[i-1] <= i-k && r[i+1] >= i+k {
			ans = append(ans, i)
		}
	}
	return ans
}

func longestSubarray(a []int) int {
	mx := 0
	for _, v := range a {
		if v > mx {
			mx = v
		}
	}
	t := 0
	ans := 0
	for i := 0; i < len(a); i++ {
		if a[i] == mx {
			t++
		} else {
			t = 0
		}
		if t > ans {
			ans = t
		}
	}
	return ans
}

func sortPeople(a []string, b []int) []string {
	h := make([]string, 1e5+2)
	for i, v := range a {
		h[b[i]] = v
	}
	ans := []string{}
	for i := 100001; i >= 0; i-- {
		if h[i] != "" {
			ans = append(ans, h[i])
		}
	}
	return ans
	//type p struct {
	//	a string
	//	b int
	//}
	//d := []*p{}
	//for i, v := range a {
	//	d = append(d, &p{v, b[i]})
	//}
	//sort.Slice(d, func(i, j int) bool {
	//	return d[i].b > d[j].b
	//})
	//ans := []string{}
	//for _, v := range d {
	//	ans = append(ans, v.a)
	//}
	//return a
}

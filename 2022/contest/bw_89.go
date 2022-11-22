package main

import "fmt"

func main() {
	fmt.Println(" ")
}

func componentValue(a []int, e [][]int) int {
	n := len(a)
	nx := make([][]int, n)
	sum := 0
	for _, v := range a {
		sum += v
	}

	for _, v := range e {
		x, y := v[0], v[1]
		nx[x] = append(nx[x], y)
		nx[y] = append(nx[y], x)
	}

	vis := make([]bool, n)
	var dfs func(i, v int) int
	dfs = func(i, v int) int {
		t := a[i]
		vis[i] = true
		for _, x := range nx[i] {
			r := dfs(x, v)
			if r != v {
				t += r
			}
		}
		fmt.Println(v, i, a[i], t)

		return t
	}

	for i := n; i >= 2; i-- {
		if sum%i != 0 {
			continue
		}
		vis = make([]bool, n)
		fmt.Println("i   ", i)
		if dfs(0, sum/i) == sum/i {
			return i - 1
		}
	}

	return 0
}

func productQueries(n int, q [][]int) []int {
	a := []int{}
	for i := 0; i < 32; i++ {
		if n&(1<<i) > 0 {
			a = append(a, 1<<i)
		}
	}

	M := int(1e9) + 7
	d := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		d[i] = make([]int, len(a))
		t := 1
		for j := i; j < len(a); j++ {
			t *= a[j]
			t %= M
			d[i][j] = t
		}
	}

	ans := make([]int, len(q))
	for i, v := range q {
		ans[i] = d[v[0]][v[1]]
	}

	return ans
}

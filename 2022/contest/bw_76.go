package main

import "fmt"

func main() {
	//fmt.Println(maximumScore([]int{5, 2, 9, 8, 4}, [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 2}, {1, 3}, {2, 4}}))
	fmt.Println(maximumScore([]int{14, 12, 10, 8, 1, 2, 3, 1}, [][]int{{1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 1}, {6, 1}, {7, 1}, {2, 1}})) // 44
}

func findClosestNumber(nums []int) int {
	a := 10000000
	for _, v := range nums {
		if abs(v) < abs(a) {
			a = v
		}
		if abs(v) == abs(a) && v > 0 && a < 0 {
			a = v
		}
	}
	return a
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func waysToBuyPensPencils(t int, a int, b int) int64 {
	ans := 0
	for i := 0; i <= t/a; i++ {
		ans += (t-a*i)/b + 1
	}
	return int64(ans)
}

type ATM struct {
	a []int
}

func Constructor() ATM {
	return ATM{make([]int, 5)}
}

func (t *ATM) Deposit(banknotesCount []int) {
	for i, v := range banknotesCount {
		t.a[i] += v
	}
}

func (t *ATM) Withdraw(c int) []int {
	x := []int{20, 50, 100, 200, 500}

	ta := append([]int{}, t.a...)
	ans := make([]int, 5)
	for i := 4; i >= 0; i-- {
		do := min(c/x[i], ta[i])
		c -= x[i] * do
		ta[i] -= do
		ans[i] += do
	}
	if c > 0 {
		return []int{-1}
	}
	t.a = ta
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximumScore(sc []int, edges [][]int) int {
	n := len(sc)
	next := make([][]int, n)
	ma := make([][][2]int, n) // 最大2个节点值
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for _, v := range edges {
		next[v[0]] = append(next[v[0]], v[1])
		next[v[1]] = append(next[v[1]], v[0])
	}

	for i, ar := range next {
		a, b, c, ai, bi, ci := 0, 0, 0, -1, -1, -1
		for _, x := range ar {
			if sc[x] >= a {
				a, b, c = sc[x], a, b
				ai, bi, ci = x, ai, bi
			} else if sc[x] >= b {
				b, c = sc[x], b
				bi, ci = x, bi
			} else if sc[x] > c {
				c = sc[x]
				ci = x
			}
		}
		if ai > -1 {
			ma[i] = append(ma[i], [2]int{a, ai}) // sc, idx
		}
		if bi > -1 {
			ma[i] = append(ma[i], [2]int{b, bi})
		}
		if ci > -1 {
			ma[i] = append(ma[i], [2]int{c, ci})
		}
	}

	ans := -1
	for _, v := range edges {
		x, y := v[0], v[1]
		for i := 0; i < 9; i++ {
			if i%3 >= len(ma[x]) || i/3 >= len(ma[y]) || ma[x][i%3][1] == y || ma[y][i/3][1] == x || ma[x][i%3][1] == ma[y][i/3][1] {
				continue
			}
			ans = max(ans, sc[x]+sc[y]+ma[x][i%3][0]+ma[y][i/3][0])
		}
	}

	return ans
}

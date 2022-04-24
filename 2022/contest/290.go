package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("")

}

func countLatticePoints(c [][]int) (ans int) {
	chk := func(i, j, a, b, r int) bool {
		return (a-i)*(a-i)+(b-j)*(b-j) <= r*r
	}

	for i := 0; i <= 200; i++ {
		for j := 0; j <= 200; j++ {
			for _, a := range c {
				if chk(i, j, a[0], a[1], a[2]) {
					ans++
					break
				}
			}
		}
	}

	return
}

func fullBloomFlowers(f [][]int, p []int) []int {
	d := map[int]int{}
	for _, v := range f {
		d[v[0]]++
		d[v[1]+1]--
	}

	ar := [][2]int{}
	for k, v := range d {
		ar = append(ar, [2]int{k, v})
	}
	sort.Slice(ar, func(i, j int) bool {
		return ar[i][0] < ar[j][0]
	})
	pre := 0
	for k, v := range ar {
		pre += v[1]
		ar[k][1] = pre
	}

	ans := []int{}
	for _, v := range p {
		x := sort.Search(len(ar), func(j int) bool {
			return ar[j][0] > v
		})
		if x == 0 {
			ans = append(ans, 0)
		} else {
			ans = append(ans, ar[x-1][1])
		}
	}

	return ans
}

func countRectangles(r [][]int, p [][]int) []int {
	dp := [101][]int{}
	for _, v := range r {
		dp[v[1]] = append(dp[v[1]], v[0])
	}
	for i := 0; i < 101; i++ {
		sort.Ints(dp[i])
	}

	ans := make([]int, len(p))
	for i, v := range p {
		x, y := v[0], v[1]
		for j := y; j < 101; j++ {
			ans[i] += len(dp[j]) - sort.SearchInts(dp[j], x)
		}
	}
	return ans
}

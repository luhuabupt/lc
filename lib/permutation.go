package main

import "fmt"

func main() {
	fmt.Println(Factorial(6))
}

// C-n-r 的排列
func Permutation(n, r int) {

}

func Factorial(n int) [][]int { // 阶乘
	a := []int{}
	for i := 0; i < n; i++ {
		a = append(a, i)
	}

	ans := [][]int{}

	var dfs func(undo, cur []int)
	dfs = func(undo, cur []int) {
		if len(cur) == n {
			ans = append(ans, cur)
			return
		}
		for i, x := range undo {
			nc := append([]int{}, cur...)
			nc = append(nc, x)

			nd := append([]int{}, undo[:i]...)
			nd = append(nd, undo[i+1:]...)

			dfs(nd, nc)
		}
	}

	dfs(a, []int{})

	return ans
}

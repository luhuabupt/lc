package main

import "fmt"

func main() {
	for i := 1; i < 5; i++ {
		for j := 0; j < 1<<(i-1); j++ {
			fmt.Printf("%d", kthGrammar(i, j+1))
		}
		fmt.Println(" ")
	}
	//fmt.Println(kthGrammar(3, 3))
}

//leetcode submit region begin(Prohibit modification and deletion)
func kthGrammar(n int, k int) int {
	k--
	d := [][]int{{0, 1}, {1, 0}} // 0-01 1-10
	x := 0
	for i := 1; i <= n; i++ {
		x = d[x][(k>>(n-i))%2]
	}

	return x
}

//leetcode submit region end(Prohibit modification and deletion)

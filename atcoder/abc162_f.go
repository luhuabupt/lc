package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func main() {
	f162(os.Stdin, os.Stdout)
}

func f162(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	max := func(a ...int) int {
		sort.Ints(a)
		return a[len(a)-1]
	}

	var n int
	var x int

	Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &x)
		a[i] = x
	}

	if n == 2 {
		Fprintln(out, max(a[0], a[1]))
		return
	} else if n == 3 {
		Fprintln(out, max(a[0], a[1], a[2]))
		return
	}

	//  floor(i/2)个  1-0 2-1 3-1 4-2 5-2 6-3 7-3
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0
	dp[2] = max(a[0], a[1])
	t := a[1]
	for i := 3; i <= n; i++ {
		if i%2 == 1 {
			t += a[i-1]
			dp[i] = dp[i-1]
		} else {
			dp[i] = t
		}

		// 选
		dp[i] = max(dp[i], dp[i-2]+a[i-1])
	}
	Fprintln(out, dp[n])

	//dp := make([][]int, n)
	//dp[0] = []int{0, a[0]}
	//dp[1] = []int{a[1], -1<<62}
	//dp[2] = []int{a[2], a[0] + a[2]}
	//for i := 3; i < n; i++ { // i/2-1 ~ i/2+1 0 0-1 |  1 1 | 2 1-2 | 3 2 |  4 2-3  | 5 3  | 6 3-4 |  7 4 |  8 4-5
	//	dp[i] = []int{-1 << 62, -1 << 62}
	//	if i%2 == 1 {
	//		dp[i][0] = max(dp[i-2][0], dp[i-3][1]) + a[i]
	//	} else {
	//		dp[i][0] = max(dp[i-2][0], dp[i-3][0], dp[i-4][1]) + a[i]
	//		dp[i][1] = dp[i-2][1] + a[i]
	//	}
	//}
	////Fprintln(out, dp)
	//if n%2 == 0 {
	//	Fprintln(out, max(dp[n-1][0], dp[n-2][1]))
	//} else {
	//	Fprintln(out, max(dp[n-1][0], dp[n-2][0], dp[n-3][1]))
	//}
}

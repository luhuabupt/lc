package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func main() {
	f998(os.Stdin, os.Stdout)
}

func f998(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var dst int
	var n int
	var m int
	var x int
	var y int

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	Fscan(in, &dst)
	Fscan(in, &n)
	Fscan(in, &m)

	r := make([]bool, dst)
	w := map[int]int{}
	wk := []int{}
	for i := 0; i < n; i++ {
		Fscan(in, &x, &y)
		for j := x; j < y; j++ {
			r[j] = true
		}
	}
	for i := 0; i < m; i++ {
		Fscan(in, &x, &y)
		if v, ok := w[x]; !ok || y < v {
			w[x] = y
		}
	}
	for k, _ := range w {
		wk = append(wk, k)
	}
	sort.Ints(wk)

	//Fprintln(out, r, w, wk)
	dp := make([]int, dst)
	if r[0] {
		if _, ok := w[0]; !ok {
			Fprintln(out, -1)
			return
		}
		dp[0] = w[0]
	}
	for i := 1; i < dst; i++ {
		if !r[i] && i >= 1 {
			dp[i] = dp[i-1]
			continue
		}

		dp[i] = 1 << 31
		for j := i; j >= 0; j-- {
			if v, ok := w[j]; ok {
				cur := (i - j + 1) * v
				if j > 0 {
					cur += dp[j-1]
				}
				dp[i] = min(dp[i], cur)
			}
		}
		if dp[i] == 1<<31 {
			Fprintln(out, -1)
			return
		}
	}

	//Fprintln(out, dp)
	Fprintln(out, dp[dst-1])
}

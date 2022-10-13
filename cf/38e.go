package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func main() {
	e38(os.Stdin, os.Stdout)
}

func e38(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var x int
	var y int64

	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}

	Fscan(in, &n)
	al := map[int]int64{}
	a := []int{}

	for i := 0; i < n; i++ {
		Fscan(in, &x)
		Fscan(in, &y)
		if v, ok := al[x]; !ok || y < v {
			al[x] = y
		}
	}
	for k, _ := range al {
		a = append(a, k)
	}
	sort.Ints(a)
	n = len(a)

	dp := make([][2]int64, n) // 0 1-do
	dp[0] = [2]int64{1 << 60, al[a[0]]}
	for i := 1; i < n; i++ {
		mi := int64(1 << 60)
		t := int64(0)
		for j := i - 1; j >= 0; j-- {
			t += int64(a[j+1]-a[j]) * int64(i-1-j)
			mi = min(mi, t+dp[j][1])
		}
		dp[i][1] = mi + al[a[i]]

		mi = int64(1 << 60)
		t = int64(0)
		for j := i - 1; j >= 0; j-- {
			t += int64(a[j+1]-a[j]) * int64(i-j)
			mi = min(mi, t+dp[j][1])
		}
		dp[i][0] = mi
	}

	Fprintln(out, dp)

	Fprintln(out, min(dp[n-1][0], dp[n-1][1]))
}

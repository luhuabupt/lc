package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	e1740(os.Stdin, os.Stdout)
}

func e1740(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)

		a := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			a[i] = x
		}

		d := make([]map[int]bool, n)
		for i := 0; i < n; i++ {
			d[i] = map[int]bool{}
			if a[i]+i < n {
				d[i][a[i]+i] = true
			}
		}
		for i := n - 1; i >= 0; i-- {
			if i-a[i] >= 0 {
				d[i-a[i]][i] = true
			}
		}

		vis := make([]bool, n)
		var dfs func(i int) bool
		dfs = func(i int) bool {
			if vis[i] {
				return false
			}
			for x, _ := range d[i] {
				if x == n-1 {
					return true
				}
				if dfs(x + 1) {
					return true
				}
			}

			vis[i] = true
			return false
		}

		ans := "YES"
		if !dfs(0) {
			ans = "NO"
		}

		Fprintln(out, ans)
	}
}

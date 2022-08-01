package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1707(os.Stdin, os.Stdout)
}

//1
//11
//110
//1110
//01111

func a1707(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var q int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			a[i] = x
		}

		ans := make([]byte, n)
		t := 0 // 到i之前
		for i := n - 1; i >= 0; i-- {
			if a[i] > t {
				if t < q {
					t++
					ans[i] = '1'
				} else {
					ans[i] = '0'
				}
			} else {
				ans[i] = '1'
			}
		}
		Fprintln(out, string(ans))
	}
}

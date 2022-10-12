package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	b1740(os.Stdin, os.Stdout)
}

func b1740(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)

		ans := []int{}
		if n == 3 {
			ans = []int{-1}
		} else if n%2 == 0 {
			for i := 0; i < n; i++ {
				ans = append(ans, n-i)
			}
		} else {
			ans = append(ans, n/2+1)
			for i := 0; i < n/2; i++ {
				ans = append(ans, ans[0]+i+1)
			}
			for i := 0; i < n/2; i++ {
				ans = append(ans, i+1)
			}
		}

		for _, v := range ans {
			Fprintf(out, "%d ", v)
		}
		Fprintln(out, "")
	}
}

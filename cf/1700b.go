package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	b1700(os.Stdin, os.Stdout)
}

func b1700(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int

	for Fscan(in, &T); T > 0; T-- {

		Fscan(in, &n)

		ans := []int{}
		do := make([]bool, n+1)
		for i := 1; i <= n; i *= 2 {
			do[i] = true
			ans = append(ans, i)
		}
		for i := 1; i <= n; i++ {
			if !do[i] {
				ans = append(ans, i)
				for j := i * 2; j <= n && !do[j]; j *= 2 {
					ans = append(ans, j)
					do[j] = true
				}
			}
		}

		Fprintln(out, 2)
		Fprintln(out, ans)
	}
}

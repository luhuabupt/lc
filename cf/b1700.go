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
	var s string

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fscan(in, &s)
		ans := make([]byte, n)
		for i := 0; i < n; i++ {
			ans[i] = '9' - s[i] + '0'
		}
		if s[0] == '9' {
			ans[0]++
			for p, i := 12, n-1; i >= 0; i-- {
				x := int(ans[i] - '0')
				x += p
				ans[i] = byte('0' + x%10)

				p = x / 10
				if p == 0 {
					break
				}
			}
		}

		Fprintln(out, string(ans))
	}
}

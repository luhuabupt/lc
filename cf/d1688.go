package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	d1688(os.Stdin, os.Stdout)
}

func d1688(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var k int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fscan(in, &k)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			a[i] = x
		}

		ans := int64(0)
		if k < n {
			s := int64(0)
			for i := 0; i < k; i++ {
				s += int64(a[i])
			}
			ans = s
			for i := k; i < n; i++ {
				s += int64(a[i] - a[i-k])
				if s > ans {
					ans = s
				}
			}
			ans += (int64(k) * int64(k-1)) / 2
		} else {
			for _, v := range a {
				ans += int64(v)
			}
			ans += int64(n) * (int64(k-1) + int64(k-n)) / 2
		}

		Fprintln(out, ans)
	}
}

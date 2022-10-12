package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	b1736(os.Stdin, os.Stdout)
}

func b1736(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var x int

	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)

		for i := 0; i < n; i++ {
			Fscan(in, &x)
			a[i] = x
		}

		ans := "YES"
		for i := 1; i < n-1; i++ {
			if gcd(a[i-1]/gcd(a[i-1], a[i]), a[i+1]/gcd(a[i+1], a[i])) != 1 {
				ans = "NO"
				break
			}
		}

		Fprintln(out, ans)
	}
}

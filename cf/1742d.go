package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	d1742(os.Stdin, os.Stdout)
}

func d1742(_r io.Reader, _w io.Writer) {
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
		r := make([]int, n)
		p := make([]int, 1001)
		for i, _ := range p {
			p[i] = -1
		}
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			r[i] = x
			p[x] = i
		}
		ans := -1
		for i := 1; i < 1001; i++ {
			for j := i; j < 1001; j++ {
				if p[i] > 0 && p[j] > 0 {
					if gcd(i, j) == 1 {
						if p[i]+p[j]+2 > ans {
							ans = p[i] + p[j] + 2
						}
					}
				}
			}
		}

		Fprintln(out, ans)
	}
}

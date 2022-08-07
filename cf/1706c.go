package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	c1706(os.Stdin, os.Stdout)
}

func c1706(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var x int

	cal := func(a, b, c int) int64 {
		if b < c {
			b, c = c, b
		}
		if a > b {
			return 0
		}
		if a == b {
			return 1
		}
		return int64(b - a + 1)
	}

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ar := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			ar[i] = x
		}

		ans := int64(0)
		if n%2 == 1 {
			for i := 1; i < n; i += 2 {
				ans += cal(ar[i], ar[i-1], ar[i+1])
			}
		} else {
			l, r := make([]int64, n), make([]int64, n)
			sl, sr := int64(0), int64(0)
			for i := 1; i <= n-3; i += 2 {
				sl += cal(ar[i], ar[i-1], ar[i+1])
				l[i] = sl
			}
			for i := n - 2; i >= 2; i -= 2 {
				sr += cal(ar[i], ar[i-1], ar[i+1])
				r[i] = sr
			}
			ans = l[n-3]
			if r[2] < ans {
				ans = r[2]
			}
			for i := 1; i+3 < n; i += 2 {
				if l[i]+r[i+3] < ans {
					ans = l[i] + r[i+3]
				}
			}
		}

		Fprintln(out, ans)
	}
}

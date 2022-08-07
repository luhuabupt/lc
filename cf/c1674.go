package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	c1674(os.Stdin, os.Stdout)
}

func c1674(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var x string
	var y string

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x)
		Fscan(in, &y)

		ac := 0
		ot := 0
		for _, v := range y {
			if v == 'a' {
				ac++
			} else {
				ot++
			}
		}
		ans := int64(0)
		if ac > 0 {
			if len(y) > 1 {
				ans = -1
			} else {
				ans = 1
			}
		} else if ac == 0 {
			ans = quickPow(2, int64(len(x)))
		}
		Fprintln(out, ans)
	}
}

func quickPow(a, n int64) int64 {
	ans := int64(1)
	for n > 0 {
		if n&1 > 0 {
			ans = ans * a
		}
		a *= a
		n /= 2
	}
	return ans
}

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	e168(os.Stdin, os.Stdout)
}

func e168(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var x int
	var y int

	gcd := func(a, b int) int {
		if a < 0 {
			a = -a
		}
		if b < 0 {
			b = -b
		}
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	Fscan(in, &n)
	m := map[int]map[int]int{}
	for i := 0; i < n; i++ {
		Fscan(in, &x, &y)
		z := gcd(x, y)
		m[(x/z)]++
	}

	M := int(1e9 + 7)
	qp := func(a, n int) int {
		ans := 1
		for n > 0 {
			if n&1 > 0 {
				ans *= a
				ans %= M
			}
			a *= a
			a %= M
			n /= 2
		}
		return ans
	}

	Fprintln(out, m)
	// -4 4 -3 3 -2 2 5
	ans := 1
	for k, v := range m {
		if k > 0 {
			if m[-k] > 0 {
				ans *= qp(2, v) + qp(2, m[-k]) - 1
			} else {
				ans *= qp(2, v)
			}
		} else if k < 0 {
			if m[-k] > 0 {
				continue
			} else {
				ans *= qp(2, v)
			}
		}
		ans %= M
	}
	ans *= m[0] + 1
	ans = (ans - 1 + M) % M

	Fprintln(out, ans)
}

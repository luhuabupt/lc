package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	c1740(os.Stdin, os.Stdout)
}

func c1740(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)

		r := make([]int, n)
		s := 0
		d := map[int]int{}
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			r[i] = x
			s += x
			d[s] = i
		}

		//Fprintln(out, r)
		//Fprintln(out, d)

		ans := n
		t := 0

	loop:
		for i := 0; i < n; i++ {
			t += r[i]
			if s%t == 0 {
				mx := i + 1
				l := 0
				for nx := t; nx <= s; nx += t {
					if _, ok := d[nx]; !ok {
						continue loop
					}
					if d[nx]-l > mx {
						mx = d[nx] - l
					}
					l = d[nx]
				}
				//Fprintln(out, i, mx)
				if mx < ans {
					ans = mx
				}
			}
		}

		Fprintln(out, ans)
	}
}

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	c1676(os.Stdin, os.Stdout)
}

func c1676(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var sn int
	var x string

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fscan(in, &sn)

		row := make([]string, n)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			row[i] = x
		}

		ans := 1<<30
		for i, v := range row {
			for j := i + 1; j < n; j++ {
				t := 0
				vv := row[j]
				for k := 0; k < sn; k++ {
					tt := int(vv[k]-v[k])
					if tt < 0 {
						tt = -tt
					}
					t += tt
				}
				if t < ans {
					ans = t
				}
			}
		}

		Fprintln(out, ans)
	}
}

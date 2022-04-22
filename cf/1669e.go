package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	CF16669E(os.Stdin, os.Stdout)
}

func CF16669E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var x string
	var n int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)

		ans := int64(0)
		cnt := map[string]int{}
		for i := 0; i < n; i++ {
			Fscan(in, &x)

			for a := byte('a'); a <= 'k'; a++ {
				if a != x[0] {
					ans += int64(cnt[string([]byte{a, x[1]})])
				}
			}
			for a := byte('a'); a <= 'k'; a++ {
				if a != x[1] {
					ans += int64(cnt[string([]byte{x[0], a})])
				}
			}

			cnt[x]++
		}

		Fprintln(out, ans)
	}
}

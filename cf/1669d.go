package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	CF16669D(os.Stdin, os.Stdout)
}

func CF16669D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var x string
	var n int

loop:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)

		//r := []byte{}
		//for i := 0; i < n; i++ {
		//	Fscan(in, &x)
		//	r = append(r, x)
		//}
		Fscan(in, &x)
		r := []byte(x)
		//Fprintln(out, string(r))

		if n == 1 && r[0] != 'W' {
			Fprintln(out, "NO")
			continue loop
		}

		c := [2]int{}
		for i := 0; i < n; i++ {
			if r[i] == 'W' {
				if c[0] > 0 && c[1] == 0 || c[0] == 0 && c[1] > 0 {
					Fprintln(out, "NO")
					continue loop
				}
				c = [2]int{}
			}
			if r[i] == 'R' {
				c[1]++
			} else if r[i] == 'B' {
				c[0]++
			}
		}

		if c[0] > 0 && c[1] == 0 || c[0] == 0 && c[1] > 0 {
			Fprintln(out, "NO")
			continue loop
		}

		Fprintln(out, "YES")
	}
}

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	CF1671A(os.Stdin, os.Stdout)
}

func CF1671A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var x string

loop:
	for Fscan(in, &T); T > 0; T-- {

		Fscan(in, &x)
		r := []byte(x)
		//Fprintln(out, string(r))

		if len(r) == 1 {
			Fprintln(out, "NO")
			continue
		}

		t := 1
		for i := 1; i < len(r); i++ {
			if r[i] == r[i-1] {
				t++
			} else {
				if t == 1 {
					Fprintln(out, "NO")
					continue loop
				}
				t = 1
			}
		}
		if t == 1 {
			Fprintln(out, "NO")
		} else {
			Fprintln(out, "YES")
		}
	}
}

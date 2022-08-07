package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a1674(os.Stdin, os.Stdout)
}

func a1674(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var x string

	m := map[string]int{}
	idx := 1
	for i := 'a'; i <= 'z'; i++ {
		for j := 'a'; j <= 'z'; j++ {
			if i != j {
				m[string(i)+string(j)] = idx
				idx++
			}
		}
	}

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x)

		Fprintln(out, m[x])
	}
}

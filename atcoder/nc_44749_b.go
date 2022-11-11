package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	b(os.Stdin, os.Stdout)
}

func b(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int

	Fscan(in, &n)
	M := int(1e9) + 7
	a := 1
	for i := 1; i <= n; i++ {
		a *= i
		a %= M
	}
	b := 1
	t := (n + 1) / 2
	for i := 2; i <= n; i += 2 {
		b *= t
		t--
		b %= M
	}
	t = (n + 1) / 2
	for i := 1; i <= n; i += 2 {
		b *= t
		t--
		b %= M
	}

	Fprintln(out, (a+M-b)%M)

}

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	a290a(os.Stdin, os.Stdout)
}

func a290a(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var x string

	Fscan(in, &x)

	c := map[byte]int{}
	for _, v := range x {
		c[byte(v)]++
	}
	for k, v := range c {
		if v == 1 {
			Fprintln(out, string(k))
			return
		}
	}
	Fprintln(out, -1)

}

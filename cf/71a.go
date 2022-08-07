package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	CF71A(os.Stdin, os.Stdout)
}

func CF71A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		o := s
		if len(o) > 10 {
			oa := []byte{s[0]}
			oa = append(oa, strconv.Itoa(len(s)-2)...)
			oa = append(oa, s[len(s)-1])
			o = string(oa)
		}
		Fprintln(out, o)
	}
}

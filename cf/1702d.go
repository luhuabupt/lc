package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main() {
	d1702(os.Stdin, os.Stdout)
}

func d1702(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s string
	var p int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		Fscan(in, &p)

		c := [27]int{}
		sum := 0
		for _, v := range s {
			c[v-'a'+1]++
			sum += int(v - 'a' + 1)
		}
		for sum > p {
			for i := 26; i > 0; i-- {
				if c[i] == 0 {
					continue
				}
				if sum-c[i]*i > p {
					sum -= c[i] * i
					c[i] = 0
				} else {
					c[i] -= (sum - p) / i
					if (sum-p)%i != 0 {
						c[i]--
					}
					sum = -1
					break
				}
			}
		}

		sb := []byte(s)
		ans := []byte{}
		for _, v := range sb {
			if c[v-'a'+1] > 0 {
				c[v-'a'+1]--
				ans = append(ans, v)
			}
		}

		Fprintln(out, string(ans))
	}
}

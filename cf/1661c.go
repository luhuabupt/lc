package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func main() {
	b1661(os.Stdin, os.Stdout)
	// 7
	//1 1 1 1 1 1 2
}

func b1661(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n int
	var x int

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		row := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			row[i] = x
		}
		sort.Ints(row)
		ma := row[n-1]
		sum := 0
		one := 0
		for _, v := range row {
			if v == ma-1 {
				one++
			}
			//else {
			sum += ma - v
			//}
		}
		ans := (sum / 3) * 2
		if sum%3 == 2 {
			ans += 2
		} else if sum%3 == 1 {
			ans++
		}
		if ans/2 < one {
			ans = one*2 - 1
		}

		if ma%3 == 2 {
			ma++
			sum := 0
			one := 0
			for _, v := range row {
				if v == ma-1 {
					one++
				}
				//else {
				sum += ma - v
				//}
			}
			t := (sum / 3) * 2
			if sum%3 == 2 {
				t += 2
			} else if sum%3 == 1 {
				t++
			}
			if t/2 < one {
				t = one*2 - 1
			}
			if t < ans {
				ans = t
			}
		}
		//if sum%3 == 1 {
		//	one++
		//} else if sum%3 == 2 {
		//	ans += 2
		//	one--
		//}

		//if one > 0 {
		//	ans += one*2 - 1
		//}
		//
		//Println(sum, one, ans)

		Fprintln(out, ans)
	}
}

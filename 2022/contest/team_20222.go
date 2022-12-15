package main

import "fmt"

func main() {
	//fmt.Println(Leetcode([]string{"hello", "leetcode"}))
	fmt.Println(Leetcode([]string{"hold", "engineer", "cost", "level"}))
}

func Leetcode(words []string) int {
	//s := "helloleetcode"
	//m := map[string]int{}
	//for _, x := range t {
	//	m[string(x)]++
	//}
	//fmt.Println(m)

	del := func(s string, d int) string {
		ar := []byte{}
		for i, v := range s {
			if i == d {
				continue
			}
			ar = append(ar, byte(v))
		}
		return string(ar)
	}

	// c:1 d:1 e:4 h:1 l:3 o:2 t:1
	// 1   1   100 1   11  10  1
	//t := 1<<0 + 1<<1 + 1<<4 + 1<<5 + 1<<6 + 1<<7 + 1<<9 + 1<<10
	t := 1779
	tw := []byte{'c', 'd', 'e', 'e', 'e', 'h', 'l', 'l', 'o', 'o', 't'}
	fmt.Printf("%2b t \n", t)
	dp := make([]int, t+1)
	for i, _ := range dp {
		dp[i] = -1
	}

	ans := 1 << 60
	var dfs func(mask, c int)
	dfs = func(mask, c int) {
		if dp[mask] != -1 && dp[mask] < c {
			return
		}
		dp[mask] = c

		for i := 0; i < 11; i++ {
			if (1<<i)&mask > 0 {
				fmt.Printf("%2b  %d %d %s %v\n", mask, c, i, string(tw[i]), words)
				// find
				mi := 1 << 60
				dw, dd := -1, -1
				for iw, w := range words {
					if mi == 0 {
						break
					}
					for ix, x := range w {
						if byte(x) == tw[i] {
							if ix*(len(w)-1-ix) < mi {
								mi = ix * (len(w) - 1 - ix)
								dw = iw
								dd = ix
							}
						}
						if mi == 0 {
							break
						}
					}
				}

				pre := words[dw]
				words[dw] = del(pre, dd)

				if tw[i] == 'e' {
					mask -= 1 << 2
				} else if tw[i] == 'l' {
					mask -= 1 << 6
				} else if tw[i] == 'o' {
					mask -= 1 << 8
				} else {
					mask -= 1 << i
				}

				dfs(mask, c+mi)

				// word back
				words[dw] = pre
			}

		}

		if mask == 0 {
			if c < ans {
				ans = c
			}
			return
		}
	}

	dfs(t, 0)

	return ans
}

// 5
// [[1,1],[2,1],[3,3],[3,5],[4,2],[5,1],[5,3],[5,5],[5,7],[5,9]]
func sandyLandManagement(n int) [][]int {
	r := make([][][]int, n+1)
	for l := 4; l <= n; l++ {
		if l%2 == 0 {
			r[l] = append(r[l], []int{l, l})
			for c := 1; c <= (l-4)/2; c++ {
				r[l] = append(r[l], []int{l, l + c*2})
				r[l] = append(r[l], []int{l, l - c*2})
			}
		} else {
			for c := 0; c <= (l-5)/2; c++ {
				r[l] = append(r[l], []int{l, l + c*2 + 1})
				r[l] = append(r[l], []int{l, l - c*2 - 1})
			}
		}
	}
	ans := [][]int{}
	for l := 1; l <= n; l++ {
		ans = append(ans, []int{l, 1})
		if l%2 == n%2 {
			ans = append(ans, r[l]...)
		}
		if l > 1 {
			ans = append(ans, []int{l, 2*l - 1})
		}
	}
	return ans
}

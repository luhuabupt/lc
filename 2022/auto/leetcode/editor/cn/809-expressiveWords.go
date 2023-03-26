package main

import "fmt"

// expressive-words 情感丰富的文字  2022-11-25 09:19:53
func main() {
	fmt.Println(expressiveWords("dddiiiinnssssssoooo", []string{"dinnssoo", "ddinso", "ddiinnso", "ddiinnssoo", "ddiinso", "dinsoo", "ddiinsso", "dinssoo", "dinso"}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func expressiveWords(s string, a []string) int {
	ans := 0

	cal := func(sv string) [][]int {
		d, p, c := [][]int{}, ' ', 1
		for i, v := range sv {
			if i > 0 {
				if v == p {
					c++
				} else {
					d = append(d, []int{int(p - 'a'), c})
					c = 1
				}
			}

			p = v
		}
		d = append(d, []int{int(p - 'a'), c})
		//fmt.Println(sv, d)
		return d
	}

	chk := func(a, b [][]int) bool {
		if len(a) != len(b) {
			return false
		}

		for i, v := range a {
			x := b[i]
			if v[0] != x[0] {
				return false
			}
			if v[1] != x[1] {
				if x[1] > v[1] || v[1] <= 2 {
					return false
				}
			}
		}

		return true
	}

	sd := cal(s)
	for _, v := range a {
		x := cal(v)
		if chk(sd, x) {
			ans++
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

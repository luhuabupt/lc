package main

import "fmt"

func main() {
	//fmt.Println(scoreOfParentheses("(()(()))"))
	fmt.Println(scoreOfParentheses("(())(())"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func scoreOfParentheses(s string) int {
	c := make([]int, 50)
	idx := 0
	for i, v := range s {
		if v == '(' {
			idx++
		} else {
			if s[i-1] == '(' {
				c[idx]++
			} else {
				c[idx] += c[idx+1] * 2
				c[idx+1] = 0
			}
			idx--
		}
	}

	//fmt.Println(c)

	return c[1]
}

//leetcode submit region end(Prohibit modification and deletion)

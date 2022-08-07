package main

import "fmt"

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	m := map[string]bool{}
	for _, v := range supplies {
		m[v] = true
	}

	do := map[string]bool{}
	for i := 0; i < len(recipes); i++ {
		for i, v := range recipes {
			flag := true
			for _, x := range ingredients[i] {
				if !m[x] {
					flag = false
					break
				}
			}
			if flag {
				m[v] = true
				do[v] = true
			}
		}
	}

	ans := []string{}
	for v, _ := range do {
		ans = append(ans, v)
	}

	return ans
}

func canBeValid(s string, locked string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	l, c := 0, 0
	for i, v := range s {
		if locked[i] == '0' {
			c++
			continue
		}
		if v == '(' {
			l++
		} else {
			if l == 0 {
				if c > 0 {
					c--
				} else {
					return false
				}
			} else {
				l--
			}
		}
	}
	if l != 0 && l > 2*c {
		return false
	}

	r, c := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if locked[i] == '0' {
			c++
			continue
		}
		if s[i] == ')' {
			r++
		} else {
			if r == 0 {
				if c > 0 {
					c--
				} else {
					return false
				}
			} else {
				r--
			}
		}
	}
	if r != 0 && r > 2*c {
		return false
	}

	return true
}

func main() {
	fmt.Println(canBeValid("()))", "1011"))
	fmt.Println(canBeValid("(()))())", "11001111"))
	fmt.Println(canBeValid("(()))()))(()((()()(((()))()()()()()())))()()(()())()(()((()()((()((((((()(()()(()()())(((((())((()))))()(((((((()()())()))())((((((()(())())()((())()()((())((((())(((())(())()()))(((()()())())))((()))))()()()((()))())(())(((()()((())(())(())())()((()))())(())()(()())((((()(((())((()()())(((()(((((()))()))))))(()((())())(())))))(())(((())()()(()))())())))(((())))()()))()())))))())()(()()))(())(()())))())()))((((())(()))()(((())())(()(()))()))((()(())()()))))())(()(((((()",
		"110001111001011100000100011110101000100110010010011001110010111111100111000100000000101111101001111111011101001110011001001100100001100000000010100010101101100000100001000110111000111110110010111011010010100011111101110011100010010001111001010001001000111101101111111011001000100111100110101000100011011001001100110011111111111100101000100111111100000100101101000101111101000011110001001011111010011010000100100000000011101011001110000110011000100001110101100101111111110100"))
}

func abbreviateProduct(left int, right int) string {

}

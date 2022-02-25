package main

import "fmt"

func main() {
	fmt.Println(alienOrder([]string{"ac", "ab", "zc", "zb"}))
	fmt.Println(alienOrder([]string{"z", "x", "a", "zb", "zx"}))
}

func alienOrder(w []string) string {
	n := len(w)

	e := [26][26]bool{}
	in, out, all := [26]int{}, [26]bool{}, [26]int{}
	for _, v := range w {
		for _, x := range v {
			all[x-'a'] = 2
		}
	}
	for i := 1; i < n; i++ {
		a, b := w[i-1], w[i]
		if len(a) > len(b) && a[:len(b)] == b {
			return ""
		}
		for j := 0; j < len(b) && j < len(a); j++ {
			if a[j] != b[j] {
				if e[b[j]-'a'][a[j]-'a'] {
					//fmt.Println(i, j)
					return ""
				}
				if !e[a[j]-'a'][b[j]-'a'] {
					in[b[j]-'a']++
				}
				e[a[j]-'a'][b[j]-'a'] = true
				out[a[j]-'a'] = true
				break
			}
		}
	}

	ans := []byte{}
	for {
		flag := false
		for i, v := range in {
			if all[i] == 2 && v == 0 && out[i] {
				ans = append(ans, byte(i+'a'))
				all[i] = 1
				for x, xt := range e[i] {
					if xt {
						in[x]--
						if in[x] == 0 {
							flag = true
						}
					}
				}
			}
		}
		if !flag {
			break
		}
	}

	//fmt.Println(in)
	for _, v := range in {
		if v > 0 {
			return ""
		}
	}
	for i, v := range all {
		if v == 2 {
			ans = append(ans, byte(i+'a'))
		}
	}

	return string(ans)
}

package main

import "fmt"

func main() {
	fmt.Println(numDifferentIntegers("a1b01c001"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func numDifferentIntegers(s string) int {
	m := map[string]bool{}
	t := []byte{}
	for _, v := range s {
		if v >= '0' && v <= '9' {
			t = append(t, byte(v))
		} else {
			if len(t) > 0 {
				if len(t) == 1 && t[0] == '0' {
					m["0"] = true
				} else {
					for i := 0; i < len(t); i++ {
						if t[i] == '0' {
							if i == len(t)-1 {
								m["0"] = true
							}
						} else {
							m[string(t[i:])] = true
							break
						}
					}
				}
			}
			t = []byte{}
		}
	}
	if len(t) > 0 {
		if len(t) == 1 && t[0] == '0' {
			m["0"] = true
		} else {
			for i := 0; i < len(t); i++ {
				if t[i] == '0' {
					if i == len(t)-1 {
						m["0"] = true
					}
					continue
				} else {
					m[string(t[i:])] = true
					break
				}
			}
		}
	}

	return len(m)
}

//leetcode submit region end(Prohibit modification and deletion)

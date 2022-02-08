package main

import (
	"fmt"
	"sort"
	"strconv"
)

// https://leetcode-cn.com/leetbook/read/bytedance-c01/euzitj/
func main() {
	var input string
	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			return
		}
		fmt.Println(parse(input))
	}
}

func parse(s string) string {
	st := []map[string]int{{}}
	for i := 0; i < len(s); i++ {
		v, cur := s[i], len(st)-1
		if v >= 'A' && v <= 'Z' {
			x, cnt := []byte{v}, 0
			for i += 1; i < len(s); i++ {
				if s[i] >= 'a' && s[i] <= 'z' {
					x = append(x, s[i])
				} else if s[i] >= '0' && s[i] <= '9' {
					cnt = cnt*10 + int(s[i]-'0')
				} else {
					break
				}
			}
			i--

			if cnt == 0 {
				cnt = 1
			}
			st[cur][string(x)] += cnt
		}
		if v == '(' || v == '[' {
			cur++
			st = append(st, map[string]int{})
		}
		if v == ')' || v == ']' {
			cnt := 0
			for i += 1; i < len(s); i++ {
				if s[i] >= '0' && s[i] <= '9' {
					cnt = cnt*10 + int(s[i]-'0')
				} else {
					break
				}
			}
			i--
			if cnt == 0 {
				cnt = 1
			}

			for kk, vv := range st[cur] {
				st[cur-1][kk] += vv * cnt
			}
			cur--
			st = st[:len(st)-1]
		}
	}

	field := []string{}
	for k, _ := range st[0] {
		field = append(field, k)
	}
	sort.Strings(field)

	ans := ""
	for _, v := range field {
		ans += v
		ans += strconv.Itoa(st[0][v])
	}

	return ans
}

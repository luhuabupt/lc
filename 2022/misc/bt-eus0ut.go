package main

import (
	"fmt"
	"sort"
	"strings"
)

// https://leetcode-cn.com/leetbook/read/bytedance-c01/eus0ut/
func main() {
	var input string
	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			return
		}
		input = strings.TrimLeft(input, "[")
		input = strings.TrimRight(input, "]")

		s := strings.Split(input, ",")

		fmt.Println(do(s))
	}
}

func do(s []string) (ans string) {
	sort.Slice(s, func(i, j int) bool {
		a, b := s[i], s[j]
		for x := 0; x < len(a) && x < len(b); x++ {
			if a[x] != b[x] {
				return a[x] > b[x]
			}
		}
		xa, xb := a[0], b[0]
		if len(a) > len(b) {
			xa = a[len(b)]
		} else if len(b) > len(a) {
			xb = b[len(a)]
		}
		return xa > xb
	})
	for _, v := range s {
		ans += v
	}

	return ans
}

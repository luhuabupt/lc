package main

import (
	"fmt"
)

func main() {
	//fmt.Println(countEven(4))
	fmt.Println(repeatLimitedString("cczazcc", 3))
	fmt.Println(repeatLimitedString("aabbdd", 2))
}

func repeatLimitedString(s string, lim int) string {
	cnt := [26]int{}
	for _, v := range s {
		cnt[v-'a']++
	}
	//fmt.Println(cnt)

	get := func(c, x int) (a int, v int) {
		for i := 25; i >= 0; i-- {
			if i == x {continue}
			if cnt[i] > 0 {
				if cnt[i] > c {
					cnt[i] -= c
					return i, c
				} else {
					v = cnt[i]
					cnt[i] = 0
					return i, v
				}
			}
		}
		return
	}

	ans := []byte{}
	k, v := 0, 0
	for {
		k, v = get(lim, -1)
		if v == 0 {
			break
		}
		for i := 0; i < v; i++ {
			ans = append(ans, byte('a'+k))
		}
		if cnt[k] == 0 {
			continue
		}

		k, v = get(1, k)
		if v == 0 {
			break
		}
		ans = append(ans, byte('a'+k))

		//fmt.Println(cnt)
	}
	return string(ans)
}

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	t := Constructor([]byte{'a', 'b', 'c', 'd'}, []string{"ei", "zf", "ei", "am"}, []string{"abcd", "acbd", "adbc", "badc", "dacb", "cadb", "cbda", "abad"})
	fmt.Println(t.Encrypt("abcd"))
	fmt.Println(t.Decrypt("eizfeiam"))
	fmt.Println(t.Decrypt("eizfeiei"))

	t2 := Constructor([]byte{'a', 'b', 'c', 'd'}, []string{"ei", "zf", "ei", "am"}, []string{"aaaa"})
	fmt.Println(t2.Encrypt("abcd"))
	fmt.Println(t2.Decrypt("eizfeiam"))
}

type Encrypter struct {
	kv map[byte]string
	vk map[string][]byte
	t  *tire
}

type tire struct {
	sub [26]*tire
	end bool
}

func Constructor(keys []byte, values []string, dictionary []string) Encrypter {
	kv := map[byte]string{}
	vk := map[string][]byte{}

	for i, v := range keys {
		kv[v] = values[i]
		if vk[values[i]] == nil {
			vk[values[i]] = []byte{}
		}
		vk[values[i]] = append(vk[values[i]], v)
	}

	d := &tire{}
	for _, s := range dictionary {
		p := d
		for _, v := range s {
			if p.sub[v-'a'] == nil {
				p.sub[v-'a'] = &tire{}
			}
			p = p.sub[v-'a']
		}
		p.end = true
	}

	return Encrypter{kv, vk, d}
}

func (t *Encrypter) Encrypt(w string) string {
	ans := []byte{}
	for _, v := range w {
		ans = append(ans, t.kv[byte(v)]...)
	}
	return string(ans)
}

func (t *Encrypter) Decrypt(w string) int {
	n := len(w)
	sa := []*tire{t.t}
	for i := 0; i < n; i += 2 {
		xa := t.vk[w[i:i+2]]
		if len(xa) == 0 || len(sa) == 0 {
			return 0
		}

		ns := []*tire{}
		for _, p := range sa {
			for _, v := range xa {
				if p.sub[v-'a'] != nil {
					ns = append(ns, p.sub[v-'a'])
				}
			}
		}
		sa = ns
	}

	ans := 0
	for _, x := range sa {
		if x.end == true {
			ans++
		}
	}

	return ans
}

/**
 * Your Encrypter object will be instantiated and called as such:
 * obj := Constructor(keys, values, dictionary);
 * param_1 := obj.Encrypt(word1);
 * param_2 := obj.Decrypt(word2);
 */

func convertTime(ss string, tt string) int {
	gt := func(v string) int {
		a := strings.Split(v, ":")
		h, _ := strconv.Atoi(a[0])
		m, _ := strconv.Atoi(a[1])
		return h*60 + m
	}

	s := gt(ss)
	t := gt(tt)

	//fmt.Println(s, t)
	ans := 0

loop:
	for s < t {
		//fmt.Println(s)
		for _, x := range []int{60, 15, 5, 1} {
			if t-s >= x {
				s += x
				ans++
				continue loop
			}
		}
	}
	return ans
}

func findWinners(m [][]int) [][]int {
	lose, do := make([]int, 100001), make([]int, 100001)
	for _, v := range m {
		lose[v[1]]++
		do[v[0]]++
		do[v[1]]++
	}

	ans := make([][]int, 2)
	for i, v := range do {
		if v > 0 {
			if lose[i] == 0 {
				ans[0] = append(ans[0], i)
			}
			if lose[i] == 1 {
				ans[1] = append(ans[1], i)
			}
		}
	}

	return ans
}

func maximumCandies(c []int, k int64) int {
	return sort.Search(1e7+1, func(r int) bool {
		do := 0
		for _, v := range c {
			do += v / r
		}
		return int64(do) < k
	}) - 1
}

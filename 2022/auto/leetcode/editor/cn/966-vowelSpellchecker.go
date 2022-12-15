package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func spellchecker(w []string, q []string) []string {
	a := map[string]bool{}
	b := map[string]string{}
	c := map[string]string{}
	m := map[int32]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

	low := func(v string) string {
		d := []byte{}
		for _, x := range v {
			if x >= 'A' && x <= 'Z' {
				d = append(d, byte(x-'A'+'a'))
			} else {
				d = append(d, byte(x))
			}
		}
		return string(d)
	}
	vowel := func(v string) string {
		d := []byte{}
		for _, x := range v {
			nx := x
			if nx >= 'A' && nx <= 'Z' {
				nx = x - 'A' + 'a'
			}
			if m[nx] {
				nx = '*'
			}
			d = append(d, byte(nx))
		}
		return string(d)
	}

	for _, v := range w {
		a[v] = true
		if _, ok := b[low(v)]; !ok {
			b[low(v)] = v
		}
		if _, ok := c[vowel(v)]; !ok {
			c[vowel(v)] = v
		}
	}

	ans := make([]string, len(q))
	for i, v := range q {
		if a[v] {
			ans[i] = v
		} else if b[low(v)] != "" {
			ans[i] = b[low(v)]
		} else if c[vowel(v)] != "" {
			ans[i] = c[vowel(v)]
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

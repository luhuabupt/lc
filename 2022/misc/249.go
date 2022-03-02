package main

func groupStrings(s []string) [][]string {
	m := map[int]map[string][]string{}

	toa := func(v rune, d int) byte {
		x := int(v-'a') - d
		if x < 0 {
			x += 26
		}
		return byte('a' + x)
	}

	for _, v := range s {
		d := int(v[0] - 'a')
		fa := []byte{}
		for _, x := range v {
			fa = append(fa, toa(x, d))
		}
		if m[len(v)] == nil {
			m[len(v)] = map[string][]string{}
		}
		if m[len(v)][string(fa)] == nil {
			m[len(v)][string(fa)] = []string{}
		}
		m[len(v)][string(fa)] = append(m[len(v)][string(fa)], v)
	}

	ans := [][]string{}
	for _, la := range m {
		for _, sa := range la {
			ans = append(ans, sa)
		}
	}

	return ans
}

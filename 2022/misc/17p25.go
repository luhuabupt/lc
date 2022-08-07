package main

func maxRectangle(words []string) []string {
	max := 0
	ans := []string{}

	type t struct {
		sub [26]*t
		end bool
	}

	hash := make([]map[string]bool, 101)
	tire := make([]t, 101)
	for _, v := range words {
		hash[len(v)][v] = true
		if len(v) > max {
			max = len(v)
			ans = []string{v}
		}
		p := tire[len(v)]
		for i := 0; i < len(v); i++ {
			p.sub[v[i]-'a'] = nil
			p = *p.sub[v[i]-'a']
		}
		p.end = true
	}

	for n, mp := range hash {
		for m := 2; m <= n && m <= len(mp); m++ { // 只算长的矩形
			// 剪枝
			if len(hash[m]) < n { // 竖行不够
				continue
			}
			if m*n <= max {
				continue
			}
			// 字典树
		}
	}

	return ans
}

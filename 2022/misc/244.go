package main

type WordDistance struct {
	pos map[string][]int
}

func Constructor(w []string) WordDistance {
	p := map[string][]int{}
	for i, v := range w {
		p[v] = append(p[v], i)
	}
	return WordDistance{p}
}

func (t *WordDistance) Shortest(w1 string, w2 string) int {
	a1 := t.pos[w1]
	a2 := t.pos[w2]

	get := func(ans, i, j int) int {
		if i < j {
			i, j = j, i
		}
		if i-j < ans {
			ans = i - j
		}
		return ans
	}

	ans := 100000
	for i, j := 0, 0; i < len(a1) && j < len(a2); i++ {
		for ; j < len(a2) && a2[j] < a1[i]; j++ {
			ans = get(ans, a1[i], a2[j])
		}
		if j < len(a2) {
			ans = get(ans, a1[i], a2[j])
		}
	}

	return ans
}

/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(wordsDict);
 * param_1 := obj.Shortest(word1,word2);
 */

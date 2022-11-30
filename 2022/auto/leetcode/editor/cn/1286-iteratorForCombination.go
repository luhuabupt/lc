package main

import "fmt"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
type CombinationIterator struct {
	a   [][]byte
	idx int
}

func Constructor(cs string, l int) CombinationIterator {
	a := [][]byte{}
	c := []byte(cs)

	var dfs func(cur, left []byte)
	dfs = func(cur, left []byte) {
		if len(a) > 1e4 {
			return
		}
		if len(cur) == l {
			a = append(a, cur)
			return
		}
		for i, x := range left {
			if len(cur)+len(left)-i < l {
				continue
			}
			nc := append(append([]byte{}, cur...), x)
			nf := []byte{}
			for j := i + 1; j < len(left); j++ {
				nf = append(nf, left[j])
			}
			dfs(nc, nf)
		}
	}
	dfs([]byte{}, c)

	return CombinationIterator{a, -1}
}

func (t *CombinationIterator) Next() string {
	t.idx++
	return string(t.a[t.idx])
}

func (t *CombinationIterator) HasNext() bool {
	return t.idx < len(t.a)-1
}

/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
//leetcode submit region end(Prohibit modification and deletion)

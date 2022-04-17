package main

import "math/rand"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
type RandomizedSet struct {
	a   []int
	m   map[int]int
	end int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		a:   make([]int, 2*1e5+1),
		m:   map[int]int{},
		end: 0,
	}
}

func (t *RandomizedSet) Insert(val int) bool {
	if _, ok := t.m[val]; ok {
		return false
	}

	t.a[t.end] = val
	t.m[val] = t.end
	t.end++

	return true
}

func (t *RandomizedSet) Remove(val int) bool {
	idx, ok := t.m[val]
	if !ok {
		return false
	}

	delete(t.m, val)
	t.end--
	if t.end != idx {
		t.m[t.a[t.end]] = idx
		t.a[idx] = t.a[t.end]
	}

	return true
}

func (t *RandomizedSet) GetRandom() int {
	return t.a[rand.Intn(t.end)]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
//leetcode submit region end(Prohibit modification and deletion)

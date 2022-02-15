package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
type Trie struct {
	ch    [26]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	for _, v := range word {
		if t.ch[v-'a'] == nil {
			t.ch[v-'a'] = &Trie{}
		}
		t = t.ch[v-'a']
	}
	t.isEnd = true
}

func (t *Trie) Search(word string) bool {
	for _, v := range word {
		if t.ch[v-'a'] == nil {
			return false
		}
		t = t.ch[v-'a']
	}
	return t.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		if t.ch[v-'a'] == nil {
			return false
		}
		t = t.ch[v-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
//leetcode submit region end(Prohibit modification and deletion)

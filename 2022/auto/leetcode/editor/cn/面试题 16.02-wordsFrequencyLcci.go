package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
type WordsFrequency struct {
	m map[string]int
}

func Constructor(book []string) WordsFrequency {
	cnt := map[string]int{}
	for _, v := range book {
		cnt[v]++
	}
	return WordsFrequency{cnt}
}

func (this *WordsFrequency) Get(word string) int {
	return this.m[word]
}

/**
 * Your WordsFrequency object will be instantiated and called as such:
 * obj := Constructor(book);
 * param_1 := obj.Get(word);
 */
//leetcode submit region end(Prohibit modification and deletion)

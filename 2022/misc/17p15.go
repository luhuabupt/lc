package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(longestWord([]string{"qlmql","qlmqlmqqlqmqqlq","mqqlqmqqlqmqqlq","mqqlq","mqqlqlmlsmqq","qmlmmmmsm","lmlsmqq","slmsqq","mslqssl","mqqlqmqqlq"}))
	fmt.Println(longestWord([]string{"cat", "banana", "dog", "nana", "walk", "walker", "dogwalker"}))
}

func longestWord(words []string) string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j]) || len(words[i]) == len(words[j]) && words[i] < words[j]
	})

	m := map[string]bool{}

	var chk func(s string) bool
	chk = func(s string) bool {
		if m[s] {
			return true
		}
		for i := 0; i < len(s)-1; i++ {
			if m[s[:i+1]] && chk(s[i+1:]) {
				return true
			}
		}
		return false
	}

	ans := ""
	for _, s := range words {
		if chk(s) {
			ans = s
		}
		m[s] = true
	}

	return ans
}

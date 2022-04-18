package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(123)
	fmt.Println(deleteText("Singing dancing in the rain", 10))
	fmt.Println(deleteText("Hello World", 5))
	fmt.Println(deleteText("Hello World", 2))
}

func deleteText(a string, idx int) string {
	do := 0
	for i, v := range a {
		if v == ' ' {
			do++
		}
		if i == idx {
			if v == ' ' {
				return a
			} else {
				x := strings.Split(a, " ")
				re := x[:do]
				re = append(re, x[do+1:]...)
				return strings.Join(re, " ")
			}
		}
	}
	return ""
}

func numFlowers(r [][]int) int {
	in := map[int]int{}
	for _, v := range r {
		in[v[0]]++
		in[v[1]]++
	}
	ans := 0
	for i, v := range in {
		if v > in[ans] {
			ans = i
		}
	}
	return in[ans] + 1
}

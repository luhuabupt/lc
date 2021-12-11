package main

import "fmt"

type Abc struct {
	node []*Abc
}

func main() {
	a := []int{1, 2, 3}
	for _, v := range a {
		a = append(a, 100+v)
		a = append(a, 10+v)
	}
	fmt.Println(a)
}

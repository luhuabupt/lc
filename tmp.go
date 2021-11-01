package main

import "fmt"

type Abc struct {
	node []*Abc
}

func main() {
	for j := 100000; j > 0; j -= (j & -j) {
		fmt.Println(j)
		fmt.Printf("%#b ", j&-j)
		fmt.Printf("%#b \n", j)
	}
}

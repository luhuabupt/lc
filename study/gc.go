package main

import (
	"fmt"
)

func main() {
	a := map[int]int{1: 2}
	ss()
	//runtime.GC()
	fmt.Println(a)
}

func ss() {
	b := map[int]int{1: 2}
	b[2] = 1
	for i := 0; i < 1000000; i++ {
		b[i] = i%7 + i%8 + i%9
	}
}

package main

import (
	"errors"
	"fmt"
)

func main() {
	var err error
	err = errors.New(" test")
	fmt.Println(err)

	a := []string{}
	b := append(a[:0:0], a...)
}

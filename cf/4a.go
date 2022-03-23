package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			return
		}
		fmt.Println(check(input))
	}
}

func check (input string) string {
	i, _ := strconv.Atoi(input)

	if i % 2 == 0 && i > 2 {
		return "YES"
	}
	return "NO"
}
package main

import "fmt"

func main() {
	fmt.Println(lcm(2, 12))
	fmt.Println(lcm(2, 3))
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	x := gcd(a, b)
	return a / x * b
}

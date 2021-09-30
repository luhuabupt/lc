package main

import "fmt"

func main() {
	fmt.Println(bubble([]int{3,1,5,4,3}))
}

func bubble(arr []int) []int {
	for i, _ := range arr {
		for j := i+1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

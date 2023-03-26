package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		test()
	}
}

func test() {
	ch := make(chan int, 2)
	go func() { ch <- 3 }()
	go func() { ch <- 4 }()

	ans := []int{}
	for len(ans) < 2 {
		select {
		case i := <-ch:
			ans = append(ans, i)
		case i := <-ch:
			ans = append(ans, i*2)
		}
	}
	close(ch)

	fmt.Println(ans)
}

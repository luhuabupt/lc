package main

func main() {
	ch := make(chan int, 2)
	ch <- 3
	ch <- 3

	select {
	case i := <-ch:
		println(i)
	case i := <-ch:
		println(i + 1)
	default:
		println("default")
	}
}

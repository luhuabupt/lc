package main

func main() {
	defer println("in main")

	defer func() {
		recover()
	}()

	defer func() {
		err := recover()
		println(err)
		defer func() {
			recover()
			panic("panic again and again")
		}()
		panic("panic again")
	}()

	panic("panic once")
}

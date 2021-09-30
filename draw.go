package main

import (
	"fmt"
	"math/rand"
)

func main() {
	seed := int64(2000)
	rand.Seed(seed)
	rand := rand.Intn(9)
	fmt.Println(rand)
}

func draw(seed, total int) {

}

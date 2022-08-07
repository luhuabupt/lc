package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type T struct {
	wg sync.WaitGroup
	ch chan bool
}

func main() {
	// resource
	data := []int{}
	for i := 0; i < 1000000; i++ {
		data = append(data, i)
	}

	// new
	t := new(T)
	t.ch = make(chan bool, 100)

	for len(data) > 0 {
		// pop
		x := data[0]
		data = data[1:]

		t.ch <- true
		t.wg.Add(1)
		go t.exec(x)
	}

	t.wg.Wait()
	close(t.ch)
}

func (t *T) exec(x int) {
	defer func() {
		_ = <-t.ch
		t.wg.Done()
	}()

	// do smt
	nano := time.Now().UnixNano()
	rand.Seed(nano)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)

	fmt.Println(x)
}

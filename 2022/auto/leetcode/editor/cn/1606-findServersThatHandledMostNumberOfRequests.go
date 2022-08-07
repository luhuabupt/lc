package main

import (
	"container/heap"
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
)

func main() {
	//fmt.Println(busiestServers(3, []int{1, 2, 3, 4, 5}, []int{5, 2, 3, 3, 3}))
	//fmt.Println(busiestServers(3, []int{1}, []int{100000}))
	//fmt.Println(busiestServers(3, []int{3, 4, 6, 8, 9, 11, 12, 16}, []int{1, 2, 8, 6, 5, 3, 8, 3}))  // 1
	fmt.Println(busiestServers(7, []int{1, 3, 4, 5, 6, 11, 12, 13, 15, 19, 20, 21, 23, 25, 31, 32}, []int{9, 16, 14, 1, 5, 15, 6, 10, 1, 1, 7, 5, 11, 4, 4, 6}))  // 0
	fmt.Println(busiestServers_(7, []int{1, 3, 4, 5, 6, 11, 12, 13, 15, 19, 20, 21, 23, 25, 31, 32}, []int{9, 16, 14, 1, 5, 15, 6, 10, 1, 1, 7, 5, 11, 4, 4, 6})) // 0
	//fmt.Println(busiestServers_(3, []int{3, 4, 6, 8, 9, 11, 12, 16}, []int{1, 2, 8, 6, 5, 3, 8, 3})) // 1
}

//leetcode submit region begin(Prohibit modification and deletion)
type ht [][2]int

func (h ht) Len() int            { return len(h) }
func (h ht) Less(i, j int) bool  { return h[i][0] < h[j][0] } // 小顶堆
func (h ht) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *ht) Push(v interface{}) { *h = append(*h, v.([2]int)) }
func (h *ht) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

type h []int

func (h h) Len() int            { return len(h) }
func (h h) Less(i, j int) bool  { return h[i] < h[j] } // 小顶堆
func (h h) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *h) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *h) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func busiestServers(k int, a []int, load []int) []int {
	n := len(a)
	do := make([]int, k)
	free := &h{}

	h := &ht{}
	for i := 0; i < k && i < n; i++ {
		heap.Push(h, [2]int{a[i] + load[i], i})
		do[i]++
	}

	for i := k; i < n; i++ {
		for h.Len() > 0 {
			cur := heap.Pop(h).([2]int)
			if cur[0] <= a[i] {
				in := cur[1]
				if in < i%k {
					in += k
				}
				heap.Push(free, in)
			} else {
				heap.Push(h, cur)
				break
			}
		}

		if free.Len() == 0 {
			continue
		}

		use := 0
		for {
			cur := heap.Pop(free).(int)
			if cur >= i {
				use = cur % k
				break
			}
			heap.Push(free, cur+k)
		}

		heap.Push(h, [2]int{a[i] + load[i], use})
		do[use]++
	}

	ma, ans := 0, []int{}
	for _, v := range do {
		if v > ma {
			ma = v
		}
	}
	for i, v := range do {
		if v == ma {
			ans = append(ans, i)
		}
	}

	return ans
}

func busiestServers_(k int, a []int, load []int) []int {
	n := len(a)
	do := make([]int, k)
	free := redblacktree.NewWithIntComparator()

	h := &ht{}
	for i := 0; i < k && i < n; i++ {
		heap.Push(h, [2]int{a[i] + load[i], i})
		do[i]++
	}

	for i := k; i < n; i++ {
		for h.Len() > 0 {
			cur := heap.Pop(h).([2]int)
			if cur[0] <= a[i] {
				free.Put(cur[1], nil)
			} else {
				heap.Push(h, cur)
				break
			}
		}

		if free.Size() == 0 {
			continue
		}

		use, _ := free.Ceiling(i % k)
		if use == nil {
			use = free.Left()
		}

		heap.Push(h, [2]int{a[i] + load[i], use.Key.(int)})
		do[use.Key.(int)]++
		fmt.Println(i, use.Key.(int), free.Size())
		free.Remove(use.Key)
	}

	ma, ans := 0, []int{}
	for _, v := range do {
		if v > ma {
			ma = v
		}
	}
	for i, v := range do {
		if v == ma {
			ans = append(ans, i)
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

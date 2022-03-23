package main

import "container/heap"

func main() {
	
}

//leetcode submit region begin(Prohibit modification and deletion)
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

func magicTower(nums []int) int {
	sum := 1
	for _,x := range nums {
		sum += x
	}
	if sum <= 0 {
		return -1
	}

	h := &h{}

	ans := 0
	cur := 1
	for _, v := range nums {
		if v < 0 {
			heap.Push(h, v)
		}
		cur += v

		if cur <= 0 {
			x := heap.Pop(h).(int)
			ans++

			cur -= x
		}
	}

	return ans
}
//leetcode submit region end(Prohibit modification and deletion)

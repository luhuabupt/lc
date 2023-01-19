package main

import (
	"container/heap"
	"fmt"
)

// number-of-orders-in-the-backlog 积压订单中的订单总数  2023-01-02 08:40:45
func main() {
	fmt.Println(getNumberOfBacklogOrders([][]int{{10, 5, 0}, {15, 2, 1}, {25, 1, 1}, {30, 4, 0}}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func getNumberOfBacklogOrders(a [][]int) int {
	sell := &ht{}
	buy := &hl{}
	M := int(1e9) + 7
	ans := 0

	for _, v := range a {
		if v[2] == 0 { // buy
			for sell.Len() > 0 && v[1] > 0 {
				x := heap.Pop(sell).([2]int)
				if x[0] <= v[0] {
					d := min(v[1], x[1])
					v[1] -= d
					x[1] -= d
					if x[1] > 0 {
						heap.Push(sell, x)
						break
					}
				} else {
					heap.Push(sell, x)
					break
				}
			}
			if v[1] > 0 {
				heap.Push(buy, [2]int{v[0], v[1]})
			}
		} else { // sell
			for buy.Len() > 0 && v[1] > 0 {
				x := heap.Pop(buy).([2]int)
				if x[0] >= v[0] {
					d := min(v[1], x[1])
					v[1] -= d
					x[1] -= d
					if x[1] > 0 {
						heap.Push(buy, x)
						break
					}
				} else {
					heap.Push(buy, x)
					break
				}
			}
			if v[1] > 0 {
				heap.Push(sell, [2]int{v[0], v[1]})
			}
		}
	}
	for sell.Len() > 0 {
		x := heap.Pop(sell).([2]int)
		ans += x[1]
	}
	for buy.Len() > 0 {
		x := heap.Pop(buy).([2]int)
		ans += x[1]
	}

	return ans % M
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

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

type hl [][2]int

func (h hl) Len() int            { return len(h) }
func (h hl) Less(i, j int) bool  { return h[i][0] > h[j][0] } // 大顶堆
func (h hl) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hl) Push(v interface{}) { *h = append(*h, v.([2]int)) }
func (h *hl) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)

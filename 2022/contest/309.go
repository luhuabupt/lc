package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestNiceSubarray([]int{1, 3, 8, 48, 10}))
	fmt.Println(longestNiceSubarray([]int{135745088, 609245787, 16, 2048, 2097152}))
	fmt.Println(longestNiceSubarray([]int{536870913, 536870914}))
}

func mostBooked_(n int, ar [][]int) int {
	do := make([]int, n)
	e := make([]int, n)
	sort.Slice(ar, func(i, j int) bool {
		return ar[i][0] < ar[j][0]
	})
	for i := 0; i < n; i++ {
		e[i] = 0
	}
loop:
	for _, v := range ar {
		mj := 0
		for j, ev := range e {
			if ev <= v[0] {
				do[j]++
				e[j] = v[1]
				continue loop
			}
			if ev < e[mj] {
				mj = j
			}
		}
		do[mj]++
		e[mj] = e[mj] + v[1] - v[0]
	}

	ans := 0
	for i, v := range do {
		if v > do[ans] {
			ans = i
		}
	}
	return ans
}

func mostBooked(n int, ar [][]int) int {
	do := make([]int, n) // 预订次数
	sort.Slice(ar, func(i, j int) bool {
		return ar[i][0] < ar[j][0]
	})
	on := &ht{}  // 开会中 [2]int: 记录结束时间和房间编号
	free := &h{} // 空闲 int 记录房间编号
	for i := 0; i < n; i++ {
		heap.Push(free, i)
	}
	for _, v := range ar {
		for on.Len() > 0 { // 处理当前空闲的会议室
			x := heap.Pop(on).([2]int)
			if x[0] <= v[0] {
				heap.Push(free, x[1])
			} else {
				heap.Push(on, x)
				break
			}
		}
		if free.Len() > 0 {
			cur := heap.Pop(free).(int)
			do[cur]++
			heap.Push(on, [2]int{v[1], cur})
		} else { // 没空闲的，等一个最近结束的
			wait := heap.Pop(on).([2]int) //
			do[wait[1]]++
			heap.Push(on, [2]int{wait[0] + v[1] - v[0], wait[1]})
		}
	}
	ans := 0
	for i, v := range do {
		if v > do[ans] {
			ans = i
		}
	}
	return ans
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

type ht [][2]int

func (h ht) Len() int            { return len(h) }
func (h ht) Less(i, j int) bool  { return h[i][0] < h[j][0] || h[i][0] == h[j][0] && h[i][1] < h[j][1] } // 小顶堆
func (h ht) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *ht) Push(v interface{}) { *h = append(*h, v.([2]int)) }
func (h *ht) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func longestNiceSubarray(ar []int) int {
	n := len(ar)
	ans := 1

loop:
	for i, _ := range ar {
		c := [31]bool{}
		flag := true
		for j := i; j < n && j <= i+30; j++ {
			//fmt.Println(i, j)
			//fmt.Printf("%b\n", ar[j])
			for b := 0; 1<<b < 1<<31; b++ {
				if (1<<b)&ar[j] > 0 {
					if c[b] == true {
						flag = false
						continue loop
					}
					c[b] = true
				}
			}
			if flag && j-i+1 > ans {
				//fmt.Println("ans", i, j, j-i+1)
				ans = j - i + 1
			}
		}
	}
	return ans

	//ans := 1
	//for i, _ := range ar {
	//	l := i
	//	r := i + 30
	//	for ; l < r; l++ {
	//		for j := l; j < n && j <= r; j++ {
	//			if ar[l]&ar[j] != 0 {
	//				r = j - 1
	//				break
	//			}
	//		}
	//	}
	//	if r -i+1 > ans {
	//		ans = r-i+1
	//	}
	//}
	//return ans
}

func checkDistances(s string, d []int) bool {
	ar := [26][]int{}
	for i, v := range s {
		ar[v-'a'] = append(ar[v-'a'], i)
	}
	for i, v := range ar {
		if len(v) == 2 {
			if v[1]-v[0] != d[i] {
				return false
			}
		}
	}
	return true
}

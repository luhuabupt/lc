package main

import (
	"container/heap"
	"github.com/emirpasic/gods/trees/redblacktree"
)

func main() {
	freqStack := Constructor()
	freqStack.Push(5) //堆栈为 [5]
	freqStack.Push(7) //堆栈是 [5,7]
	freqStack.Push(5) //堆栈是 [5,7,5]
	freqStack.Push(7) //堆栈是 [5,7,5,7]
	freqStack.Push(4) //堆栈是 [5,7,5,7,4]
	freqStack.Push(5) //堆栈是 [5,7,5,7,4,5]
	freqStack.Pop()   //返回 5 ，因为 5 出现频率最高。堆栈变成 [5,7,5,7,4]。
	freqStack.Pop()   //返回 7 ，因为 5 和 7 出现频率最高，但7最接近顶部。堆栈变成 [5,7,5,4]。
	freqStack.Pop()   //返回 5 ，因为 5 出现频率最高。堆栈变成 [5,7,4]。
	freqStack.Pop()   //返回 4 ，因为 4, 5 和 7 出现频率最高，但 4 是最接近顶部的。堆栈变成 [5,7]。
}

//leetcode submit region begin(Prohibit modification and deletion)
type FreqStack struct {
	vm  map[int][]int
	rbt *redblacktree.Tree
	idx int
}

func Constructor() FreqStack {
	rbt := redblacktree.NewWith(func(ai, bi interface{}) int {
		a := ai.([]int)
		b := bi.([]int)
		if len(a) < len(b) || len(a) == len(b) && a[len(a)-1] < b[len(b)-1] {
			return 1
		}
		if len(a) > len(b) || len(a) == len(b) && a[len(a)-1] > b[len(b)-1] {
			return -1
		}
		return 0
	})

	return FreqStack{map[int][]int{}, rbt, 0}
}

func (t *FreqStack) Push(val int) {
	t.rbt.Remove(t.vm[val])
	t.vm[val] = append(t.vm[val], t.idx)
	t.rbt.Put(t.vm[val], val)

	//fmt.Println("Push =+++", val,  t.rbt.String())
	t.idx++
}

func (t *FreqStack) Pop() int {
	r := t.rbt.Left()
	ans := r.Value.(int)
	t.rbt.Remove(t.vm[ans])

	na := t.vm[ans]
	na = na[:len(na)-1]

	t.vm[ans] = na
	if len(na) > 0 {
		t.rbt.Put(t.vm[ans], ans)
	}
	//fmt.Println("Pop =+++",  t.rbt.String())

	return ans
}

//type ht [][]int
//
//func (h ht) Len() int { return len(h) }
//func (h ht) Less(i, j int) bool {
//	return len(h[i]) > len(h[j]) || len(h[i]) == len(h[j]) && h[i][len(h[i])-1] > h[j][len(h[j])-1]
//} // 大顶堆
//func (h ht) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
//func (h *ht) Push(v interface{}) { *h = append(*h, v.([]int)) }
//func (h *ht) Pop() interface{} {
//	a := *h
//	v := a[len(a)-1]
//	*h = a[:len(a)-1]
//	return v
//}
//
//type FreqStack struct {
//	vm  map[int][]int
//	h   *ht
//	idx int
//}
//
//func Constructor() FreqStack {
//	return FreqStack{map[int][]int{}, &ht{}, 0}
//}
//
//func (t *FreqStack) Push(val int) {
//	if len(t.vm[val]) == 0 {
//		t.vm[val] = []int{val}
//	}
//	t.vm[val] = append(t.vm[val], t.idx)
//	heap.Push(t.h, t.vm[val])
//	t.idx++
//}
//
//func (t *FreqStack) Pop() int {
//	x := heap.Pop(t.h).([]int)
//	t.vm[x[0]] = t.vm[x[0]][:len(t.vm[x[0]])-1]
//	return x[0]
//}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
//leetcode submit region end(Prohibit modification and deletion)

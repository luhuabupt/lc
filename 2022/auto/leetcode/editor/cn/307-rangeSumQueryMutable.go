package main

import "fmt"

// range-sum-query-mutable 区域和检索 - 数组可修改  2022-04-04 22:02:22
func main() {
	fmt.Println(rangeSumQueryMutable())
}

//leetcode submit region begin(Prohibit modification and deletion)
type NumArray struct {
	tr *tree
	a  []int
}

type tree struct {
	arr []int
}

func (t tree) add(i, v int) {
	for i < len(t.arr) {
		t.arr[i] = t.arr[i] + v
		i += lowbit(i)
	}
}

func (t tree) sum(i int) (ans int) {
	for i > 0 {
		ans += t.arr[i]
		i -= lowbit(i)
	}
	return
}

func (t tree) query(l, r int) (ans int) {
	return t.sum(r) - t.sum(l-1)
}

func lowbit(v int) int {
	return v & -v
}

func Constructor(nums []int) NumArray {
	t := tree{make([]int, len(nums)+1)}
	for i, v := range nums {
		t.add(i+1, v)
	}
	return NumArray{&t, nums}
}

func (t *NumArray) Update(index int, val int) {
	t.tr.add(index+1, val-t.a[index])
	t.a[index] = val
}

func (t *NumArray) SumRange(left int, right int) int {
	return t.tr.query(left+1, right+1)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
//leetcode submit region end(Prohibit modification and deletion)

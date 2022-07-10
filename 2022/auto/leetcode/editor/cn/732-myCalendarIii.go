package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)

type MyCalendarThree struct {
}

func Constructor() MyCalendarThree {

}

func (this *MyCalendarThree) Book(start int, end int) int {

}

//type MyCalendarThree struct {
//	*redblacktree.Tree
//}
//
//func Constructor() MyCalendarThree {
//	return MyCalendarThree{redblacktree.NewWithIntComparator()}
//}
//
//func (t *MyCalendarThree) Book(s int, e int) int {
//	t.add(s, 1)
//	t.add(e, -1)
//
//	ans, cur := 0, 0
//	for v := t.Iterator(); v.Next(); {
//		cur += v.Value().(int)
//		if cur > ans {
//			ans = cur
//		}
//	}
//
//	return ans
//}
//
//func (t *MyCalendarThree) add(idx, d int) {
//	if v, ok := t.Get(idx); ok {
//		d += v.(int)
//	}
//	t.Put(idx, d)
//}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
//leetcode submit region end(Prohibit modification and deletion)

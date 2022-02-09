package main

// my-calendar-i 我的日程安排表 I  2022-01-28 23:06:15
func main() {
}

//leetcode submit region begin(Prohibit modification and deletion)
type MyCalendar struct {
	arr [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{[][]int{}}
}

func (t *MyCalendar) Book(start int, end int) bool {
	i := 0
	for ; i < len(t.arr); i++ {
		a := t.arr[i]
		if a[0] <= start && a[1] > start || a[0] >= start && a[0] < end {
			return false
		}
		if a[0] >= end {
			break
		}
	}

	nr := [][]int{}
	nr = append(nr, t.arr[:i]...)
	nr = append(nr, []int{start, end})
	nr = append(nr, t.arr[i:]...)
	t.arr = nr

	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
//leetcode submit region end(Prohibit modification and deletion)

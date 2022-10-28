package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func countStudents(st []int, sand []int) int {
	c := make([]int, 2)
	for _, v := range st {
		c[v]++
	}
	j := 0
	for i, v := range sand {
		if c[v] == 0 {
			return len(sand) - i
		}
		for j < len(st) {
			if st[j] != v {
				j++
				st = append(st, st[j])
			} else {
				c[v]--
				break
			}
		}
	}
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)

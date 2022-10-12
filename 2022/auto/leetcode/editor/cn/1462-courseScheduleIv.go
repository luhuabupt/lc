package main

// course-schedule-iv 课程表 IV  2022-09-25 13:43:23
func main() {
	//fmt.Println(checkIfPrerequisite())
}

//leetcode submit region begin(Prohibit modification and deletion)
func checkIfPrerequisite(n int, pre [][]int, q [][]int) []bool {
	in := make([]int, n)
	nx := make([][]int, n)
	d := make([][]bool, n)
	for _, v := range pre {
		nx[v[0]] = append(nx[v[0]], v[1])
		in[v[1]]++
	}

	st := []int{}
	for i, v := range in {
		d[i] = make([]bool, n)
		if v == 0 {
			st = append(st, i)
		}
	}

	for i := 0; i < len(st); i++ {
		for _, x := range nx[st[i]] {
			in[x]--
			d[x][i] = true
			for j, jv := range d[i] {
				if jv {
					d[x][j] = true
				}
			}
			if in[x] == 0 {
				st = append(st, x)
			}
		}
	}

	ans := []bool{}
	for _, v := range q {
		ans = append(ans, d[v[0]][v[1]])
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

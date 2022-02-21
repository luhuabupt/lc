package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func pushDominoes(d string) string {
	n := len(d)
	st := [][2]int{}
	ans := make([]byte, n)
	for i, v := range d {
		if v == 'R' {
			st = append(st, [2]int{i, 1})
		} else if v == 'L' {
			st = append(st, [2]int{i, -1})
		}
		ans[i] = byte(v)
	}

	for len(st) > 0 {
		do := map[int]int{}
		for _, v := range st {
			i, d := v[0], v[1]
			if i+d >= 0 && i+d < n && ans[i+d] == '.' {
				do[i+d] += d
			}
		}

		st = [][2]int{}
		for k, v := range do {
			if v != 0 {
				if v == 1 {
					ans[k] = 'R'
				} else if v == -1 {
					ans[k] = 'L'
				}
				st = append(st, [2]int{k, v})
			}
		}
	}

	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)

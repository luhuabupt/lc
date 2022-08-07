package main

import "sort"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func smallestDifference(a []int, b []int) int {
	abs := func(a, b, ans int) int {
		x := a - b
		if x < 0 {
			x = b - a
		}

		if x < ans {
			return x
		}
		return ans
	}

	type p struct {
		v, pos int
	}
	mg := []p{}

	ma := map[int]bool{}
	for _, v := range a {
		mg = append(mg, p{v, 0})
		ma[v] = true
	}
	for _, v := range b {
		if ma[v] {
			return 0
		}
		mg = append(mg, p{v, 1})
	}

	sort.Slice(mg, func(i, j int) bool {
		return mg[i].v < mg[j].v
	})

	ans := 1 << 31
	ap, bp, ado, bdo := 0, 0, false, false
	for i := 0; i < len(mg); i++ {
		p := mg[i]
		if p.pos == 0 {
			ap = p.v
			ado = true
		} else {
			bp = p.v
			bdo = true
		}
		if ado && bdo {
			ans = abs(ap, bp, ans)
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

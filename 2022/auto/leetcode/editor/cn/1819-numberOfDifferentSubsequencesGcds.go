package main

import "fmt"

func main() {
	fmt.Println(countDifferentSubsequenceGCDs([]int{4, 6, 9}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func countDifferentSubsequenceGCDs(nums []int) int {
	// 所有质数
	prime := []int{}
	d := make([]bool, 501)
	for l := 2; l < 501; l++ {
		for i := l + l; i < 501; i += l {
			d[i] = true
		}
	}
	for i, v := range d {
		if i >= 2 && !v {
			prime = append(prime, i)
		}
	}
	// fmt.Println(prime)
	// []int{2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97,101,103,107,109,113,127,131,137,139,149,151,157,163,167,173,179,181,191,193,197,199,211,223,227,229,233,239,241,251,257,263,269,271,277,281,283,293,307,311,313,317,331,337,347,349,353,359,367,373,379,383,389,397,401,409,419,421,431,433,439,443,449,457,461,463,467,479,487,491,499}

	// 分解质因数
	var do func(v int) []int
	do = func(v int) []int {
		if v == 1 {
			return []int{}
		}
		for _, x := range prime {
			if v%x == 0 {
				return append([]int{x}, do(v/x)...)
			}
		}
		return []int{v}
	}
	//fmt.Println(do(6))
	//fmt.Println(do(8))
	//fmt.Println(do(12))

	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	type tp struct {
		end  bool
		next map[int]*tp
	}

	head := &tp{false, map[int]*tp{}}
	ans := map[int]bool{}
	c := map[int]bool{}
	all := nums[0]
	for _, v := range nums {
		all = gcd(v, nums[0])
		if v == 1 {
			continue
		}
		if c[v] {
			ans[v] = true
		} else {
			pa := do(v)
			t := head
			stv := 1
			fmt.Println(v, pa)
			for _, x := range pa {
				stv *= x
				if _, ok := t.next[x]; !ok {
					t.next[x] = &tp{false, map[int]*tp{}}
				}
				t = t.next[x]

				if t.end && len(t.next) > 0 {
					ans[stv] = true
				}
				if len(t.next) > 0 {
					ans[stv] = true
				}
			}
			t.end = true
		}
		c[v] = true
	}

	fmt.Println(ans)

	ans[all] = true
	return len(ans)
}

//leetcode submit region end(Prohibit modification and deletion)

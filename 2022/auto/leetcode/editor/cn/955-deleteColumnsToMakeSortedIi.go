package main

import "fmt"

func main() {
	//fmt.Println(minDeletionSize([]string{"zyx","wvu","tsr"}))
	//fmt.Println(minDeletionSize([]string{"jwkwdc", "etukoz"}))
	fmt.Println(minDeletionSize([]string{"doeeqiy", "yabhbqe", "twckqte"}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func minDeletionSize(s []string) int {
	n := len(s[0])
	d := make([]bool, n)
	f := make([]int, n)
	for i, _ := range f {
		f[i] = -1
	}
	for {
		for i := 1; i < len(s); i++ {
			v := s[i]
			for j, x := range v {
				//fmt.Println(i, j, d, string(x), x, s[i-1][j])
				if d[j] {
					continue
				}
				if int32(s[i-1][j]) < x {
					if f[j] == -1 {
						f[j] = i
					}
					break
				} else if int32(s[i-1][j]) > x {
					d[j] = true
					if f[j] > 0 {
						i = f[j] - 1
						break
					}
				}
			}
		}

		flag := true
		for i := 1; i < len(s); i++ {
			v := s[i]
			for j, x := range v {
				if d[j] {
					continue
				}
				if int32(s[i-1][j]) < x {
					break
				} else if int32(s[i-1][j]) > x {
					flag = false
					break
				}
			}
			if !flag {
				break
			}
		}
		if flag {
			break
		}
	}

	ans := 0
	for _, v := range d {
		if v {
			ans++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
